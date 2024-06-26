package comment

import (
	"context"
	"github.com/iot-synergy/synergy-member-rpc/ent/reply"

	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminGetReplyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminGetReplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminGetReplyListLogic {
	return &AdminGetReplyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: comment
func (l *AdminGetReplyListLogic) AdminGetReplyList(in *mms.ReplyReq) (*mms.ReplyList, error) {
	// todo: add your logic here and delete this line

	query := l.svcCtx.DB.Reply.Query().
		Where(reply.AdminId(in.GetAdminId()))
	if in.GetReply() != "" {
		query.Where(reply.ReplyContains(in.GetReply()))
	}

	count := int64(query.CountX(l.ctx))
	if in.GetPage() <= 0 {
		*in.Page = 1
	}
	if in.GetPageSize() <= 0 {
		*in.PageSize = 1
	}
	query.Limit(int(in.GetPageSize())).
		Offset(int((in.GetPage() - 1) * in.GetPageSize()))
	replies, err := query.All(l.ctx)

	if err != nil {
		return nil, err
	}

	replyList := make([]*mms.ReplyInfo, 0)

	for _, reply := range replies {
		id := int64(reply.ID)
		commentId := int64(reply.CommentID)
		createTime := reply.CreateTime.UnixMilli()
		updateTime := reply.UpdateTime.UnixMilli()
		replyInfo := mms.ReplyInfo{
			Id:         &id,
			CommentId:  &commentId,
			Reply:      &reply.Reply,
			AdminId:    &reply.AdminId,
			AdminName:  &reply.AdminName,
			CreateTime: &createTime,
			UpdateTime: &updateTime,
		}
		replyList = append(replyList, &replyInfo)
	}

	return &mms.ReplyList{
		ReplyList: replyList,
		Count:     &count,
	}, nil
}
