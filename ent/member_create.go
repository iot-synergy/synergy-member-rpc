// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/synergy-member-rpc/ent/member"
	"github.com/iot-synergy/synergy-member-rpc/ent/memberrank"
)

// MemberCreate is the builder for creating a Member entity.
type MemberCreate struct {
	config
	mutation *MemberMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mc *MemberCreate) SetCreatedAt(t time.Time) *MemberCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MemberCreate) SetNillableCreatedAt(t *time.Time) *MemberCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MemberCreate) SetUpdatedAt(t time.Time) *MemberCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MemberCreate) SetNillableUpdatedAt(t *time.Time) *MemberCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetStatus sets the "status" field.
func (mc *MemberCreate) SetStatus(u uint8) *MemberCreate {
	mc.mutation.SetStatus(u)
	return mc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mc *MemberCreate) SetNillableStatus(u *uint8) *MemberCreate {
	if u != nil {
		mc.SetStatus(*u)
	}
	return mc
}

// SetUsername sets the "username" field.
func (mc *MemberCreate) SetUsername(s string) *MemberCreate {
	mc.mutation.SetUsername(s)
	return mc
}

// SetPassword sets the "password" field.
func (mc *MemberCreate) SetPassword(s string) *MemberCreate {
	mc.mutation.SetPassword(s)
	return mc
}

// SetNickname sets the "nickname" field.
func (mc *MemberCreate) SetNickname(s string) *MemberCreate {
	mc.mutation.SetNickname(s)
	return mc
}

// SetRankID sets the "rank_id" field.
func (mc *MemberCreate) SetRankID(u uint64) *MemberCreate {
	mc.mutation.SetRankID(u)
	return mc
}

// SetNillableRankID sets the "rank_id" field if the given value is not nil.
func (mc *MemberCreate) SetNillableRankID(u *uint64) *MemberCreate {
	if u != nil {
		mc.SetRankID(*u)
	}
	return mc
}

// SetMobile sets the "mobile" field.
func (mc *MemberCreate) SetMobile(s string) *MemberCreate {
	mc.mutation.SetMobile(s)
	return mc
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (mc *MemberCreate) SetNillableMobile(s *string) *MemberCreate {
	if s != nil {
		mc.SetMobile(*s)
	}
	return mc
}

// SetEmail sets the "email" field.
func (mc *MemberCreate) SetEmail(s string) *MemberCreate {
	mc.mutation.SetEmail(s)
	return mc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (mc *MemberCreate) SetNillableEmail(s *string) *MemberCreate {
	if s != nil {
		mc.SetEmail(*s)
	}
	return mc
}

// SetAvatar sets the "avatar" field.
func (mc *MemberCreate) SetAvatar(s string) *MemberCreate {
	mc.mutation.SetAvatar(s)
	return mc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (mc *MemberCreate) SetNillableAvatar(s *string) *MemberCreate {
	if s != nil {
		mc.SetAvatar(*s)
	}
	return mc
}

// SetWechatOpenID sets the "wechat_open_id" field.
func (mc *MemberCreate) SetWechatOpenID(s string) *MemberCreate {
	mc.mutation.SetWechatOpenID(s)
	return mc
}

// SetNillableWechatOpenID sets the "wechat_open_id" field if the given value is not nil.
func (mc *MemberCreate) SetNillableWechatOpenID(s *string) *MemberCreate {
	if s != nil {
		mc.SetWechatOpenID(*s)
	}
	return mc
}

// SetExpiredAt sets the "expired_at" field.
func (mc *MemberCreate) SetExpiredAt(t time.Time) *MemberCreate {
	mc.mutation.SetExpiredAt(t)
	return mc
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (mc *MemberCreate) SetNillableExpiredAt(t *time.Time) *MemberCreate {
	if t != nil {
		mc.SetExpiredAt(*t)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MemberCreate) SetID(u uuid.UUID) *MemberCreate {
	mc.mutation.SetID(u)
	return mc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mc *MemberCreate) SetNillableID(u *uuid.UUID) *MemberCreate {
	if u != nil {
		mc.SetID(*u)
	}
	return mc
}

// SetRanksID sets the "ranks" edge to the MemberRank entity by ID.
func (mc *MemberCreate) SetRanksID(id uint64) *MemberCreate {
	mc.mutation.SetRanksID(id)
	return mc
}

// SetNillableRanksID sets the "ranks" edge to the MemberRank entity by ID if the given value is not nil.
func (mc *MemberCreate) SetNillableRanksID(id *uint64) *MemberCreate {
	if id != nil {
		mc = mc.SetRanksID(*id)
	}
	return mc
}

// SetRanks sets the "ranks" edge to the MemberRank entity.
func (mc *MemberCreate) SetRanks(m *MemberRank) *MemberCreate {
	return mc.SetRanksID(m.ID)
}

// Mutation returns the MemberMutation object of the builder.
func (mc *MemberCreate) Mutation() *MemberMutation {
	return mc.mutation
}

// Save creates the Member in the database.
func (mc *MemberCreate) Save(ctx context.Context) (*Member, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MemberCreate) SaveX(ctx context.Context) *Member {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MemberCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MemberCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MemberCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := member.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := member.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.Status(); !ok {
		v := member.DefaultStatus
		mc.mutation.SetStatus(v)
	}
	if _, ok := mc.mutation.RankID(); !ok {
		v := member.DefaultRankID
		mc.mutation.SetRankID(v)
	}
	if _, ok := mc.mutation.Avatar(); !ok {
		v := member.DefaultAvatar
		mc.mutation.SetAvatar(v)
	}
	if _, ok := mc.mutation.ExpiredAt(); !ok {
		v := member.DefaultExpiredAt
		mc.mutation.SetExpiredAt(v)
	}
	if _, ok := mc.mutation.ID(); !ok {
		v := member.DefaultID()
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MemberCreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Member.created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Member.updated_at"`)}
	}
	if _, ok := mc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "Member.username"`)}
	}
	if _, ok := mc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Member.password"`)}
	}
	if _, ok := mc.mutation.Nickname(); !ok {
		return &ValidationError{Name: "nickname", err: errors.New(`ent: missing required field "Member.nickname"`)}
	}
	return nil
}

func (mc *MemberCreate) sqlSave(ctx context.Context) (*Member, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MemberCreate) createSpec() (*Member, *sqlgraph.CreateSpec) {
	var (
		_node = &Member{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(member.Table, sqlgraph.NewFieldSpec(member.FieldID, field.TypeUUID))
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(member.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(member.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.Status(); ok {
		_spec.SetField(member.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := mc.mutation.Username(); ok {
		_spec.SetField(member.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := mc.mutation.Password(); ok {
		_spec.SetField(member.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := mc.mutation.Nickname(); ok {
		_spec.SetField(member.FieldNickname, field.TypeString, value)
		_node.Nickname = value
	}
	if value, ok := mc.mutation.Mobile(); ok {
		_spec.SetField(member.FieldMobile, field.TypeString, value)
		_node.Mobile = value
	}
	if value, ok := mc.mutation.Email(); ok {
		_spec.SetField(member.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := mc.mutation.Avatar(); ok {
		_spec.SetField(member.FieldAvatar, field.TypeString, value)
		_node.Avatar = value
	}
	if value, ok := mc.mutation.WechatOpenID(); ok {
		_spec.SetField(member.FieldWechatOpenID, field.TypeString, value)
		_node.WechatOpenID = value
	}
	if value, ok := mc.mutation.ExpiredAt(); ok {
		_spec.SetField(member.FieldExpiredAt, field.TypeTime, value)
		_node.ExpiredAt = value
	}
	if nodes := mc.mutation.RanksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   member.RanksTable,
			Columns: []string{member.RanksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memberrank.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RankID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MemberCreateBulk is the builder for creating many Member entities in bulk.
type MemberCreateBulk struct {
	config
	err      error
	builders []*MemberCreate
}

// Save creates the Member entities in the database.
func (mcb *MemberCreateBulk) Save(ctx context.Context) ([]*Member, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Member, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MemberMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MemberCreateBulk) SaveX(ctx context.Context) []*Member {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MemberCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MemberCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
