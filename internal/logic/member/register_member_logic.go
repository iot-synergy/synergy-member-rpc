package member

import (
	"context"
	"errors"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/iot-synergy/synergy-common/utils/encrypt"
	"github.com/iot-synergy/synergy-common/utils/pointy"
	"github.com/iot-synergy/synergy-member-rpc/ent/member"
	"github.com/iot-synergy/synergy-member-rpc/internal/utils/dberrorhandler"
	"google.golang.org/grpc/metadata"
	"strings"

	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterMemberLogic {
	return &RegisterMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: member
func (l *RegisterMemberLogic) RegisterMember(in *mms.MemberInfo) (*mms.BaseUUIDResp, error) {
	// todo: add your logic here and delete this line
	if in.Username == nil || *in.Username == "" {
		return nil, errors.New("username is null")
	}
	//判断username的唯一性
	if l.svcCtx.DB.Member.Query().Where(member.Username(in.GetUsername())).CountX(l.ctx) > 0 {
		return nil, errors.New("username already exists")
	}

	query := l.svcCtx.DB.Member.Create().
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetNotNilUsername(in.Username).
		SetNotNilRankID(in.RankId).
		SetNotNilMobile(in.Mobile).
		SetNotNilEmail(in.Email).
		SetNotNilAvatar(in.Avatar).
		SetNotNilWechatOpenID(in.WechatId).
		SetNotNilExpiredAt(pointy.GetTimeMilliPointer(in.ExpiredAt))

	if in.Nickname == nil || *in.Nickname == "" {
		v1, _ := uuid.NewV1()
		nickName := "User-" + v1.String()
		query.SetNotNilNickname(&nickName)
	} else {
		//校验昵称的唯一性
		if l.svcCtx.DB.Member.Query().Where(member.Nickname(in.GetNickname())).CountX(l.ctx) > 0 {
			return nil, errors.New("nickname already exists")
		}
		query.SetNotNilNickname(in.Nickname)
	}

	value := metadata.ValueFromIncomingContext(l.ctx, "gateway-firebaseid")
	uid, err := uuid.FromString(strings.Join(value, ""))
	if err != nil {
		return nil, err
	}
	//校验id的唯一性
	if l.svcCtx.DB.Member.Query().Where(member.ID(uid)).CountX(l.ctx) > 0 {
		return nil, errors.New("id already exists")
	}
	query.SetID(uid)

	if in.Password != nil {
		query.SetNotNilPassword(pointy.GetPointer(encrypt.BcryptEncrypt(*in.Password)))
	} else {
		password := "123456" //默认密码
		query.SetNotNilPassword(pointy.GetPointer(encrypt.BcryptEncrypt(password)))
	}
	result, err := query.Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mms.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil
}
