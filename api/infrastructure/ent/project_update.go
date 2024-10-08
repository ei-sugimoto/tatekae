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

// ProjectUpdate is the builder for updating Project entities.
type ProjectUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectMutation
}

// Where appends a list predicates to the ProjectUpdate builder.
func (pu *ProjectUpdate) Where(ps ...predicate.Project) *ProjectUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *ProjectUpdate) SetName(s string) *ProjectUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableName(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetCreatedBy sets the "created_by" field.
func (pu *ProjectUpdate) SetCreatedBy(i int) *ProjectUpdate {
	pu.mutation.ResetCreatedBy()
	pu.mutation.SetCreatedBy(i)
	return pu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableCreatedBy(i *int) *ProjectUpdate {
	if i != nil {
		pu.SetCreatedBy(*i)
	}
	return pu
}

// AddCreatedBy adds i to the "created_by" field.
func (pu *ProjectUpdate) AddCreatedBy(i int) *ProjectUpdate {
	pu.mutation.AddCreatedBy(i)
	return pu
}

// AddBillIDs adds the "bills" edge to the Bill entity by IDs.
func (pu *ProjectUpdate) AddBillIDs(ids ...int) *ProjectUpdate {
	pu.mutation.AddBillIDs(ids...)
	return pu
}

// AddBills adds the "bills" edges to the Bill entity.
func (pu *ProjectUpdate) AddBills(b ...*Bill) *ProjectUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pu.AddBillIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (pu *ProjectUpdate) AddUserIDs(ids ...int) *ProjectUpdate {
	pu.mutation.AddUserIDs(ids...)
	return pu
}

// AddUsers adds the "users" edges to the User entity.
func (pu *ProjectUpdate) AddUsers(u ...*User) *ProjectUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.AddUserIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (pu *ProjectUpdate) Mutation() *ProjectMutation {
	return pu.mutation
}

// ClearBills clears all "bills" edges to the Bill entity.
func (pu *ProjectUpdate) ClearBills() *ProjectUpdate {
	pu.mutation.ClearBills()
	return pu
}

// RemoveBillIDs removes the "bills" edge to Bill entities by IDs.
func (pu *ProjectUpdate) RemoveBillIDs(ids ...int) *ProjectUpdate {
	pu.mutation.RemoveBillIDs(ids...)
	return pu
}

// RemoveBills removes "bills" edges to Bill entities.
func (pu *ProjectUpdate) RemoveBills(b ...*Bill) *ProjectUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pu.RemoveBillIDs(ids...)
}

// ClearUsers clears all "users" edges to the User entity.
func (pu *ProjectUpdate) ClearUsers() *ProjectUpdate {
	pu.mutation.ClearUsers()
	return pu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (pu *ProjectUpdate) RemoveUserIDs(ids ...int) *ProjectUpdate {
	pu.mutation.RemoveUserIDs(ids...)
	return pu
}

// RemoveUsers removes "users" edges to User entities.
func (pu *ProjectUpdate) RemoveUsers(u ...*User) *ProjectUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProjectUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProjectUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProjectUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProjectUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.CreatedBy(); ok {
		_spec.SetField(project.FieldCreatedBy, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(project.FieldCreatedBy, field.TypeInt, value)
	}
	if pu.mutation.BillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.BillsTable,
			Columns: []string{project.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedBillsIDs(); len(nodes) > 0 && !pu.mutation.BillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.BillsTable,
			Columns: []string{project.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.BillsTable,
			Columns: []string{project.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.UsersTable,
			Columns: project.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !pu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.UsersTable,
			Columns: project.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.UsersTable,
			Columns: project.UsersPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProjectUpdateOne is the builder for updating a single Project entity.
type ProjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectMutation
}

// SetName sets the "name" field.
func (puo *ProjectUpdateOne) SetName(s string) *ProjectUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableName(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetCreatedBy sets the "created_by" field.
func (puo *ProjectUpdateOne) SetCreatedBy(i int) *ProjectUpdateOne {
	puo.mutation.ResetCreatedBy()
	puo.mutation.SetCreatedBy(i)
	return puo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableCreatedBy(i *int) *ProjectUpdateOne {
	if i != nil {
		puo.SetCreatedBy(*i)
	}
	return puo
}

// AddCreatedBy adds i to the "created_by" field.
func (puo *ProjectUpdateOne) AddCreatedBy(i int) *ProjectUpdateOne {
	puo.mutation.AddCreatedBy(i)
	return puo
}

// AddBillIDs adds the "bills" edge to the Bill entity by IDs.
func (puo *ProjectUpdateOne) AddBillIDs(ids ...int) *ProjectUpdateOne {
	puo.mutation.AddBillIDs(ids...)
	return puo
}

// AddBills adds the "bills" edges to the Bill entity.
func (puo *ProjectUpdateOne) AddBills(b ...*Bill) *ProjectUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return puo.AddBillIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (puo *ProjectUpdateOne) AddUserIDs(ids ...int) *ProjectUpdateOne {
	puo.mutation.AddUserIDs(ids...)
	return puo
}

// AddUsers adds the "users" edges to the User entity.
func (puo *ProjectUpdateOne) AddUsers(u ...*User) *ProjectUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.AddUserIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (puo *ProjectUpdateOne) Mutation() *ProjectMutation {
	return puo.mutation
}

// ClearBills clears all "bills" edges to the Bill entity.
func (puo *ProjectUpdateOne) ClearBills() *ProjectUpdateOne {
	puo.mutation.ClearBills()
	return puo
}

// RemoveBillIDs removes the "bills" edge to Bill entities by IDs.
func (puo *ProjectUpdateOne) RemoveBillIDs(ids ...int) *ProjectUpdateOne {
	puo.mutation.RemoveBillIDs(ids...)
	return puo
}

// RemoveBills removes "bills" edges to Bill entities.
func (puo *ProjectUpdateOne) RemoveBills(b ...*Bill) *ProjectUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return puo.RemoveBillIDs(ids...)
}

// ClearUsers clears all "users" edges to the User entity.
func (puo *ProjectUpdateOne) ClearUsers() *ProjectUpdateOne {
	puo.mutation.ClearUsers()
	return puo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (puo *ProjectUpdateOne) RemoveUserIDs(ids ...int) *ProjectUpdateOne {
	puo.mutation.RemoveUserIDs(ids...)
	return puo
}

// RemoveUsers removes "users" edges to User entities.
func (puo *ProjectUpdateOne) RemoveUsers(u ...*User) *ProjectUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the ProjectUpdate builder.
func (puo *ProjectUpdateOne) Where(ps ...predicate.Project) *ProjectUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProjectUpdateOne) Select(field string, fields ...string) *ProjectUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Project entity.
func (puo *ProjectUpdateOne) Save(ctx context.Context) (*Project, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProjectUpdateOne) SaveX(ctx context.Context) *Project {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProjectUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProjectUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProjectUpdateOne) sqlSave(ctx context.Context) (_node *Project, err error) {
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Project.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, project.FieldID)
		for _, f := range fields {
			if !project.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != project.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.CreatedBy(); ok {
		_spec.SetField(project.FieldCreatedBy, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(project.FieldCreatedBy, field.TypeInt, value)
	}
	if puo.mutation.BillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.BillsTable,
			Columns: []string{project.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedBillsIDs(); len(nodes) > 0 && !puo.mutation.BillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.BillsTable,
			Columns: []string{project.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.BillsTable,
			Columns: []string{project.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.UsersTable,
			Columns: project.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !puo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.UsersTable,
			Columns: project.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.UsersTable,
			Columns: project.UsersPrimaryKey,
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
	_node = &Project{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
