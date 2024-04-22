package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Alarm_config struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	// TODO: Fill your own fields
	UpdateAt     time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt     time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
	AddxUserId   string    `bson:"addxUserId,omitempty" json:"userId,omitempty"`
	SerialNumber string    `bson:"serialNumber,omitempty" json:"deviceSn,omitempty"`
	//0 提示
	//1 不提示
	Config int32 `bson:"config,omitempty" json:"config,omitempty"`
}
