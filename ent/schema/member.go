package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/iot-synergy/synergy-common/orm/ent/mixins"
)

type Member struct {
	ent.Schema
}

func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.String("forein_id").Unique().
			SchemaType(map[string]string{dialect.MySQL: "varchar(32)"}).
			Comment("Member's forein id | 外部ID").
			Annotations(entsql.WithComments(true)),
		field.String("username").Unique().
			Comment("Member's login name | 登录名").
			Annotations(entsql.WithComments(true)),
		field.String("password").
			Comment("Password | 密码").
			Annotations(entsql.WithComments(true)),
		field.String("nickname").Unique().
			Comment("Nickname | 昵称").
			Annotations(entsql.WithComments(true)),
		field.Uint64("rank_id").Optional().Default(2).
			Comment("Member Rank ID | 会员等级ID").
			Annotations(entsql.WithComments(true)),
		field.String("mobile").Optional().
			Comment("Mobile number | 手机号").
			Annotations(entsql.WithComments(true)),
		field.String("email").Optional().
			Comment("Email | 邮箱号").
			Annotations(entsql.WithComments(true)),
		field.String("avatar").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Default("").
			Comment("Avatar | 头像路径").
			Annotations(entsql.WithComments(true)),
		field.String("wechat_open_id").Optional().
			Comment("Wechat Open ID | 微信 Open ID").
			Annotations(entsql.WithComments(true)),
		field.Time("expired_at").Optional().
			Comment("Member expired time | 会员到期时间").
			Default(time.Date(2036, 12, 31, 23, 59, 59, 0, time.Local)).
			Annotations(entsql.WithComments(true)),
	}
}

func (Member) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
	}
}

func (Member) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ranks", MemberRank.Type).Unique().Field("rank_id"),
	}
}

func (Member) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username", "email").
			Unique(),
		index.Fields("wechat_open_id").Unique(),
	}
}

func (Member) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "mms_members"}, // mms means member management service
	}
}
