package comment

import (
	"context"
	"time"

	"github.com/iot-synergy/synergy-member-rpc/ent"
	"github.com/iot-synergy/synergy-member-rpc/ent/comment"
	"github.com/iot-synergy/synergy-member-rpc/ent/member"
	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

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
	query.Order(ent.Desc(comment.FieldCreateTime))
	count := int64(query.CountX(l.ctx))
	if in.GetPage() <= 0 {
		*in.Page = 1
	}
	if in.GetPageSize() <= 0 {
		*in.PageSize = 1
	}
	query.Limit(int(in.GetPageSize()))
	query.Offset(int((in.GetPage() - 1) * in.GetPageSize()))
	//

	all, err := query.All(l.ctx)

	if err != nil {
		return nil, err
	}

	infos := make([]*mms.CommentInfo, 0)
	for _, comm := range all {
		id := int64(comm.ID)
		createTime := comm.CreateTime.UnixMilli()
		result, err := l.svcCtx.DB.Member.Query().Where(member.ForeinIDEQ(comm.MemberId)).WithRanks().First(l.ctx)
		if err != nil {
			l.Logger.Error(err, err.Error())
			result = new(ent.Member)
		}
		info := mms.CommentInfo{
			Id:         &id,
			Title:      &comm.Title,
			Content:    &comm.Content,
			MemberId:   &comm.MemberId,
			CreateTime: &createTime,
			IsReply:    &comm.IsReply,
			UserName:   &result.Username,
			NickName:   &result.Nickname,
			Email:      &result.Email,
			Avatar:     &result.Avatar,
		}
		infos = append(infos, &info)
	}

	return &mms.CommentList{
		List:  infos,
		Count: &count,
	}, nil
}
