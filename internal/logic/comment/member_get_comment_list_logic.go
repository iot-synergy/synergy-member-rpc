package comment

import (
	"context"
	"github.com/iot-synergy/synergy-member-rpc/ent/comment"
	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"
	"google.golang.org/grpc/metadata"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberGetCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberGetCommentListLogic {
	return &MemberGetCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: comment
func (l *MemberGetCommentListLogic) MemberGetCommentList(in *mms.CommentListReq) (*mms.CommentListResp, error) {
	// todo: add your logic here and delete this line
	query := l.svcCtx.DB.Comment.Query()

	incomingContext, ok := metadata.FromIncomingContext(l.ctx)
	if ok {
		value := incomingContext.Get("gateway-firebaseid")
		if value == nil || len(value) == 0 {
			return &mms.CommentListResp{
				Code: -1,
				Msg:  "member token is null",
				Data: nil,
			}, nil
		}
		query.Where(comment.MemberId(strings.Join(value, "")))
	}

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

	count := int64(query.CountX(l.ctx))
	if in.GetPage() <= 0 {
		*in.Page = 1
	}
	if in.GetPageSize() <= 0 {
		*in.PageSize = 1
	}
	query.Limit(int(in.GetPageSize()))
	query.Offset(int((in.GetPage() - 1) * in.GetPageSize()))

	all, err := query.All(l.ctx)

	if err != nil {
		return &mms.CommentListResp{
			Code: -1,
			Msg:  err.Error(),
			Data: nil,
		}, nil
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

	return &mms.CommentListResp{
		Code: 0,
		Msg:  "成功",
		Data: &mms.CommentListData{
			Data:  infos,
			Total: count,
		},
	}, nil
}
