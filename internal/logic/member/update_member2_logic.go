package member

import (
	"context"
	"errors"
	"github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/iot-synergy/synergy-member-rpc/internal/utils/dberrorhandler"
	"google.golang.org/grpc/metadata"
	"strings"

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
func (l *UpdateMember2Logic) UpdateMember2(in *mms.MemberInfo) (*mms.BaseResp, error) {
	// todo: add your logic here and delete this line
	value := metadata.ValueFromIncomingContext(l.ctx, "gateway-firebaseid")
	uid, err := uuid.FromString(strings.Join(value, ""))
	if err != nil {
		return nil, err
	}

	if len(in.GetNickname()) > 255 || len(in.GetAvatar()) > 512 {
		return nil, errors.New("nickname length more 255 or avatar length more 512")
	}

	query := l.svcCtx.DB.Member.UpdateOneID(uid).
		SetNotNilNickname(in.Nickname).
		SetNotNilAvatar(in.Avatar)

	err = query.Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mms.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
