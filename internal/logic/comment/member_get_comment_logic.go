package comment

import (
	"context"
	"github.com/iot-synergy/synergy-member-rpc/ent/reply"
	"google.golang.org/grpc/metadata"
	"strings"

	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberGetCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberGetCommentLogic {
	return &MemberGetCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: comment
func (l *MemberGetCommentLogic) MemberGetComment(in *mms.CommentIdReq) (*mms.CommentInfoResp, error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.DB.Comment.Get(l.ctx, uint64(in.GetId()))

	if err != nil {
		return nil, err
	}

	incomingContext, ok := metadata.FromIncomingContext(l.ctx)
	if ok {
		value := incomingContext.Get("gateway-firebaseid")
		if value == nil || len(value) == 0 {
			return &mms.CommentInfoResp{
				Code: -1,
				Msg:  "member token is null",
				Data: nil,
			}, nil
		}
		if data.MemberId != strings.Join(value, "") {
			return &mms.CommentInfoResp{
				Code: -1,
				Msg:  "member token error",
				Data: nil,
			}, nil
		}
	}
	id := int64(data.ID)
	createTime := data.CreateTime.UnixMilli()
	updateTime := data.UpdateTime.UnixMilli()

	replies, err := l.svcCtx.DB.Reply.Query().Where(reply.CommentID(data.ID)).All(l.ctx)

	if err != nil {
		return &mms.CommentInfoResp{
			Code: -1,
			Msg:  err.Error(),
			Data: nil,
		}, nil
	}

	replyInfos := make([]*mms.ReplyInfo, 0)

	for _, reply := range replies {
		replyId := int64(reply.ID)
		createTime2 := reply.CreateTime.UnixMilli()
		updateTime2 := reply.UpdateTime.UnixMilli()
		m := &mms.ReplyInfo{
			Id:         &replyId,
			CommentId:  &id,
			Reply:      &reply.Reply,
			AdminId:    &reply.AdminId,
			AdminName:  &reply.AdminName,
			CreateTime: &createTime2,
			UpdateTime: &updateTime2,
		}
		replyInfos = append(replyInfos, m)
	}

	return &mms.CommentInfoResp{
		Code: 0,
		Msg:  "成功",
		Data: &mms.CommentInfo{
			Id:         &id,
			Title:      &data.Title,
			Content:    &data.Content,
			MemberId:   &data.MemberId,
			CreateTime: &createTime,
			UpdateTime: &updateTime,
			Reply:      replyInfos,
		},
	}, nil
}
