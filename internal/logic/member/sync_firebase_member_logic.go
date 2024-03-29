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

func (l *SyncFirebaseMemberLogic) SyncFirebaseMember(in *mms.SyncMemberReq) (*mms.SyncMemberResp, error) {
	// todo: add your logic here and delete this line
	page := 0
	pageSize := 100

	all := 0
	new := 0
	updated := 0

	resp, err := l.svcCtx.Fcm.GetUsers(context.Background(), &synergyFCM.PageInfoReq{
		Page:      uint64(page),
		PageSize:  uint64(pageSize),
		PageToken: in.NextPageToken,
	})

	if err != nil {
		return &mms.SyncMemberResp{
			Code:          -1,
			Msg:           "error",
			Total:         0,
			NextPageToken: resp.NextPageToken,
		}, err
	}

	syncMemberResp := &mms.SyncMemberResp{
		Code:          0,
		Msg:           "success",
		Total:         0,
		NextPageToken: resp.NextPageToken,
		Data:          []*mms.MemberInfo{},
	}

	for _, d := range resp.Data {
		mem, err := l.svcCtx.DB.Member.Query().Where(member.ForeinID(d.Uid)).First(context.Background())
		if err == nil && mem != nil {
			if mem.Username != d.UserName || mem.Nickname != d.NickName || mem.Avatar != *d.Avatar || mem.Email != *d.Email {
				mem.Username = d.UserName
				mem.Nickname = d.NickName
				mem.Avatar = *d.Avatar
				mem.Email = *d.Email
				mem.UpdatedAt = time.Now()
				l.svcCtx.DB.Member.UpdateOneID(mem.ID).
					SetNotNilUsername(&d.UserName).
					SetNotNilEmail(d.Email).
					SetNotNilAvatar(d.Avatar).
					SetNotNilNickname(&d.NickName).Exec(context.Background())
				// l.svcCtx.DB.Member.UpdateOneID(mem.ID)..Exec(context.Background())
				updated += 1
			}

			all += 1

		} else {
			all += 1
			query := l.svcCtx.DB.Member.Create().
				SetForeinID(d.Uid).
				SetNotNilUsername(&d.UserName).
				SetNotNilEmail(d.Email).
				SetNotNilAvatar(d.Avatar).
				SetNotNilNickname(&d.NickName).
				SetNotNilPassword(pointy.GetPointer(encrypt.BcryptEncrypt("Abcd1234")))

			_, err := query.Save(context.Background())

			if err != nil {
				l.Error("query.Save error:" + err.Error())

			} else {
				new += 1
			}

		}

		syncMemberResp.Data = append(syncMemberResp.Data, &mms.MemberInfo{
			Username: &d.UserName,
			Nickname: &d.NickName,
			Email:    d.Email,
			Avatar:   d.Avatar,
			ForeinId: &d.Uid,
		})

	}

	return syncMemberResp, nil
}
