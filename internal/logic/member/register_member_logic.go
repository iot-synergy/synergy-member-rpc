package member

import (
	"context"
	"strings"

	uuid "github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/synergy-common/utils/encrypt"
	"github.com/iot-synergy/synergy-common/utils/pointy"
	"github.com/iot-synergy/synergy-member-rpc/ent/member"
	"google.golang.org/grpc/metadata"

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
func (l *RegisterMemberLogic) RegisterMember(in *mms.MemberInfo) (*mms.RegisterMemberResp, error) {
	// todo: add your logic here and delete this line
	if in.Username == nil || *in.Username == "" {
		return &mms.RegisterMemberResp{
			Code: -1,
			Msg:  "username is null",
			Data: nil,
		}, nil
	}
	//判断username的唯一性
	if l.svcCtx.DB.Member.Query().Where(member.Username(in.GetUsername())).CountX(l.ctx) > 0 {
		return &mms.RegisterMemberResp{
			Code: -1,
			Msg:  "username already exists",
			Data: nil,
		}, nil
	}

	forein_id := metadata.ValueFromIncomingContext(l.ctx, "gateway-firebaseid")
	if len(forein_id) <= 0 {
		return &mms.RegisterMemberResp{
			Code: -1,
			Msg:  "firebaseid is null",
			Data: nil,
		}, nil
	}

	query := l.svcCtx.DB.Member.Create().
		SetForeinID(strings.Join(forein_id, "")).
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
			return &mms.RegisterMemberResp{
				Code: -1,
				Msg:  "nickname already exists",
				Data: nil,
			}, nil
		}
		query.SetNotNilNickname(in.Nickname)
	}

	// value := metadata.ValueFromIncomingContext(l.ctx, "gateway-firebaseid")
	// uid, err := uuid.FromString(strings.Join(value, ""))
	// if err != nil {
	// 	return nil, err
	// }
	// query.SetID(uid)

	if in.Password != nil {
		query.SetNotNilPassword(pointy.GetPointer(encrypt.BcryptEncrypt(*in.Password)))
	} else {
		password := "123456" //默认密码
		query.SetNotNilPassword(pointy.GetPointer(encrypt.BcryptEncrypt(password)))
	}
	result, err := query.Save(l.ctx)

	if err != nil {
		return &mms.RegisterMemberResp{
			Code: -1,
			Msg:  err.Error(),
			Data: nil,
		}, nil
	}

	return &mms.RegisterMemberResp{
		Code: 0,
		Msg:  "成功",
		Data: &mms.RegisterMemberRespData{Id: result.ID.String()},
	}, nil
}
