package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/iot-synergy/synergy-common/orm/ent/mixins"
	"time"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			SchemaType(map[string]string{dialect.MySQL: "varchar(256)"}).
			Optional().
			Default("").
			Comment("标签").
			Annotations(entsql.WithComments(true)),
		field.String("content").
			SchemaType(map[string]string{dialect.MySQL: "varchar(2048)"}).
			Optional().
			Default("").
			Comment("正文").
			Annotations(entsql.WithComments(true)),
		field.String("memberId").NotEmpty().
			SchemaType(map[string]string{dialect.MySQL: "varchar(64)"}).
			Comment("会员id").
			Annotations(entsql.WithComments(true)),
		field.Time("create_time").
			Default(time.Now()).
			Comment("发布时间").
			Annotations(entsql.WithComments(true)),
		field.Time("update_time").
			Default(time.Now()).
			Comment("更新时间").
			Annotations(entsql.WithComments(true)),
		field.Bool("is_reply").
			Default(false).
			Comment("是否已回复").
			Annotations(entsql.WithComments(true)),
	}
}

func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("replys", Reply.Type),
	}
}
