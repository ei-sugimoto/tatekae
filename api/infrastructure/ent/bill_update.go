// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ei-sugimoto/tatekae/api/infrastructure/ent/bill"
	"github.com/ei-sugimoto/tatekae/api/infrastructure/ent/predicate"
	"github.com/ei-sugimoto/tatekae/api/infrastructure/ent/project"
	"github.com/ei-sugimoto/tatekae/api/infrastructure/ent/user"
)

// BillUpdate is the builder for updating Bill entities.
type BillUpdate struct {
	config
	hooks    []Hook
	mutation *BillMutation
}

// Where appends a list predicates to the BillUpdate builder.
func (bu *BillUpdate) Where(ps ...predicate.Bill) *BillUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetPrice sets the "price" field.
func (bu *BillUpdate) SetPrice(i int) *BillUpdate {
	bu.mutation.ResetPrice()
	bu.mutation.SetPrice(i)
	return bu
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (bu *BillUpdate) SetNillablePrice(i *int) *BillUpdate {
	if i != nil {
		bu.SetPrice(*i)
	}
	return bu
}

// AddPrice adds i to the "price" field.
func (bu *BillUpdate) AddPrice(i int) *BillUpdate {
	bu.mutation.AddPrice(i)
	return bu
}

// SetProjectIDID sets the "project_id" edge to the Project entity by ID.
func (bu *BillUpdate) SetProjectIDID(id int) *BillUpdate {
	bu.mutation.SetProjectIDID(id)
	return bu
}

// SetNillableProjectIDID sets the "project_id" edge to the Project entity by ID if the given value is not nil.
func (bu *BillUpdate) SetNillableProjectIDID(id *int) *BillUpdate {
	if id != nil {
		bu = bu.SetProjectIDID(*id)
	}
	return bu
}

// SetProjectID sets the "project_id" edge to the Project entity.
func (bu *BillUpdate) SetProjectID(p *Project) *BillUpdate {
	return bu.SetProjectIDID(p.ID)
}

// SetSrcUserID sets the "src_user" edge to the User entity by ID.
func (bu *BillUpdate) SetSrcUserID(id int) *BillUpdate {
	bu.mutation.SetSrcUserID(id)
	return bu
}

// SetNillableSrcUserID sets the "src_user" edge to the User entity by ID if the given value is not nil.
func (bu *BillUpdate) SetNillableSrcUserID(id *int) *BillUpdate {
	if id != nil {
		bu = bu.SetSrcUserID(*id)
	}
	return bu
}

// SetSrcUser sets the "src_user" edge to the User entity.
func (bu *BillUpdate) SetSrcUser(u *User) *BillUpdate {
	return bu.SetSrcUserID(u.ID)
}

// SetDstUserID sets the "dst_user" edge to the User entity by ID.
func (bu *BillUpdate) SetDstUserID(id int) *BillUpdate {
	bu.mutation.SetDstUserID(id)
	return bu
}

// SetNillableDstUserID sets the "dst_user" edge to the User entity by ID if the given value is not nil.
func (bu *BillUpdate) SetNillableDstUserID(id *int) *BillUpdate {
	if id != nil {
		bu = bu.SetDstUserID(*id)
	}
	return bu
}

// SetDstUser sets the "dst_user" edge to the User entity.
func (bu *BillUpdate) SetDstUser(u *User) *BillUpdate {
	return bu.SetDstUserID(u.ID)
}

// Mutation returns the BillMutation object of the builder.
func (bu *BillUpdate) Mutation() *BillMutation {
	return bu.mutation
}

// ClearProjectID clears the "project_id" edge to the Project entity.
func (bu *BillUpdate) ClearProjectID() *BillUpdate {
	bu.mutation.ClearProjectID()
	return bu
}

// ClearSrcUser clears the "src_user" edge to the User entity.
func (bu *BillUpdate) ClearSrcUser() *BillUpdate {
	bu.mutation.ClearSrcUser()
	return bu
}

// ClearDstUser clears the "dst_user" edge to the User entity.
func (bu *BillUpdate) ClearDstUser() *BillUpdate {
	bu.mutation.ClearDstUser()
	return bu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BillUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BillUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BillUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BillUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BillUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(bill.Table, bill.Columns, sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Price(); ok {
		_spec.SetField(bill.FieldPrice, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedPrice(); ok {
		_spec.AddField(bill.FieldPrice, field.TypeInt, value)
	}
	if bu.mutation.ProjectIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.ProjectIDTable,
			Columns: []string{bill.ProjectIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.ProjectIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.ProjectIDTable,
			Columns: []string{bill.ProjectIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.SrcUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   bill.SrcUserTable,
			Columns: []string{bill.SrcUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.SrcUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   bill.SrcUserTable,
			Columns: []string{bill.SrcUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.DstUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   bill.DstUserTable,
			Columns: []string{bill.DstUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.DstUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   bill.DstUserTable,
			Columns: []string{bill.DstUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bill.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BillUpdateOne is the builder for updating a single Bill entity.
type BillUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BillMutation
}

// SetPrice sets the "price" field.
func (buo *BillUpdateOne) SetPrice(i int) *BillUpdateOne {
	buo.mutation.ResetPrice()
	buo.mutation.SetPrice(i)
	return buo
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (buo *BillUpdateOne) SetNillablePrice(i *int) *BillUpdateOne {
	if i != nil {
		buo.SetPrice(*i)
	}
	return buo
}

// AddPrice adds i to the "price" field.
func (buo *BillUpdateOne) AddPrice(i int) *BillUpdateOne {
	buo.mutation.AddPrice(i)
	return buo
}

// SetProjectIDID sets the "project_id" edge to the Project entity by ID.
func (buo *BillUpdateOne) SetProjectIDID(id int) *BillUpdateOne {
	buo.mutation.SetProjectIDID(id)
	return buo
}

// SetNillableProjectIDID sets the "project_id" edge to the Project entity by ID if the given value is not nil.
func (buo *BillUpdateOne) SetNillableProjectIDID(id *int) *BillUpdateOne {
	if id != nil {
		buo = buo.SetProjectIDID(*id)
	}
	return buo
}

// SetProjectID sets the "project_id" edge to the Project entity.
func (buo *BillUpdateOne) SetProjectID(p *Project) *BillUpdateOne {
	return buo.SetProjectIDID(p.ID)
}

// SetSrcUserID sets the "src_user" edge to the User entity by ID.
func (buo *BillUpdateOne) SetSrcUserID(id int) *BillUpdateOne {
	buo.mutation.SetSrcUserID(id)
	return buo
}

// SetNillableSrcUserID sets the "src_user" edge to the User entity by ID if the given value is not nil.
func (buo *BillUpdateOne) SetNillableSrcUserID(id *int) *BillUpdateOne {
	if id != nil {
		buo = buo.SetSrcUserID(*id)
	}
	return buo
}

// SetSrcUser sets the "src_user" edge to the User entity.
func (buo *BillUpdateOne) SetSrcUser(u *User) *BillUpdateOne {
	return buo.SetSrcUserID(u.ID)
}

// SetDstUserID sets the "dst_user" edge to the User entity by ID.
func (buo *BillUpdateOne) SetDstUserID(id int) *BillUpdateOne {
	buo.mutation.SetDstUserID(id)
	return buo
}

// SetNillableDstUserID sets the "dst_user" edge to the User entity by ID if the given value is not nil.
func (buo *BillUpdateOne) SetNillableDstUserID(id *int) *BillUpdateOne {
	if id != nil {
		buo = buo.SetDstUserID(*id)
	}
	return buo
}

// SetDstUser sets the "dst_user" edge to the User entity.
func (buo *BillUpdateOne) SetDstUser(u *User) *BillUpdateOne {
	return buo.SetDstUserID(u.ID)
}

// Mutation returns the BillMutation object of the builder.
func (buo *BillUpdateOne) Mutation() *BillMutation {
	return buo.mutation
}

// ClearProjectID clears the "project_id" edge to the Project entity.
func (buo *BillUpdateOne) ClearProjectID() *BillUpdateOne {
	buo.mutation.ClearProjectID()
	return buo
}

// ClearSrcUser clears the "src_user" edge to the User entity.
func (buo *BillUpdateOne) ClearSrcUser() *BillUpdateOne {
	buo.mutation.ClearSrcUser()
	return buo
}

// ClearDstUser clears the "dst_user" edge to the User entity.
func (buo *BillUpdateOne) ClearDstUser() *BillUpdateOne {
	buo.mutation.ClearDstUser()
	return buo
}

// Where appends a list predicates to the BillUpdate builder.
func (buo *BillUpdateOne) Where(ps ...predicate.Bill) *BillUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BillUpdateOne) Select(field string, fields ...string) *BillUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Bill entity.
func (buo *BillUpdateOne) Save(ctx context.Context) (*Bill, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BillUpdateOne) SaveX(ctx context.Context) *Bill {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BillUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BillUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BillUpdateOne) sqlSave(ctx context.Context) (_node *Bill, err error) {
	_spec := sqlgraph.NewUpdateSpec(bill.Table, bill.Columns, sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Bill.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bill.FieldID)
		for _, f := range fields {
			if !bill.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != bill.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Price(); ok {
		_spec.SetField(bill.FieldPrice, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedPrice(); ok {
		_spec.AddField(bill.FieldPrice, field.TypeInt, value)
	}
	if buo.mutation.ProjectIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.ProjectIDTable,
			Columns: []string{bill.ProjectIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.ProjectIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.ProjectIDTable,
			Columns: []string{bill.ProjectIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.SrcUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   bill.SrcUserTable,
			Columns: []string{bill.SrcUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.SrcUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   bill.SrcUserTable,
			Columns: []string{bill.SrcUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.DstUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   bill.DstUserTable,
			Columns: []string{bill.DstUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.DstUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   bill.DstUserTable,
			Columns: []string{bill.DstUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Bill{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bill.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
