package member

import (
	"context"

	"github.com/iot-synergy/synergy-common/utils/encrypt"
	"github.com/iot-synergy/synergy-common/utils/pointy"

	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/iot-synergy/synergy-member-rpc/internal/utils/dberrorhandler"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/iot-synergy/synergy-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberLogic {
	return &UpdateMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMemberLogic) UpdateMember(in *mms.MemberInfo) (*mms.BaseResp, error) {
	query := l.svcCtx.DB.Member.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetNotNilUsername(in.Username).
		SetNotNilNickname(in.Nickname).
		SetNotNilRankID(in.RankId).
		SetNotNilMobile(in.Mobile).
		SetNotNilEmail(in.Email).
		SetNotNilAvatar(in.Avatar).
		SetNotNilWechatOpenID(in.WechatId).
		SetNotNilExpiredAt(pointy.GetTimeMilliPointer(in.ExpiredAt)).
		SetNotNilGender(in.Gender).
		SetNotNilBirthday(in.Birthday)

	if in.Password != nil && *in.Password != "" {
		query.SetNotNilPassword(pointy.GetPointer(encrypt.BcryptEncrypt(*in.Password)))
	}

	err := query.Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mms.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
