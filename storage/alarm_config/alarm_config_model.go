package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

var _ Alarm_configModel = (*customAlarm_configModel)(nil)

type (
	// Alarm_configModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAlarm_configModel.
	Alarm_configModel interface {
		alarm_configModel
		FindOneByUserIdAndDeviceSn(ctx context.Context, addxUserId, serialNumber string) (*Alarm_config, error)
		SetConfigByUserIdAndDeviceSn(ctx context.Context, addxUserId, serialNumber string, config int32) error
	}

	customAlarm_configModel struct {
		*defaultAlarm_configModel
	}
)

// NewAlarm_configModel returns a model for the mongo.
func NewAlarm_configModel(url, db, collection string) Alarm_configModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customAlarm_configModel{
		defaultAlarm_configModel: newDefaultAlarm_configModel(conn),
	}
}

func (m *customAlarm_configModel) FindOneByUserIdAndDeviceSn(ctx context.Context, userId, deviceSn string) (*Alarm_config, error) {
	var data Alarm_config
	err := m.conn.FindOne(ctx, &data, bson.M{"addxUserId": userId, "serialNumber": deviceSn})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *customAlarm_configModel) SetConfigByUserIdAndDeviceSn(ctx context.Context, addxUserId, serialNumber string, config int32) error {
	_, err := m.conn.UpdateOne(ctx, bson.M{"addxUserId": addxUserId, "serialNumber": serialNumber}, bson.M{"config": config})
	return err
}
