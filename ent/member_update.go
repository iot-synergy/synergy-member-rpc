// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-member-rpc/ent/member"
	"github.com/suyuan32/simple-admin-member-rpc/ent/memberrank"
	"github.com/suyuan32/simple-admin-member-rpc/ent/predicate"
)

// MemberUpdate is the builder for updating Member entities.
type MemberUpdate struct {
	config
	hooks    []Hook
	mutation *MemberMutation
}

// Where appends a list predicates to the MemberUpdate builder.
func (mu *MemberUpdate) Where(ps ...predicate.Member) *MemberUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MemberUpdate) SetUpdatedAt(t time.Time) *MemberUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetStatus sets the "status" field.
func (mu *MemberUpdate) SetStatus(u uint8) *MemberUpdate {
	mu.mutation.ResetStatus()
	mu.mutation.SetStatus(u)
	return mu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableStatus(u *uint8) *MemberUpdate {
	if u != nil {
		mu.SetStatus(*u)
	}
	return mu
}

// AddStatus adds u to the "status" field.
func (mu *MemberUpdate) AddStatus(u int8) *MemberUpdate {
	mu.mutation.AddStatus(u)
	return mu
}

// ClearStatus clears the value of the "status" field.
func (mu *MemberUpdate) ClearStatus() *MemberUpdate {
	mu.mutation.ClearStatus()
	return mu
}

// SetUsername sets the "username" field.
func (mu *MemberUpdate) SetUsername(s string) *MemberUpdate {
	mu.mutation.SetUsername(s)
	return mu
}

// SetPassword sets the "password" field.
func (mu *MemberUpdate) SetPassword(s string) *MemberUpdate {
	mu.mutation.SetPassword(s)
	return mu
}

// SetNickname sets the "nickname" field.
func (mu *MemberUpdate) SetNickname(s string) *MemberUpdate {
	mu.mutation.SetNickname(s)
	return mu
}

// SetRankID sets the "rank_id" field.
func (mu *MemberUpdate) SetRankID(u uint64) *MemberUpdate {
	mu.mutation.SetRankID(u)
	return mu
}

// SetNillableRankID sets the "rank_id" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableRankID(u *uint64) *MemberUpdate {
	if u != nil {
		mu.SetRankID(*u)
	}
	return mu
}

// ClearRankID clears the value of the "rank_id" field.
func (mu *MemberUpdate) ClearRankID() *MemberUpdate {
	mu.mutation.ClearRankID()
	return mu
}

// SetMobile sets the "mobile" field.
func (mu *MemberUpdate) SetMobile(s string) *MemberUpdate {
	mu.mutation.SetMobile(s)
	return mu
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableMobile(s *string) *MemberUpdate {
	if s != nil {
		mu.SetMobile(*s)
	}
	return mu
}

// ClearMobile clears the value of the "mobile" field.
func (mu *MemberUpdate) ClearMobile() *MemberUpdate {
	mu.mutation.ClearMobile()
	return mu
}

// SetEmail sets the "email" field.
func (mu *MemberUpdate) SetEmail(s string) *MemberUpdate {
	mu.mutation.SetEmail(s)
	return mu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableEmail(s *string) *MemberUpdate {
	if s != nil {
		mu.SetEmail(*s)
	}
	return mu
}

// ClearEmail clears the value of the "email" field.
func (mu *MemberUpdate) ClearEmail() *MemberUpdate {
	mu.mutation.ClearEmail()
	return mu
}

// SetAvatar sets the "avatar" field.
func (mu *MemberUpdate) SetAvatar(s string) *MemberUpdate {
	mu.mutation.SetAvatar(s)
	return mu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableAvatar(s *string) *MemberUpdate {
	if s != nil {
		mu.SetAvatar(*s)
	}
	return mu
}

// ClearAvatar clears the value of the "avatar" field.
func (mu *MemberUpdate) ClearAvatar() *MemberUpdate {
	mu.mutation.ClearAvatar()
	return mu
}

// SetRanksID sets the "ranks" edge to the MemberRank entity by ID.
func (mu *MemberUpdate) SetRanksID(id uint64) *MemberUpdate {
	mu.mutation.SetRanksID(id)
	return mu
}

// SetNillableRanksID sets the "ranks" edge to the MemberRank entity by ID if the given value is not nil.
func (mu *MemberUpdate) SetNillableRanksID(id *uint64) *MemberUpdate {
	if id != nil {
		mu = mu.SetRanksID(*id)
	}
	return mu
}

// SetRanks sets the "ranks" edge to the MemberRank entity.
func (mu *MemberUpdate) SetRanks(m *MemberRank) *MemberUpdate {
	return mu.SetRanksID(m.ID)
}

// Mutation returns the MemberMutation object of the builder.
func (mu *MemberUpdate) Mutation() *MemberMutation {
	return mu.mutation
}

// ClearRanks clears the "ranks" edge to the MemberRank entity.
func (mu *MemberUpdate) ClearRanks() *MemberUpdate {
	mu.mutation.ClearRanks()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MemberUpdate) Save(ctx context.Context) (int, error) {
	mu.defaults()
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MemberUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MemberUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MemberUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MemberUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := member.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

func (mu *MemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(member.Table, member.Columns, sqlgraph.NewFieldSpec(member.FieldID, field.TypeUUID))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(member.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.Status(); ok {
		_spec.SetField(member.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := mu.mutation.AddedStatus(); ok {
		_spec.AddField(member.FieldStatus, field.TypeUint8, value)
	}
	if mu.mutation.StatusCleared() {
		_spec.ClearField(member.FieldStatus, field.TypeUint8)
	}
	if value, ok := mu.mutation.Username(); ok {
		_spec.SetField(member.FieldUsername, field.TypeString, value)
	}
	if value, ok := mu.mutation.Password(); ok {
		_spec.SetField(member.FieldPassword, field.TypeString, value)
	}
	if value, ok := mu.mutation.Nickname(); ok {
		_spec.SetField(member.FieldNickname, field.TypeString, value)
	}
	if value, ok := mu.mutation.Mobile(); ok {
		_spec.SetField(member.FieldMobile, field.TypeString, value)
	}
	if mu.mutation.MobileCleared() {
		_spec.ClearField(member.FieldMobile, field.TypeString)
	}
	if value, ok := mu.mutation.Email(); ok {
		_spec.SetField(member.FieldEmail, field.TypeString, value)
	}
	if mu.mutation.EmailCleared() {
		_spec.ClearField(member.FieldEmail, field.TypeString)
	}
	if value, ok := mu.mutation.Avatar(); ok {
		_spec.SetField(member.FieldAvatar, field.TypeString, value)
	}
	if mu.mutation.AvatarCleared() {
		_spec.ClearField(member.FieldAvatar, field.TypeString)
	}
	if mu.mutation.RanksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RanksIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MemberUpdateOne is the builder for updating a single Member entity.
type MemberUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MemberMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MemberUpdateOne) SetUpdatedAt(t time.Time) *MemberUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetStatus sets the "status" field.
func (muo *MemberUpdateOne) SetStatus(u uint8) *MemberUpdateOne {
	muo.mutation.ResetStatus()
	muo.mutation.SetStatus(u)
	return muo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableStatus(u *uint8) *MemberUpdateOne {
	if u != nil {
		muo.SetStatus(*u)
	}
	return muo
}

// AddStatus adds u to the "status" field.
func (muo *MemberUpdateOne) AddStatus(u int8) *MemberUpdateOne {
	muo.mutation.AddStatus(u)
	return muo
}

// ClearStatus clears the value of the "status" field.
func (muo *MemberUpdateOne) ClearStatus() *MemberUpdateOne {
	muo.mutation.ClearStatus()
	return muo
}

// SetUsername sets the "username" field.
func (muo *MemberUpdateOne) SetUsername(s string) *MemberUpdateOne {
	muo.mutation.SetUsername(s)
	return muo
}

// SetPassword sets the "password" field.
func (muo *MemberUpdateOne) SetPassword(s string) *MemberUpdateOne {
	muo.mutation.SetPassword(s)
	return muo
}

// SetNickname sets the "nickname" field.
func (muo *MemberUpdateOne) SetNickname(s string) *MemberUpdateOne {
	muo.mutation.SetNickname(s)
	return muo
}

// SetRankID sets the "rank_id" field.
func (muo *MemberUpdateOne) SetRankID(u uint64) *MemberUpdateOne {
	muo.mutation.SetRankID(u)
	return muo
}

// SetNillableRankID sets the "rank_id" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableRankID(u *uint64) *MemberUpdateOne {
	if u != nil {
		muo.SetRankID(*u)
	}
	return muo
}

// ClearRankID clears the value of the "rank_id" field.
func (muo *MemberUpdateOne) ClearRankID() *MemberUpdateOne {
	muo.mutation.ClearRankID()
	return muo
}

// SetMobile sets the "mobile" field.
func (muo *MemberUpdateOne) SetMobile(s string) *MemberUpdateOne {
	muo.mutation.SetMobile(s)
	return muo
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableMobile(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetMobile(*s)
	}
	return muo
}

// ClearMobile clears the value of the "mobile" field.
func (muo *MemberUpdateOne) ClearMobile() *MemberUpdateOne {
	muo.mutation.ClearMobile()
	return muo
}

// SetEmail sets the "email" field.
func (muo *MemberUpdateOne) SetEmail(s string) *MemberUpdateOne {
	muo.mutation.SetEmail(s)
	return muo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableEmail(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetEmail(*s)
	}
	return muo
}

// ClearEmail clears the value of the "email" field.
func (muo *MemberUpdateOne) ClearEmail() *MemberUpdateOne {
	muo.mutation.ClearEmail()
	return muo
}

// SetAvatar sets the "avatar" field.
func (muo *MemberUpdateOne) SetAvatar(s string) *MemberUpdateOne {
	muo.mutation.SetAvatar(s)
	return muo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableAvatar(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetAvatar(*s)
	}
	return muo
}

// ClearAvatar clears the value of the "avatar" field.
func (muo *MemberUpdateOne) ClearAvatar() *MemberUpdateOne {
	muo.mutation.ClearAvatar()
	return muo
}

// SetRanksID sets the "ranks" edge to the MemberRank entity by ID.
func (muo *MemberUpdateOne) SetRanksID(id uint64) *MemberUpdateOne {
	muo.mutation.SetRanksID(id)
	return muo
}

// SetNillableRanksID sets the "ranks" edge to the MemberRank entity by ID if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableRanksID(id *uint64) *MemberUpdateOne {
	if id != nil {
		muo = muo.SetRanksID(*id)
	}
	return muo
}

// SetRanks sets the "ranks" edge to the MemberRank entity.
func (muo *MemberUpdateOne) SetRanks(m *MemberRank) *MemberUpdateOne {
	return muo.SetRanksID(m.ID)
}

// Mutation returns the MemberMutation object of the builder.
func (muo *MemberUpdateOne) Mutation() *MemberMutation {
	return muo.mutation
}

// ClearRanks clears the "ranks" edge to the MemberRank entity.
func (muo *MemberUpdateOne) ClearRanks() *MemberUpdateOne {
	muo.mutation.ClearRanks()
	return muo
}

// Where appends a list predicates to the MemberUpdate builder.
func (muo *MemberUpdateOne) Where(ps ...predicate.Member) *MemberUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MemberUpdateOne) Select(field string, fields ...string) *MemberUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Member entity.
func (muo *MemberUpdateOne) Save(ctx context.Context) (*Member, error) {
	muo.defaults()
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MemberUpdateOne) SaveX(ctx context.Context) *Member {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MemberUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MemberUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MemberUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := member.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

func (muo *MemberUpdateOne) sqlSave(ctx context.Context) (_node *Member, err error) {
	_spec := sqlgraph.NewUpdateSpec(member.Table, member.Columns, sqlgraph.NewFieldSpec(member.FieldID, field.TypeUUID))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Member.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, member.FieldID)
		for _, f := range fields {
			if !member.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != member.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(member.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.Status(); ok {
		_spec.SetField(member.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := muo.mutation.AddedStatus(); ok {
		_spec.AddField(member.FieldStatus, field.TypeUint8, value)
	}
	if muo.mutation.StatusCleared() {
		_spec.ClearField(member.FieldStatus, field.TypeUint8)
	}
	if value, ok := muo.mutation.Username(); ok {
		_spec.SetField(member.FieldUsername, field.TypeString, value)
	}
	if value, ok := muo.mutation.Password(); ok {
		_spec.SetField(member.FieldPassword, field.TypeString, value)
	}
	if value, ok := muo.mutation.Nickname(); ok {
		_spec.SetField(member.FieldNickname, field.TypeString, value)
	}
	if value, ok := muo.mutation.Mobile(); ok {
		_spec.SetField(member.FieldMobile, field.TypeString, value)
	}
	if muo.mutation.MobileCleared() {
		_spec.ClearField(member.FieldMobile, field.TypeString)
	}
	if value, ok := muo.mutation.Email(); ok {
		_spec.SetField(member.FieldEmail, field.TypeString, value)
	}
	if muo.mutation.EmailCleared() {
		_spec.ClearField(member.FieldEmail, field.TypeString)
	}
	if value, ok := muo.mutation.Avatar(); ok {
		_spec.SetField(member.FieldAvatar, field.TypeString, value)
	}
	if muo.mutation.AvatarCleared() {
		_spec.ClearField(member.FieldAvatar, field.TypeString)
	}
	if muo.mutation.RanksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RanksIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Member{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
