package token

import (
	"context"
	"time"

	"github.com/iot-synergy/synergy-common/config"
	"github.com/iot-synergy/synergy-member-rpc/internal/utils/dberrorhandler"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/iot-synergy/synergy-common/msg/logmsg"
	"github.com/iot-synergy/synergy-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/iot-synergy/synergy-common/i18n"

	"github.com/iot-synergy/synergy-member-rpc/ent/token"

	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type BlockUserAllTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBlockUserAllTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlockUserAllTokenLogic {
	return &BlockUserAllTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BlockUserAllTokenLogic) BlockUserAllToken(in *mms.UUIDReq) (*mms.BaseResp, error) {
	err := l.svcCtx.DB.Token.Update().Where(token.UUIDEQ(uuidx.ParseUUIDString(in.Id))).SetStatus(0).Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	tokenData, err := l.svcCtx.DB.Token.Query().
		Where(token.UUIDEQ(uuidx.ParseUUIDString(in.Id))).
		Where(token.StatusEQ(0)).
		Where(token.ExpiredAtGT(time.Now())).
		All(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	for _, v := range tokenData {
		expiredTime := v.ExpiredAt.Sub(time.Now())
		if expiredTime > 0 {
			err = l.svcCtx.Redis.Set(l.ctx, config.RedisTokenPrefix+v.Token, "1", expiredTime).Err()
			if err != nil {
				logx.Errorw(logmsg.RedisError, logx.Field("detail", err.Error()))
				return nil, errorx.NewInternalError(i18n.RedisError)
			}
		}
	}

	return &mms.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
