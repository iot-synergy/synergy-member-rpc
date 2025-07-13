package member

import (
	"context"
	"errors"
	"strings"

	"github.com/iot-synergy/synergy-common/utils/pointy"
	"google.golang.org/grpc/metadata"

	"github.com/iot-synergy/synergy-member-rpc/ent/member"
	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberLogic {
	return &GetMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: member
func (l *GetMemberLogic) GetMember(in *mms.Empty) (*mms.MemberInfoResp, error) {
	// todo: add your logic here and delete this line
	forein_id := metadata.ValueFromIncomingContext(l.ctx, "gateway-firebaseid")

	if len(forein_id) <= 0 {
		return nil, errors.New("firebaseid is null")
	}

	result, err := l.svcCtx.DB.Member.Query().Where(member.ForeinIDEQ(strings.Join(forein_id, ""))).First(l.ctx)
	if result == nil || err != nil {
		return &mms.MemberInfoResp{
			Code: 0,
			Msg:  "成功",
			Data: &mms.MemberInfoRespData{
				IsExist:    0,
				MemberInfo: nil,
			},
		}, nil
	}

	return &mms.MemberInfoResp{
		Code: 0,
		Msg:  "成功",
		Data: &mms.MemberInfoRespData{
			IsExist: 1,
			MemberInfo: &mms.MemberInfo{
				Id:        pointy.GetPointer(result.ID.String()),
				CreatedAt: pointy.GetPointer(result.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(result.UpdatedAt.UnixMilli()),
				Status:    pointy.GetPointer(uint32(result.Status)),
				Username:  &result.Username,
				Nickname:  &result.Nickname,
				RankId:    &result.RankID,
				Mobile:    &result.Mobile,
				Email:     &result.Email,
				Avatar:    &result.Avatar,
				ExpiredAt: pointy.GetPointer(result.ExpiredAt.UnixMilli()),
				ForeinId:  &result.ForeinID,
				Gender:    &result.Gender,
				Birthday:  &result.Birthday,
			},
		},
	}, nil
}
