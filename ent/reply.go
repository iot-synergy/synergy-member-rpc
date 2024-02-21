// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/iot-synergy/synergy-member-rpc/ent/comment"
	"github.com/iot-synergy/synergy-member-rpc/ent/reply"
)

// Reply is the model entity for the Reply schema.
type Reply struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 评论id
	CommentID uint64 `json:"comment_id,omitempty"`
	// 回复内容
	Reply string `json:"reply,omitempty"`
	// 管理员id
	AdminId int64 `json:"adminId,omitempty"`
	// 管理员名字
	AdminName string `json:"adminName,omitempty"`
	// 发布时间
	CreateTime time.Time `json:"create_time,omitempty"`
	// 更新时间
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReplyQuery when eager-loading is set.
	Edges        ReplyEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ReplyEdges holds the relations/edges for other nodes in the graph.
type ReplyEdges struct {
	// Comment holds the value of the comment edge.
	Comment *Comment `json:"comment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CommentOrErr returns the Comment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReplyEdges) CommentOrErr() (*Comment, error) {
	if e.loadedTypes[0] {
		if e.Comment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: comment.Label}
		}
		return e.Comment, nil
	}
	return nil, &NotLoadedError{edge: "comment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Reply) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case reply.FieldID, reply.FieldCommentID, reply.FieldAdminId:
			values[i] = new(sql.NullInt64)
		case reply.FieldReply, reply.FieldAdminName:
			values[i] = new(sql.NullString)
		case reply.FieldCreatedAt, reply.FieldUpdatedAt, reply.FieldCreateTime, reply.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Reply fields.
func (r *Reply) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case reply.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = uint64(value.Int64)
		case reply.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case reply.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case reply.FieldCommentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field comment_id", values[i])
			} else if value.Valid {
				r.CommentID = uint64(value.Int64)
			}
		case reply.FieldReply:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reply", values[i])
			} else if value.Valid {
				r.Reply = value.String
			}
		case reply.FieldAdminId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field adminId", values[i])
			} else if value.Valid {
				r.AdminId = value.Int64
			}
		case reply.FieldAdminName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field adminName", values[i])
			} else if value.Valid {
				r.AdminName = value.String
			}
		case reply.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				r.CreateTime = value.Time
			}
		case reply.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				r.UpdateTime = value.Time
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Reply.
// This includes values selected through modifiers, order, etc.
func (r *Reply) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryComment queries the "comment" edge of the Reply entity.
func (r *Reply) QueryComment() *CommentQuery {
	return NewReplyClient(r.config).QueryComment(r)
}

// Update returns a builder for updating this Reply.
// Note that you need to call Reply.Unwrap() before calling this method if this Reply
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Reply) Update() *ReplyUpdateOne {
	return NewReplyClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Reply entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Reply) Unwrap() *Reply {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Reply is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Reply) String() string {
	var builder strings.Builder
	builder.WriteString("Reply(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("comment_id=")
	builder.WriteString(fmt.Sprintf("%v", r.CommentID))
	builder.WriteString(", ")
	builder.WriteString("reply=")
	builder.WriteString(r.Reply)
	builder.WriteString(", ")
	builder.WriteString("adminId=")
	builder.WriteString(fmt.Sprintf("%v", r.AdminId))
	builder.WriteString(", ")
	builder.WriteString("adminName=")
	builder.WriteString(r.AdminName)
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(r.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(r.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Replies is a parsable slice of Reply.
type Replies []*Reply
