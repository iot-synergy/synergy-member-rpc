// Code generated by ent, DO NOT EDIT.

package comment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/iot-synergy/synergy-member-rpc/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldUpdatedAt, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldTitle, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldContent, v))
}

// MemberId applies equality check predicate on the "memberId" field. It's identical to MemberIdEQ.
func MemberId(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldMemberId, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldUpdateTime, v))
}

// IsReply applies equality check predicate on the "is_reply" field. It's identical to IsReplyEQ.
func IsReply(v bool) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldIsReply, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldUpdatedAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleIsNil applies the IsNil predicate on the "title" field.
func TitleIsNil() predicate.Comment {
	return predicate.Comment(sql.FieldIsNull(FieldTitle))
}

// TitleNotNil applies the NotNil predicate on the "title" field.
func TitleNotNil() predicate.Comment {
	return predicate.Comment(sql.FieldNotNull(FieldTitle))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldTitle, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasSuffix(FieldContent, v))
}

// ContentIsNil applies the IsNil predicate on the "content" field.
func ContentIsNil() predicate.Comment {
	return predicate.Comment(sql.FieldIsNull(FieldContent))
}

// ContentNotNil applies the NotNil predicate on the "content" field.
func ContentNotNil() predicate.Comment {
	return predicate.Comment(sql.FieldNotNull(FieldContent))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldContent, v))
}

// MemberIdEQ applies the EQ predicate on the "memberId" field.
func MemberIdEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldMemberId, v))
}

// MemberIdNEQ applies the NEQ predicate on the "memberId" field.
func MemberIdNEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldMemberId, v))
}

// MemberIdIn applies the In predicate on the "memberId" field.
func MemberIdIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldMemberId, vs...))
}

// MemberIdNotIn applies the NotIn predicate on the "memberId" field.
func MemberIdNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldMemberId, vs...))
}

// MemberIdGT applies the GT predicate on the "memberId" field.
func MemberIdGT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldMemberId, v))
}

// MemberIdGTE applies the GTE predicate on the "memberId" field.
func MemberIdGTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldMemberId, v))
}

// MemberIdLT applies the LT predicate on the "memberId" field.
func MemberIdLT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldMemberId, v))
}

// MemberIdLTE applies the LTE predicate on the "memberId" field.
func MemberIdLTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldMemberId, v))
}

// MemberIdContains applies the Contains predicate on the "memberId" field.
func MemberIdContains(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContains(FieldMemberId, v))
}

// MemberIdHasPrefix applies the HasPrefix predicate on the "memberId" field.
func MemberIdHasPrefix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasPrefix(FieldMemberId, v))
}

// MemberIdHasSuffix applies the HasSuffix predicate on the "memberId" field.
func MemberIdHasSuffix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasSuffix(FieldMemberId, v))
}

// MemberIdEqualFold applies the EqualFold predicate on the "memberId" field.
func MemberIdEqualFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldMemberId, v))
}

// MemberIdContainsFold applies the ContainsFold predicate on the "memberId" field.
func MemberIdContainsFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldMemberId, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldUpdateTime, v))
}

// IsReplyEQ applies the EQ predicate on the "is_reply" field.
func IsReplyEQ(v bool) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldIsReply, v))
}

// IsReplyNEQ applies the NEQ predicate on the "is_reply" field.
func IsReplyNEQ(v bool) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldIsReply, v))
}

// HasReplys applies the HasEdge predicate on the "replys" edge.
func HasReplys() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReplysTable, ReplysColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReplysWith applies the HasEdge predicate on the "replys" edge with a given conditions (other predicates).
func HasReplysWith(preds ...predicate.Reply) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := newReplysStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Comment) predicate.Comment {
	return predicate.Comment(sql.NotPredicates(p))
}
