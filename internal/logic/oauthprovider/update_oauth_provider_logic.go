package oauthprovider

import (
	"context"

	"github.com/iot-synergy/synergy-member-rpc/internal/utils/dberrorhandler"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/iot-synergy/synergy-common/i18n"
)

type UpdateOauthProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOauthProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOauthProviderLogic {
	return &UpdateOauthProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOauthProviderLogic) UpdateOauthProvider(in *mms.OauthProviderInfo) (*mms.BaseResp, error) {
	err := l.svcCtx.DB.OauthProvider.UpdateOneID(*in.Id).
		SetNotNilName(in.Name).
		SetNotNilClientID(in.ClientId).
		SetNotNilClientSecret(in.ClientSecret).
		SetNotNilRedirectURL(in.RedirectUrl).
		SetNotNilScopes(in.Scopes).
		SetNotNilAuthURL(in.AuthUrl).
		SetNotNilTokenURL(in.TokenUrl).
		SetNotNilAuthStyle(in.AuthStyle).
		SetNotNilInfoURL(in.InfoUrl).
		Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	if _, ok := providerConfig[*in.Name]; ok {
		delete(providerConfig, *in.Name)
	}

	return &mms.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
