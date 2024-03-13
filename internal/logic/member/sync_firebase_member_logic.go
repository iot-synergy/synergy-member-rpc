package member

import (
	"context"
	"time"

	"github.com/iot-synergy/synergy-common/utils/encrypt"
	"github.com/iot-synergy/synergy-common/utils/pointy"
	"github.com/iot-synergy/synergy-fcm/synergyFCM"
	"github.com/iot-synergy/synergy-member-rpc/ent/member"
	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncFirebaseMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncFirebaseMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncFirebaseMemberLogic {
	return &SyncFirebaseMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SyncFirebaseMemberLogic) SyncFirebaseMember(in *mms.Empty) (*mms.SyncMemberResp, error) {
	// todo: add your logic here and delete this line
	page := 0
	pageSize := 10
	pageToken := ""

	all := 0
	new := 0
	updated := 0

	for {
		resp, err := l.svcCtx.Fcm.GetUsers(l.ctx, &synergyFCM.PageInfoReq{
			Page:      uint64(page),
			PageSize:  uint64(pageSize),
			PageToken: pageToken,
		})

		if err != nil {
			return &mms.SyncMemberResp{
				Code: -1,
				Msg:  "error",
				Data: &mms.SyncResult{
					All:     int64(all),
					New:     int64(new),
					Updated: int64(updated),
				},
			}, err
		}

		for _, d := range resp.Data {

			member, err := l.svcCtx.DB.Member.Query().Where(member.ForeinID(d.Uid)).First(l.ctx)
			if err != nil && member != nil {
				member.Username = d.UserName
				member.Nickname = d.NickName
				member.Avatar = *d.Avatar
				member.Email = *d.Email
				member.UpdatedAt = time.Now()
				l.svcCtx.DB.Member.UpdateOne(member)
				all += 1
				updated += 1
			} else {
				all += 1
				query := l.svcCtx.DB.Member.Create().
					SetForeinID(d.Uid).
					SetNotNilUsername(&d.UserName).
					SetNotNilEmail(d.Email).
					SetNotNilAvatar(d.Avatar).
					SetNotNilNickname(&d.NickName).
					SetNotNilPassword(pointy.GetPointer(encrypt.BcryptEncrypt("Abcd1234")))

				_, err := query.Save(l.ctx)

				if err != nil {
					continue
				}
				new += 1
			}

		}

		if resp.NextPageToken == "" {
			break
		}

		pageToken = resp.NextPageToken
	}

	return &mms.SyncMemberResp{
		Code: 0,
		Msg:  "success",
		Data: &mms.SyncResult{
			All:     int64(all),
			New:     int64(new),
			Updated: int64(updated),
		},
	}, nil
}
