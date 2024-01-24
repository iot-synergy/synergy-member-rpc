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

// MemberRankCreate is the builder for creating a MemberRank entity.
type MemberRankCreate struct {
	config
	mutation *MemberRankMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mrc *MemberRankCreate) SetCreatedAt(t time.Time) *MemberRankCreate {
	mrc.mutation.SetCreatedAt(t)
	return mrc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mrc *MemberRankCreate) SetNillableCreatedAt(t *time.Time) *MemberRankCreate {
	if t != nil {
		mrc.SetCreatedAt(*t)
	}
	return mrc
}

// SetUpdatedAt sets the "updated_at" field.
func (mrc *MemberRankCreate) SetUpdatedAt(t time.Time) *MemberRankCreate {
	mrc.mutation.SetUpdatedAt(t)
	return mrc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mrc *MemberRankCreate) SetNillableUpdatedAt(t *time.Time) *MemberRankCreate {
	if t != nil {
		mrc.SetUpdatedAt(*t)
	}
	return mrc
}

// SetName sets the "name" field.
func (mrc *MemberRankCreate) SetName(s string) *MemberRankCreate {
	mrc.mutation.SetName(s)
	return mrc
}

// SetCode sets the "code" field.
func (mrc *MemberRankCreate) SetCode(s string) *MemberRankCreate {
	mrc.mutation.SetCode(s)
	return mrc
}

// SetDescription sets the "description" field.
func (mrc *MemberRankCreate) SetDescription(s string) *MemberRankCreate {
	mrc.mutation.SetDescription(s)
	return mrc
}

// SetRemark sets the "remark" field.
func (mrc *MemberRankCreate) SetRemark(s string) *MemberRankCreate {
	mrc.mutation.SetRemark(s)
	return mrc
}

// SetID sets the "id" field.
func (mrc *MemberRankCreate) SetID(u uint64) *MemberRankCreate {
	mrc.mutation.SetID(u)
	return mrc
}

// AddMemberIDs adds the "members" edge to the Member entity by IDs.
func (mrc *MemberRankCreate) AddMemberIDs(ids ...uuid.UUID) *MemberRankCreate {
	mrc.mutation.AddMemberIDs(ids...)
	return mrc
}

// AddMembers adds the "members" edges to the Member entity.
func (mrc *MemberRankCreate) AddMembers(m ...*Member) *MemberRankCreate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mrc.AddMemberIDs(ids...)
}

// Mutation returns the MemberRankMutation object of the builder.
func (mrc *MemberRankCreate) Mutation() *MemberRankMutation {
	return mrc.mutation
}

// Save creates the MemberRank in the database.
func (mrc *MemberRankCreate) Save(ctx context.Context) (*MemberRank, error) {
	mrc.defaults()
	return withHooks(ctx, mrc.sqlSave, mrc.mutation, mrc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mrc *MemberRankCreate) SaveX(ctx context.Context) *MemberRank {
	v, err := mrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mrc *MemberRankCreate) Exec(ctx context.Context) error {
	_, err := mrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mrc *MemberRankCreate) ExecX(ctx context.Context) {
	if err := mrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mrc *MemberRankCreate) defaults() {
	if _, ok := mrc.mutation.CreatedAt(); !ok {
		v := memberrank.DefaultCreatedAt()
		mrc.mutation.SetCreatedAt(v)
	}
	if _, ok := mrc.mutation.UpdatedAt(); !ok {
		v := memberrank.DefaultUpdatedAt()
		mrc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mrc *MemberRankCreate) check() error {
	if _, ok := mrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MemberRank.created_at"`)}
	}
	if _, ok := mrc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "MemberRank.updated_at"`)}
	}
	if _, ok := mrc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "MemberRank.name"`)}
	}
	if _, ok := mrc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "MemberRank.code"`)}
	}
	if _, ok := mrc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "MemberRank.description"`)}
	}
	if _, ok := mrc.mutation.Remark(); !ok {
		return &ValidationError{Name: "remark", err: errors.New(`ent: missing required field "MemberRank.remark"`)}
	}
	return nil
}

func (mrc *MemberRankCreate) sqlSave(ctx context.Context) (*MemberRank, error) {
	if err := mrc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	mrc.mutation.id = &_node.ID
	mrc.mutation.done = true
	return _node, nil
}

func (mrc *MemberRankCreate) createSpec() (*MemberRank, *sqlgraph.CreateSpec) {
	var (
		_node = &MemberRank{config: mrc.config}
		_spec = sqlgraph.NewCreateSpec(memberrank.Table, sqlgraph.NewFieldSpec(memberrank.FieldID, field.TypeUint64))
	)
	if id, ok := mrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mrc.mutation.CreatedAt(); ok {
		_spec.SetField(memberrank.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mrc.mutation.UpdatedAt(); ok {
		_spec.SetField(memberrank.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mrc.mutation.Name(); ok {
		_spec.SetField(memberrank.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mrc.mutation.Code(); ok {
		_spec.SetField(memberrank.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := mrc.mutation.Description(); ok {
		_spec.SetField(memberrank.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := mrc.mutation.Remark(); ok {
		_spec.SetField(memberrank.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if nodes := mrc.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   memberrank.MembersTable,
			Columns: []string{memberrank.MembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MemberRankCreateBulk is the builder for creating many MemberRank entities in bulk.
type MemberRankCreateBulk struct {
	config
	err      error
	builders []*MemberRankCreate
}

// Save creates the MemberRank entities in the database.
func (mrcb *MemberRankCreateBulk) Save(ctx context.Context) ([]*MemberRank, error) {
	if mrcb.err != nil {
		return nil, mrcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mrcb.builders))
	nodes := make([]*MemberRank, len(mrcb.builders))
	mutators := make([]Mutator, len(mrcb.builders))
	for i := range mrcb.builders {
		func(i int, root context.Context) {
			builder := mrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MemberRankMutation)
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
					_, err = mutators[i+1].Mutate(root, mrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mrcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, mrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mrcb *MemberRankCreateBulk) SaveX(ctx context.Context) []*MemberRank {
	v, err := mrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mrcb *MemberRankCreateBulk) Exec(ctx context.Context) error {
	_, err := mrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mrcb *MemberRankCreateBulk) ExecX(ctx context.Context) {
	if err := mrcb.Exec(ctx); err != nil {
		panic(err)
	}
}
