package comment

import (
	"context"
	"errors"
	"github.com/iot-synergy/synergy-member-rpc/ent/comment"
	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"
	"google.golang.org/grpc/metadata"
	"strings"

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
func (l *MemberGetCommentListLogic) MemberGetCommentList(in *mms.CommentListReq) (*mms.CommentList, error) {
	// todo: add your logic here and delete this line
	query := l.svcCtx.DB.Comment.Query()

	incomingContext, ok := metadata.FromIncomingContext(l.ctx)
	if ok {
		value := incomingContext.Get("gateway-firebaseid")
		if value == nil || len(value) == 0 {
			return nil, errors.New("member token is null")
		}
		query.Where(comment.MemberId(strings.Join(value, "")))
	}

	if in.GetIsReply() == 0 {
		query.Where(comment.IsReply(false))
	} else if in.GetIsReply() == 1 {
		query.Where(comment.IsReply(true))
	}

	query.Limit(int(in.GetPageSize()))
	query.Offset(int((in.GetPageNo() - 1) * in.GetPageSize()))

	all, err := query.All(l.ctx)

	query.Select("title")

	if err != nil {
		return nil, err
	}

	infos := make([]*mms.CommentInfo, 0)
	for _, comm := range all {
		info := mms.CommentInfo{
			Title: &comm.Title,
		}
		infos = append(infos, &info)
	}

	return &mms.CommentList{
		Titles: infos,
	}, nil
}
