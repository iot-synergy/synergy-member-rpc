package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/iot-synergy/synergy-common/orm/ent/mixins"
	"time"
)

// Reply holds the schema definition for the Reply entity.
type Reply struct {
	ent.Schema
}

// Fields of the Reply.
func (Reply) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("comment_id").
			Comment("评论id").
			Annotations(entsql.WithComments(true)),
		field.String("reply").
			SchemaType(map[string]string{dialect.MySQL: "varchar(2048)"}).
			Comment("回复内容").
			Annotations(entsql.WithComments(true)),
		field.String("adminId").NotEmpty().
			SchemaType(map[string]string{dialect.MySQL: "varchar(64)"}).
			Comment("管理员id").
			Annotations(entsql.WithComments(true)),
		field.String("adminName").
			SchemaType(map[string]string{dialect.MySQL: "varchar(128)"}).
			Comment("管理员名字").
			Annotations(entsql.WithComments(true)),
		field.Time("create_time").
			Default(time.Now()).
			Comment("发布时间").
			Annotations(entsql.WithComments(true)),
		field.Time("update_time").
			Default(time.Now()).
			Comment("更新时间").
			Annotations(entsql.WithComments(true)),
	}
}

func (r Reply) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

// Edges of the Reply.
func (Reply) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("comment", Comment.Type).
			Ref("replys").
			Unique().
			Field("comment_id").
			Required(),
	}
}

func (Reply) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "mms_reply"},
	}
}
