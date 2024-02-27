package member

import (
	"context"
	"github.com/iot-synergy/synergy-common/utils/pointy"
	"github.com/iot-synergy/synergy-common/utils/uuidx"
	"github.com/iot-synergy/synergy-member-rpc/internal/utils/dberrorhandler"
	"google.golang.org/grpc/metadata"
	"strings"

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
func (l *GetMemberLogic) GetMember(in *mms.Empty) (*mms.MemberInfo, error) {
	// todo: add your logic here and delete this line
	value := metadata.ValueFromIncomingContext(l.ctx, "gateway-firebaseid")

	result, err := l.svcCtx.DB.Member.Get(l.ctx, uuidx.ParseUUIDString(strings.Join(value, "")))
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mms.MemberInfo{
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
	}, nil
	return &mms.MemberInfo{}, nil
}
