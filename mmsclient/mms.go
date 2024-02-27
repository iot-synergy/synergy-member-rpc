// Code generated by goctl. DO NOT EDIT.
// Source: mms.proto

package mmsclient

import (
	"context"

	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseIDResp            = mms.BaseIDResp
	BaseResp              = mms.BaseResp
	BaseUUIDResp          = mms.BaseUUIDResp
	CallbackReq           = mms.CallbackReq
	CommentIdReq          = mms.CommentIdReq
	CommentInfo           = mms.CommentInfo
	CommentList           = mms.CommentList
	CommentListReq        = mms.CommentListReq
	Empty                 = mms.Empty
	IDReq                 = mms.IDReq
	IDsReq                = mms.IDsReq
	MemberInfo            = mms.MemberInfo
	MemberListReq         = mms.MemberListReq
	MemberListResp        = mms.MemberListResp
	MemberLoginResp       = mms.MemberLoginResp
	MemberRankInfo        = mms.MemberRankInfo
	MemberRankListReq     = mms.MemberRankListReq
	MemberRankListResp    = mms.MemberRankListResp
	MemberRegisterReq     = mms.MemberRegisterReq
	OauthLoginReq         = mms.OauthLoginReq
	OauthProviderInfo     = mms.OauthProviderInfo
	OauthProviderListReq  = mms.OauthProviderListReq
	OauthProviderListResp = mms.OauthProviderListResp
	OauthRedirectResp     = mms.OauthRedirectResp
	PageInfoReq           = mms.PageInfoReq
	ReplyInfo             = mms.ReplyInfo
	ReplyList             = mms.ReplyList
	ReplyReq              = mms.ReplyReq
	TokenInfo             = mms.TokenInfo
	TokenListReq          = mms.TokenListReq
	TokenListResp         = mms.TokenListResp
	UUIDReq               = mms.UUIDReq
	UUIDsReq              = mms.UUIDsReq
	UsernameReq           = mms.UsernameReq

	Mms interface {
		// group: base
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
		// Member management
		CreateMember(ctx context.Context, in *MemberInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error)
		// group: member
		RegisterMember(ctx context.Context, in *MemberInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error)
		// group: member
		UpdateMember(ctx context.Context, in *MemberInfo, opts ...grpc.CallOption) (*BaseResp, error)
		// group: member
		GetMemberList(ctx context.Context, in *MemberListReq, opts ...grpc.CallOption) (*MemberListResp, error)
		// group: member
		DeleteMember(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// group: member
		GetMemberById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*MemberInfo, error)
		// group: member
		GetMemberByUsername(ctx context.Context, in *UsernameReq, opts ...grpc.CallOption) (*MemberInfo, error)
		// group: member
		GetMember(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MemberInfo, error)
		// group: member
		UpdateMember2(ctx context.Context, in *MemberInfo, opts ...grpc.CallOption) (*BaseResp, error)
		// MemberRank management
		CreateMemberRank(ctx context.Context, in *MemberRankInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		// group: memberrank
		UpdateMemberRank(ctx context.Context, in *MemberRankInfo, opts ...grpc.CallOption) (*BaseResp, error)
		// group: memberrank
		GetMemberRankList(ctx context.Context, in *MemberRankListReq, opts ...grpc.CallOption) (*MemberRankListResp, error)
		// group: memberrank
		GetMemberRankById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*MemberRankInfo, error)
		// group: memberrank
		DeleteMemberRank(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// OauthProvider management
		CreateOauthProvider(ctx context.Context, in *OauthProviderInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		// group: oauthprovider
		UpdateOauthProvider(ctx context.Context, in *OauthProviderInfo, opts ...grpc.CallOption) (*BaseResp, error)
		// group: oauthprovider
		GetOauthProviderList(ctx context.Context, in *OauthProviderListReq, opts ...grpc.CallOption) (*OauthProviderListResp, error)
		// group: oauthprovider
		GetOauthProviderById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*OauthProviderInfo, error)
		// group: oauthprovider
		DeleteOauthProvider(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// group: oauthprovider
		OauthLogin(ctx context.Context, in *OauthLoginReq, opts ...grpc.CallOption) (*OauthRedirectResp, error)
		// group: oauthprovider
		OauthCallback(ctx context.Context, in *CallbackReq, opts ...grpc.CallOption) (*MemberInfo, error)
		// group: oauthprovider
		WechatMiniProgramLogin(ctx context.Context, in *OauthLoginReq, opts ...grpc.CallOption) (*BaseResp, error)
		// Token management
		CreateToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error)
		// group: token
		DeleteToken(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// group: token
		GetTokenList(ctx context.Context, in *TokenListReq, opts ...grpc.CallOption) (*TokenListResp, error)
		// group: token
		GetTokenById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*TokenInfo, error)
		// group: token
		BlockUserAllToken(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*BaseResp, error)
		// group: token
		UpdateToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*BaseResp, error)
		// Comment management
		MemberComment(ctx context.Context, in *CommentInfo, opts ...grpc.CallOption) (*BaseResp, error)
		// group: comment
		MemberGetCommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentList, error)
		// group: comment
		MemberGetComment(ctx context.Context, in *CommentIdReq, opts ...grpc.CallOption) (*CommentInfo, error)
		// group: comment
		ReplyComment(ctx context.Context, in *ReplyInfo, opts ...grpc.CallOption) (*BaseResp, error)
		// group: comment
		AdminGetCommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentList, error)
		// group: comment
		AdminGetComment(ctx context.Context, in *CommentIdReq, opts ...grpc.CallOption) (*CommentInfo, error)
		// group: comment
		AdminGetReplyList(ctx context.Context, in *ReplyReq, opts ...grpc.CallOption) (*ReplyList, error)
	}

	defaultMms struct {
		cli zrpc.Client
	}
)

func NewMms(cli zrpc.Client) Mms {
	return &defaultMms{
		cli: cli,
	}
}

// group: base
func (m *defaultMms) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.InitDatabase(ctx, in, opts...)
}

// Member management
func (m *defaultMms) CreateMember(ctx context.Context, in *MemberInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.CreateMember(ctx, in, opts...)
}

// group: member
func (m *defaultMms) RegisterMember(ctx context.Context, in *MemberInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.RegisterMember(ctx, in, opts...)
}

// group: member
func (m *defaultMms) UpdateMember(ctx context.Context, in *MemberInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.UpdateMember(ctx, in, opts...)
}

// group: member
func (m *defaultMms) GetMemberList(ctx context.Context, in *MemberListReq, opts ...grpc.CallOption) (*MemberListResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.GetMemberList(ctx, in, opts...)
}

// group: member
func (m *defaultMms) DeleteMember(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.DeleteMember(ctx, in, opts...)
}

// group: member
func (m *defaultMms) GetMemberById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*MemberInfo, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.GetMemberById(ctx, in, opts...)
}

// group: member
func (m *defaultMms) GetMemberByUsername(ctx context.Context, in *UsernameReq, opts ...grpc.CallOption) (*MemberInfo, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.GetMemberByUsername(ctx, in, opts...)
}

// group: member
func (m *defaultMms) GetMember(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MemberInfo, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.GetMember(ctx, in, opts...)
}

// group: member
func (m *defaultMms) UpdateMember2(ctx context.Context, in *MemberInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.UpdateMember2(ctx, in, opts...)
}

// MemberRank management
func (m *defaultMms) CreateMemberRank(ctx context.Context, in *MemberRankInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.CreateMemberRank(ctx, in, opts...)
}

// group: memberrank
func (m *defaultMms) UpdateMemberRank(ctx context.Context, in *MemberRankInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.UpdateMemberRank(ctx, in, opts...)
}

// group: memberrank
func (m *defaultMms) GetMemberRankList(ctx context.Context, in *MemberRankListReq, opts ...grpc.CallOption) (*MemberRankListResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.GetMemberRankList(ctx, in, opts...)
}

// group: memberrank
func (m *defaultMms) GetMemberRankById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*MemberRankInfo, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.GetMemberRankById(ctx, in, opts...)
}

// group: memberrank
func (m *defaultMms) DeleteMemberRank(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.DeleteMemberRank(ctx, in, opts...)
}

// OauthProvider management
func (m *defaultMms) CreateOauthProvider(ctx context.Context, in *OauthProviderInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.CreateOauthProvider(ctx, in, opts...)
}

// group: oauthprovider
func (m *defaultMms) UpdateOauthProvider(ctx context.Context, in *OauthProviderInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.UpdateOauthProvider(ctx, in, opts...)
}

// group: oauthprovider
func (m *defaultMms) GetOauthProviderList(ctx context.Context, in *OauthProviderListReq, opts ...grpc.CallOption) (*OauthProviderListResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.GetOauthProviderList(ctx, in, opts...)
}

// group: oauthprovider
func (m *defaultMms) GetOauthProviderById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*OauthProviderInfo, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.GetOauthProviderById(ctx, in, opts...)
}

// group: oauthprovider
func (m *defaultMms) DeleteOauthProvider(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.DeleteOauthProvider(ctx, in, opts...)
}

// group: oauthprovider
func (m *defaultMms) OauthLogin(ctx context.Context, in *OauthLoginReq, opts ...grpc.CallOption) (*OauthRedirectResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.OauthLogin(ctx, in, opts...)
}

// group: oauthprovider
func (m *defaultMms) OauthCallback(ctx context.Context, in *CallbackReq, opts ...grpc.CallOption) (*MemberInfo, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.OauthCallback(ctx, in, opts...)
}

// group: oauthprovider
func (m *defaultMms) WechatMiniProgramLogin(ctx context.Context, in *OauthLoginReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.WechatMiniProgramLogin(ctx, in, opts...)
}

// Token management
func (m *defaultMms) CreateToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.CreateToken(ctx, in, opts...)
}

// group: token
func (m *defaultMms) DeleteToken(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.DeleteToken(ctx, in, opts...)
}

// group: token
func (m *defaultMms) GetTokenList(ctx context.Context, in *TokenListReq, opts ...grpc.CallOption) (*TokenListResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.GetTokenList(ctx, in, opts...)
}

// group: token
func (m *defaultMms) GetTokenById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*TokenInfo, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.GetTokenById(ctx, in, opts...)
}

// group: token
func (m *defaultMms) BlockUserAllToken(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.BlockUserAllToken(ctx, in, opts...)
}

// group: token
func (m *defaultMms) UpdateToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.UpdateToken(ctx, in, opts...)
}

// Comment management
func (m *defaultMms) MemberComment(ctx context.Context, in *CommentInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.MemberComment(ctx, in, opts...)
}

// group: comment
func (m *defaultMms) MemberGetCommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentList, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.MemberGetCommentList(ctx, in, opts...)
}

// group: comment
func (m *defaultMms) MemberGetComment(ctx context.Context, in *CommentIdReq, opts ...grpc.CallOption) (*CommentInfo, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.MemberGetComment(ctx, in, opts...)
}

// group: comment
func (m *defaultMms) ReplyComment(ctx context.Context, in *ReplyInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.ReplyComment(ctx, in, opts...)
}

// group: comment
func (m *defaultMms) AdminGetCommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentList, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.AdminGetCommentList(ctx, in, opts...)
}

// group: comment
func (m *defaultMms) AdminGetComment(ctx context.Context, in *CommentIdReq, opts ...grpc.CallOption) (*CommentInfo, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.AdminGetComment(ctx, in, opts...)
}

// group: comment
func (m *defaultMms) AdminGetReplyList(ctx context.Context, in *ReplyReq, opts ...grpc.CallOption) (*ReplyList, error) {
	client := mms.NewMmsClient(m.cli.Conn())
	return client.AdminGetReplyList(ctx, in, opts...)
}
