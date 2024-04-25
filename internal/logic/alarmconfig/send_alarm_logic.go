package alarmconfig

import (
	"context"
	"github.com/iot-synergy/synergy-addx-proxy/synergy_addx_proxy_client"

	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendAlarmLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendAlarmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendAlarmLogic {
	return &SendAlarmLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendAlarmLogic) SendAlarm(in *mms.SendAlarmReq) (*mms.EmptyResp, error) {
	l.Error("send alarm(addxUserId:" + in.GetAddxUserId() + ",serialNumber:" + in.GetSerialNumber() + ")")
	alarmConfig, err := l.svcCtx.AlarmConfigModel.FindOneByUserIdAndDeviceSn(l.ctx, in.GetAddxUserId(), in.GetSerialNumber())
	if err != nil {
		l.Error(err)
		return nil, err
	}
	if alarmConfig == nil || alarmConfig.Config == 0 {
		return nil, nil
	}
	_, err = l.svcCtx.AddxRpc.SendAlarm(l.ctx, &synergy_addx_proxy_client.SendAlarmReq{
		AddxUserId:   in.AddxUserId,
		SerialNumber: in.SerialNumber,
	})
	if err != nil {
		l.Error(err)
		return nil, err
	}
	l.Error("send alarm exit")
	return &mms.EmptyResp{}, nil
}
