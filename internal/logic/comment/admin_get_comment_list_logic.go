package comment

import (
	"context"
	"github.com/iot-synergy/synergy-member-rpc/ent/comment"
	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminGetCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminGetCommentListLogic {
	return &AdminGetCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: comment
func (l *AdminGetCommentListLogic) AdminGetCommentList(in *mms.CommentListReq) (*mms.CommentList, error) {
	// todo: add your logic here and delete this line
	query := l.svcCtx.DB.Comment.Query()

	if in.GetIsReply() == 0 {
		query.Where(comment.IsReply(false))
	} else if in.GetIsReply() == 1 {
		query.Where(comment.IsReply(true))
	}
	if in.GetTitle() != "" {
		query.Where(comment.TitleContains(in.GetTitle()))
	}
	if in.GetContent() != "" {
		query.Where(comment.ContentContains(in.GetContent()))
	}
	if in.GetCommentTime() != nil && len(in.GetCommentTime()) == 2 {
		query.Where(comment.CreateTimeGTE(time.UnixMilli(in.GetCommentTime()[0])))
		query.Where(comment.CreateTimeLTE(time.UnixMilli(in.GetCommentTime()[1])))
	}

	query.Limit(int(in.GetPageSize()))
	query.Offset(int((in.GetPage() - 1) * in.GetPageSize()))

	all, err := query.All(l.ctx)

	if err != nil {
		return nil, err
	}

	infos := make([]*mms.CommentInfo, 0)
	for _, comm := range all {
		id := int64(comm.ID)
		createTime := comm.CreateTime.UnixMilli()
		info := mms.CommentInfo{
			Id:         &id,
			Title:      &comm.Title,
			Content:    &comm.Content,
			MemberId:   &comm.MemberId,
			CreateTime: &createTime,
			IsReply:    &comm.IsReply,
		}
		infos = append(infos, &info)
	}

	return &mms.CommentList{
		List: infos,
	}, nil
}
