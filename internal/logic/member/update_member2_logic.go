package member

import (
	"context"
	"strings"

	"github.com/iot-synergy/synergy-member-rpc/ent/member"
	"google.golang.org/grpc/metadata"

	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMember2Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMember2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMember2Logic {
	return &UpdateMember2Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: member
func (l *UpdateMember2Logic) UpdateMember2(in *mms.MemberInfo) (*mms.UpdateMember2Resp, error) {
	// todo: add your logic here and delete this line
	value := metadata.ValueFromIncomingContext(l.ctx, "gateway-firebaseid")
	forein_id := strings.Join(value, "")
	if len(forein_id) <= 0 {
		return &mms.UpdateMember2Resp{
			Code: -1,
			Msg:  "forein_id is null",
			Data: false,
		}, nil
	}

	if len(in.GetNickname()) > 255 || len(in.GetAvatar()) > 512 {
		return &mms.UpdateMember2Resp{
			Code: -1,
			Msg:  "nickname length more 255 or avatar length more 512",
			Data: false,
		}, nil
	}

	//校验昵称的唯一性
	if l.svcCtx.DB.Member.Query().Where(member.Nickname(in.GetNickname())).CountX(l.ctx) > 0 {
		return &mms.UpdateMember2Resp{
			Code: -1,
			Msg:  "nickname already exists",
			Data: false,
		}, nil
	}

	query := l.svcCtx.DB.Member.Update().Where(member.ForeinIDEQ(forein_id)).
		SetNotNilNickname(in.Nickname).
		SetNotNilAvatar(in.Avatar).
		SetNotNilMobile(in.Mobile).
		SetNotNilEmail(in.Email)

	err := query.Exec(l.ctx)

	if err != nil {
		return &mms.UpdateMember2Resp{
			Code: -1,
			Msg:  err.Error(),
			Data: false,
		}, nil
	}

	return &mms.UpdateMember2Resp{
		Code: 0,
		Msg:  "成功",
		Data: true,
	}, nil
}
