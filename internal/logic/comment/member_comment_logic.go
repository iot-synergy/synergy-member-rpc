package comment

import (
	"context"
	"google.golang.org/grpc/metadata"
	"strconv"
	"strings"
	"time"

	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberCommentLogic {
	return &MemberCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Comment management
func (l *MemberCommentLogic) MemberComment(in *mms.CommentInfo) (*mms.MemberCommentResp, error) {
	// todo: add your logic here and delete this line

	if *in.Title == "" || *in.Content == "" {
		return &mms.MemberCommentResp{
			Code: -1,
			Msg:  "title or content is null",
			Data: "",
		}, nil
	}

	if len(*in.Title) > 256 || len(*in.Content) > 2048 {
		return &mms.MemberCommentResp{
			Code: -1,
			Msg:  "title length exceeds 256 or content length exceeds 2048",
			Data: "",
		}, nil
	}

	query := l.svcCtx.DB.Comment.Create().
		SetTitle(*in.Title).
		SetContent(*in.Content).
		SetCreateTime(time.Now()).
		SetUpdateTime(time.Now())

	incomingContext, ok := metadata.FromIncomingContext(l.ctx)

	if ok {
		value := incomingContext.Get("gateway-firebaseid")
		if value == nil || len(value) == 0 {
			return &mms.MemberCommentResp{
				Code: -1,
				Msg:  "member token is null",
				Data: "",
			}, nil
		}
		query.SetMemberId(strings.Join(value, ""))
	}

	save, err := query.Save(l.ctx)

	if err != nil {
		return &mms.MemberCommentResp{
			Code: -1,
			Msg:  err.Error(),
			Data: "",
		}, nil
	}
	return &mms.MemberCommentResp{
		Code: 0,
		Msg:  "成功",
		Data: strconv.FormatUint(save.ID, 10),
	}, nil
}
