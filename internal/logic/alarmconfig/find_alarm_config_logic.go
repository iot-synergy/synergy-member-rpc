package alarmconfig

import (
	"context"
	model "github.com/iot-synergy/synergy-member-rpc/storage/alarm_config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/metadata"
	"regexp"
	"strings"
	"time"

	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAlarmConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAlarmConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAlarmConfigLogic {
	return &FindAlarmConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Member management
func (l *FindAlarmConfigLogic) FindAlarmConfig(in *mms.FindAlarmConfigReq) (*mms.AlarmConfigResp, error) {
	// 获取用户id
	value := metadata.ValueFromIncomingContext(l.ctx, "gateway-firebaseid")
	if len(value) <= 0 {
		return &mms.AlarmConfigResp{
			Code: -1,
			Msg:  "用户未登录",
			Data: nil,
		}, nil
	}
	forein_id := regexp.MustCompile("^peckperk-").ReplaceAllLiteralString(strings.Join(value, ""), "")
	alarmConfig, err := l.svcCtx.AlarmConfigModel.FindOneByUserIdAndDeviceSn(l.ctx, forein_id, in.SerialNumber)
	if err != nil {
		return nil, err
	}
	if alarmConfig == nil || alarmConfig.ID.Hex() == "" {
		alarmConfig = &model.Alarm_config{
			ID:           primitive.ObjectID{},
			UpdateAt:     time.Time{},
			CreateAt:     time.Time{},
			AddxUserId:   forein_id,
			SerialNumber: in.SerialNumber,
			Config:       0,
		}
		err = l.svcCtx.AlarmConfigModel.Insert(l.ctx, alarmConfig)
		if err != nil {
			return nil, err
		}
	}

	id := alarmConfig.ID.Hex()
	return &mms.AlarmConfigResp{
		Code: 0,
		Msg:  "成功",
		Data: &mms.AlarmConfigInit{
			Id:           &id,
			AddxUserId:   &alarmConfig.AddxUserId,
			SerialNumber: &alarmConfig.SerialNumber,
			Config:       &alarmConfig.Config,
		},
	}, nil
}
