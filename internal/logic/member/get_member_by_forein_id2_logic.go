package member

import (
	"context"
	"github.com/iot-synergy/synergy-common/utils/pointy"
	"github.com/iot-synergy/synergy-member-rpc/ent/member"
	"github.com/iot-synergy/synergy-member-rpc/internal/utils/dberrorhandler"
	"regexp"

	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberByForeinId2Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberByForeinId2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberByForeinId2Logic {
	return &GetMemberByForeinId2Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberByForeinId2Logic) GetMemberByForeinId2(in *mms.UUIDReq) (*mms.MemberInfoResp, error) {
	re := regexp.MustCompile("^peckperk_peckperk-")
	result, err := l.svcCtx.DB.Member.Query().Where(member.ForeinIDEQ(re.ReplaceAllLiteralString(in.Id, ""))).WithRanks().First(l.ctx)
	if err != nil {
		return &mms.MemberInfoResp{
			Code: 0,
			Msg:  "成功",
			Data: &mms.MemberInfoRespData{
				IsExist:    0,
				MemberInfo: nil,
			},
		}, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mms.MemberInfoResp{
		Code: 0,
		Msg:  "成功",
		Data: &mms.MemberInfoRespData{
			IsExist: 1,
			MemberInfo: &mms.MemberInfo{
				Id:        pointy.GetPointer(result.ID.String()),
				CreatedAt: pointy.GetPointer(result.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(result.UpdatedAt.UnixMilli()),
				Status:    pointy.GetPointer(uint32(result.Status)),
				Username:  &result.Username,
				Password:  &result.Password,
				Nickname:  &result.Nickname,
				RankId:    &result.RankID,
				Mobile:    &result.Mobile,
				Email:     &result.Email,
				Avatar:    &result.Avatar,
				ExpiredAt: pointy.GetPointer(result.ExpiredAt.UnixMilli()),
				ForeinId:  &result.ForeinID,
			},
		},
	}, nil
}
