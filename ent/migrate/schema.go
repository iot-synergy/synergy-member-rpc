// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MmsCommentColumns holds the columns for the "mms_comment" table.
	MmsCommentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "title", Type: field.TypeString, Nullable: true, Comment: "标签", Default: "", SchemaType: map[string]string{"mysql": "varchar(256)"}},
		{Name: "content", Type: field.TypeString, Nullable: true, Comment: "正文", Default: "", SchemaType: map[string]string{"mysql": "varchar(2048)"}},
		{Name: "member_id", Type: field.TypeString, Comment: "会员id", SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "create_time", Type: field.TypeTime, Comment: "发布时间"},
		{Name: "update_time", Type: field.TypeTime, Comment: "更新时间"},
		{Name: "is_reply", Type: field.TypeBool, Comment: "是否已回复", Default: false},
	}
	// MmsCommentTable holds the schema information for the "mms_comment" table.
	MmsCommentTable = &schema.Table{
		Name:       "mms_comment",
		Columns:    MmsCommentColumns,
		PrimaryKey: []*schema.Column{MmsCommentColumns[0]},
	}
	// MmsMembersColumns holds the columns for the "mms_members" table.
	MmsMembersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Comment: "Status 1: normal 2: ban | 状态 1 正常 2 禁用", Default: 1},
		{Name: "forein_id", Type: field.TypeString, Unique: true, Comment: "Member's forein id | 外部ID", SchemaType: map[string]string{"mysql": "varchar(32)"}},
		{Name: "username", Type: field.TypeString, Unique: true, Comment: "Member's login name | 登录名"},
		{Name: "password", Type: field.TypeString, Comment: "Password | 密码"},
		{Name: "nickname", Type: field.TypeString, Unique: true, Comment: "Nickname | 昵称"},
		{Name: "mobile", Type: field.TypeString, Nullable: true, Comment: "Mobile number | 手机号"},
		{Name: "email", Type: field.TypeString, Nullable: true, Comment: "Email | 邮箱号"},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Comment: "Avatar | 头像路径", Default: "", SchemaType: map[string]string{"mysql": "varchar(512)"}},
		{Name: "wechat_open_id", Type: field.TypeString, Nullable: true, Comment: "Wechat Open ID | 微信 Open ID"},
		{Name: "expired_at", Type: field.TypeTime, Nullable: true, Comment: "Member expired time | 会员到期时间"},
		{Name: "rank_id", Type: field.TypeUint64, Nullable: true, Comment: "Member Rank ID | 会员等级ID", Default: 2},
	}
	// MmsMembersTable holds the schema information for the "mms_members" table.
	MmsMembersTable = &schema.Table{
		Name:       "mms_members",
		Columns:    MmsMembersColumns,
		PrimaryKey: []*schema.Column{MmsMembersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "mms_members_mms_ranks_ranks",
				Columns:    []*schema.Column{MmsMembersColumns[13]},
				RefColumns: []*schema.Column{MmsRanksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "member_username_email",
				Unique:  true,
				Columns: []*schema.Column{MmsMembersColumns[5], MmsMembersColumns[9]},
			},
			{
				Name:    "member_wechat_open_id",
				Unique:  true,
				Columns: []*schema.Column{MmsMembersColumns[11]},
			},
		},
	}
	// MmsRanksColumns holds the columns for the "mms_ranks" table.
	MmsRanksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "name", Type: field.TypeString},
		{Name: "code", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "remark", Type: field.TypeString},
	}
	// MmsRanksTable holds the schema information for the "mms_ranks" table.
	MmsRanksTable = &schema.Table{
		Name:       "mms_ranks",
		Columns:    MmsRanksColumns,
		PrimaryKey: []*schema.Column{MmsRanksColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "memberrank_code",
				Unique:  true,
				Columns: []*schema.Column{MmsRanksColumns[4]},
			},
		},
	}
	// MmsOauthProvidersColumns holds the columns for the "mms_oauth_providers" table.
	MmsOauthProvidersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "name", Type: field.TypeString, Unique: true, Comment: "The provider's name | 提供商名称"},
		{Name: "client_id", Type: field.TypeString, Comment: "The client id | 客户端 id"},
		{Name: "client_secret", Type: field.TypeString, Comment: "The client secret | 客户端密钥"},
		{Name: "redirect_url", Type: field.TypeString, Comment: "The redirect url | 跳转地址"},
		{Name: "scopes", Type: field.TypeString, Comment: "The scopes | 权限范围"},
		{Name: "auth_url", Type: field.TypeString, Comment: "The auth url of the provider | 认证地址"},
		{Name: "token_url", Type: field.TypeString, Comment: "The token url of the provider | 获取 token地址"},
		{Name: "auth_style", Type: field.TypeUint64, Comment: "The auth style, 0: auto detect 1: third party log in 2: log in with username and password"},
		{Name: "info_url", Type: field.TypeString, Comment: "The URL to request user information by token | 用户信息请求地址"},
	}
	// MmsOauthProvidersTable holds the schema information for the "mms_oauth_providers" table.
	MmsOauthProvidersTable = &schema.Table{
		Name:       "mms_oauth_providers",
		Columns:    MmsOauthProvidersColumns,
		PrimaryKey: []*schema.Column{MmsOauthProvidersColumns[0]},
	}
	// MmsReplyColumns holds the columns for the "mms_reply" table.
	MmsReplyColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "reply", Type: field.TypeString, Comment: "回复内容", SchemaType: map[string]string{"mysql": "varchar(2048)"}},
		{Name: "admin_id", Type: field.TypeString, Comment: "管理员id", SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "admin_name", Type: field.TypeString, Comment: "管理员名字", SchemaType: map[string]string{"mysql": "varchar(128)"}},
		{Name: "create_time", Type: field.TypeTime, Comment: "发布时间"},
		{Name: "update_time", Type: field.TypeTime, Comment: "更新时间"},
		{Name: "comment_id", Type: field.TypeUint64, Comment: "评论id"},
	}
	// MmsReplyTable holds the schema information for the "mms_reply" table.
	MmsReplyTable = &schema.Table{
		Name:       "mms_reply",
		Columns:    MmsReplyColumns,
		PrimaryKey: []*schema.Column{MmsReplyColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "mms_reply_mms_comment_replys",
				Columns:    []*schema.Column{MmsReplyColumns[8]},
				RefColumns: []*schema.Column{MmsCommentColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MmsTokensColumns holds the columns for the "mms_tokens" table.
	MmsTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Comment: "Status 1: normal 2: ban | 状态 1 正常 2 禁用", Default: 1},
		{Name: "uuid", Type: field.TypeUUID, Comment: " User's UUID | 用户的UUID"},
		{Name: "token", Type: field.TypeString, Comment: "Token string | Token 字符串"},
		{Name: "username", Type: field.TypeString, Comment: "Username | 用户名", Default: "unknown"},
		{Name: "source", Type: field.TypeString, Comment: "Log in source such as GitHub | Token 来源 （本地为core, 第三方如github等）"},
		{Name: "expired_at", Type: field.TypeTime, Comment: " Expire time | 过期时间"},
	}
	// MmsTokensTable holds the schema information for the "mms_tokens" table.
	MmsTokensTable = &schema.Table{
		Name:       "mms_tokens",
		Columns:    MmsTokensColumns,
		PrimaryKey: []*schema.Column{MmsTokensColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "mms_token_uuid",
				Unique:  false,
				Columns: []*schema.Column{MmsTokensColumns[4]},
			},
			{
				Name:    "mms_token_expired_at",
				Unique:  false,
				Columns: []*schema.Column{MmsTokensColumns[8]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MmsCommentTable,
		MmsMembersTable,
		MmsRanksTable,
		MmsOauthProvidersTable,
		MmsReplyTable,
		MmsTokensTable,
	}
)

func init() {
	MmsCommentTable.Annotation = &entsql.Annotation{
		Table: "mms_comment",
	}
	MmsMembersTable.ForeignKeys[0].RefTable = MmsRanksTable
	MmsMembersTable.Annotation = &entsql.Annotation{
		Table: "mms_members",
	}
	MmsRanksTable.Annotation = &entsql.Annotation{
		Table: "mms_ranks",
	}
	MmsOauthProvidersTable.Annotation = &entsql.Annotation{
		Table: "mms_oauth_providers",
	}
	MmsReplyTable.ForeignKeys[0].RefTable = MmsCommentTable
	MmsReplyTable.Annotation = &entsql.Annotation{
		Table: "mms_reply",
	}
	MmsTokensTable.Annotation = &entsql.Annotation{
		Table: "mms_tokens",
	}
}
