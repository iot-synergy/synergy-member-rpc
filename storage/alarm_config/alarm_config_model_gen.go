// Code generated by goctl. DO NOT EDIT.
package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type alarm_configModel interface {
	Insert(ctx context.Context, data *Alarm_config) error
	FindOne(ctx context.Context, id string) (*Alarm_config, error)
	Update(ctx context.Context, data *Alarm_config) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type defaultAlarm_configModel struct {
	conn *mon.Model
}

func newDefaultAlarm_configModel(conn *mon.Model) *defaultAlarm_configModel {
	return &defaultAlarm_configModel{conn: conn}
}

func (m *defaultAlarm_configModel) Insert(ctx context.Context, data *Alarm_config) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	_, err := m.conn.InsertOne(ctx, data)
	return err
}

func (m *defaultAlarm_configModel) FindOne(ctx context.Context, id string) (*Alarm_config, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data Alarm_config

	err = m.conn.FindOne(ctx, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAlarm_configModel) Update(ctx context.Context, data *Alarm_config) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()

	res, err := m.conn.UpdateOne(ctx, bson.M{"_id": data.ID}, bson.M{"$set": data})
	return res, err
}

func (m *defaultAlarm_configModel) Delete(ctx context.Context, id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, ErrInvalidObjectId
	}

	res, err := m.conn.DeleteOne(ctx, bson.M{"_id": oid})
	return res, err
}
