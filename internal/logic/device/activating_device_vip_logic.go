package device

import (
	"context"
	"github.com/iot-synergy/synergy-activation-code-rpc/activationcoderpcclient"

	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivatingDeviceVipLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActivatingDeviceVipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivatingDeviceVipLogic {
	return &ActivatingDeviceVipLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActivatingDeviceVipLogic) ActivatingDeviceVip(in *mms.ActivatingDeviceVipReq) (*mms.ActivatingDeviceVipResp, error) {
	data, err := l.svcCtx.SynergyActivationCodeRpc.ActivatingDeviceVip(l.ctx, &activationcoderpcclient.ActivatingDeviceVipReq{
		SerialNumber:   in.SerialNumber,
		ProductId:      in.ProductId,
		ActivationCode: in.ActivationCode,
	})
	if err != nil {
		return &mms.ActivatingDeviceVipResp{
			Code: data.Code,
			Msg:  data.Msg,
			Data: "",
		}, err
	}

	return &mms.ActivatingDeviceVipResp{
		Code: data.Code,
		Msg:  data.Msg,
		Data: "",
	}, nil
}
