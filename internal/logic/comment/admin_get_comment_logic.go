package comment

import (
	"context"
	"github.com/iot-synergy/synergy-member-rpc/ent/reply"
	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminGetCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminGetCommentLogic {
	return &AdminGetCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: comment
func (l *AdminGetCommentLogic) AdminGetComment(in *mms.CommentIdReq) (*mms.CommentInfo, error) {
	// todo: add your logic here and delete this line

	data, err := l.svcCtx.DB.Comment.Get(l.ctx, uint64(in.GetId()))

	if err != nil {
		return nil, err
	}

	id := int64(data.ID)
	createTime := data.CreateTime.UnixMilli()
	updateTime := data.UpdateTime.UnixMilli()

	replies, err := l.svcCtx.DB.Reply.Query().Where(reply.CommentID(data.ID)).All(l.ctx)

	if err != nil {
		return nil, err
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
			CreatTime:  &createTime2,
			UpdateTime: &updateTime2,
		}
		replyInfos = append(replyInfos, m)
	}

	return &mms.CommentInfo{
		Id:         &id,
		Title:      &data.Title,
		Content:    &data.Content,
		MemberId:   &data.MemberId,
		CreatTime:  &createTime,
		UpdateTime: &updateTime,
		Reply:      replyInfos,
	}, nil
}
