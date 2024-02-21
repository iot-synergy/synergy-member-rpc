package comment

import (
	"context"
	"errors"
	"github.com/iot-synergy/synergy-member-rpc/internal/utils/dberrorhandler"
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
func (l *MemberCommentLogic) MemberComment(in *mms.CommentInfo) (*mms.BaseResp, error) {
	// todo: add your logic here and delete this line
	query := l.svcCtx.DB.Comment.Create().
		SetTitle(*in.Title).
		SetContent(*in.Content).
		SetCreateTime(time.Now()).
		SetUpdateTime(time.Now())

	incomingContext, ok := metadata.FromIncomingContext(l.ctx)

	if ok {
		value := incomingContext.Get("gateway-firebaseid")
		if value == nil || len(value) == 0 {
			return nil, errors.New("member token is null")
		}
		query.SetMemberId(strings.Join(value, ""))
	}

	save, err := query.Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &mms.BaseResp{Msg: strconv.FormatUint(save.ID, 10)}, nil
}
