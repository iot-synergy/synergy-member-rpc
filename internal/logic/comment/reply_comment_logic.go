package comment

import (
	"context"
	"errors"
	"github.com/iot-synergy/synergy-member-rpc/ent/comment"
	"github.com/iot-synergy/synergy-member-rpc/internal/utils/dberrorhandler"
	"strconv"
	"time"

	"github.com/iot-synergy/synergy-member-rpc/internal/svc"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReplyCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReplyCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReplyCommentLogic {
	return &ReplyCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: comment
func (l *ReplyCommentLogic) ReplyComment(in *mms.ReplyInfo) (*mms.BaseResp, error) {
	// todo: add your logic here and delete this line

	get := l.svcCtx.DB.Comment.GetX(l.ctx, uint64(*in.CommentId))
	if get == nil {
		return nil, errors.New("commentId no data")
	}

	if *in.Reply == "" {
		return nil, errors.New("reply is null")
	}

	if len(*in.Reply) > 2048 {
		return nil, errors.New("content length exceeds 2048")
	}

	create := l.svcCtx.DB.Reply.Create().
		SetCommentID(uint64(in.GetCommentId())).
		SetReply(in.GetReply()).
		SetAdminId(in.GetAdminId()).
		SetAdminName(in.GetAdminName())

	save, err := create.Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	if get.IsReply == false {
		l.svcCtx.DB.Comment.Update().SetUpdateTime(time.Now()).SetIsReply(true).Where(comment.ID(get.ID)).SaveX(l.ctx)
	}
	return &mms.BaseResp{Msg: strconv.FormatUint(save.ID, 10)}, nil
}
