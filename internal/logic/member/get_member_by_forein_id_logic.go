package member

import (
	"context"
	"github.com/iot-synergy/synergy-common/utils/pointy"
	"github.com/iot-synergy/synergy-member-rpc/ent/member"
	"github.com/iot-synergy/synergy-member-rpc/internal/utils/dberrorhandler"

	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberByForeinIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberByForeinIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberByForeinIdLogic {
	return &GetMemberByForeinIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: member
func (l *GetMemberByForeinIdLogic) GetMemberByForeinId(in *mms.UUIDReq) (*mms.MemberInfo, error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.DB.Member.Query().Where(member.ForeinIDEQ(in.Id)).WithRanks().First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mms.MemberInfo{
		Id:        pointy.GetPointer(result.ID.String()),
		CreatedAt: pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt: pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Status:    pointy.GetPointer(uint32(result.Status)),
		Username:  &result.Username,
		Password:  &result.Password,
		Nickname:  &result.Nickname,
		RankId:    &result.RankID,
		Mobile:    &result.Mobile,
		Email:     &result.Email,
		Avatar:    &result.Avatar,
		ExpiredAt: pointy.GetPointer(result.ExpiredAt.UnixMilli()),
		ForeinId:  &result.ForeinID,
	}, nil
	return &mms.MemberInfo{}, nil
}
