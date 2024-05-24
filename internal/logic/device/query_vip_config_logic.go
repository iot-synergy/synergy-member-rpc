package device

import (
	"context"
	"github.com/iot-synergy/synergy-activation-code-rpc/activationcoderpcclient"

	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryVipConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryVipConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryVipConfigLogic {
	return &QueryVipConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryVipConfigLogic) QueryVipConfig(in *mms.QueryVipConfigReq) (*mms.QueryVipConfigResp, error) {
	config, err := l.svcCtx.SynergyActivationCodeRpc.QueryVipConfig(l.ctx, &activationcoderpcclient.QueryVipConfigReq{
		SerialNumber: in.SerialNumber,
	})
	if err != nil {
		return &mms.QueryVipConfigResp{
			Code: -1,
			Msg:  "失败",
			Data: nil,
		}, err
	}
	return &mms.QueryVipConfigResp{
		Code: 0,
		Msg:  "成功",
		Data: &mms.QueryVipConfigRespData{
			VipLevel: config.VipLevel,
			EndTime:  config.EndTime,
			Sn:       config.Sn,
		},
	}, nil
}
