// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ei-sugimoto/tatekae/api/infrastructure/ent/bill"
	"github.com/ei-sugimoto/tatekae/api/infrastructure/ent/project"
	"github.com/ei-sugimoto/tatekae/api/infrastructure/ent/user"
)

// BillCreate is the builder for creating a Bill entity.
type BillCreate struct {
	config
	mutation *BillMutation
	hooks    []Hook
}

// SetPrice sets the "price" field.
func (bc *BillCreate) SetPrice(i int) *BillCreate {
	bc.mutation.SetPrice(i)
	return bc
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (bc *BillCreate) SetProjectID(id int) *BillCreate {
	bc.mutation.SetProjectID(id)
	return bc
}

// SetNillableProjectID sets the "project" edge to the Project entity by ID if the given value is not nil.
func (bc *BillCreate) SetNillableProjectID(id *int) *BillCreate {
	if id != nil {
		bc = bc.SetProjectID(*id)
	}
	return bc
}

// SetProject sets the "project" edge to the Project entity.
func (bc *BillCreate) SetProject(p *Project) *BillCreate {
	return bc.SetProjectID(p.ID)
}

// SetSrcUserID sets the "src_user" edge to the User entity by ID.
func (bc *BillCreate) SetSrcUserID(id int) *BillCreate {
	bc.mutation.SetSrcUserID(id)
	return bc
}

// SetNillableSrcUserID sets the "src_user" edge to the User entity by ID if the given value is not nil.
func (bc *BillCreate) SetNillableSrcUserID(id *int) *BillCreate {
	if id != nil {
		bc = bc.SetSrcUserID(*id)
	}
	return bc
}

// SetSrcUser sets the "src_user" edge to the User entity.
func (bc *BillCreate) SetSrcUser(u *User) *BillCreate {
	return bc.SetSrcUserID(u.ID)
}

// SetDstUserID sets the "dst_user" edge to the User entity by ID.
func (bc *BillCreate) SetDstUserID(id int) *BillCreate {
	bc.mutation.SetDstUserID(id)
	return bc
}

// SetNillableDstUserID sets the "dst_user" edge to the User entity by ID if the given value is not nil.
func (bc *BillCreate) SetNillableDstUserID(id *int) *BillCreate {
	if id != nil {
		bc = bc.SetDstUserID(*id)
	}
	return bc
}

// SetDstUser sets the "dst_user" edge to the User entity.
func (bc *BillCreate) SetDstUser(u *User) *BillCreate {
	return bc.SetDstUserID(u.ID)
}

// Mutation returns the BillMutation object of the builder.
func (bc *BillCreate) Mutation() *BillMutation {
	return bc.mutation
}

// Save creates the Bill in the database.
func (bc *BillCreate) Save(ctx context.Context) (*Bill, error) {
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BillCreate) SaveX(ctx context.Context) *Bill {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BillCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BillCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BillCreate) check() error {
	if _, ok := bc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "Bill.price"`)}
	}
	return nil
}

func (bc *BillCreate) sqlSave(ctx context.Context) (*Bill, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BillCreate) createSpec() (*Bill, *sqlgraph.CreateSpec) {
	var (
		_node = &Bill{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(bill.Table, sqlgraph.NewFieldSpec(bill.FieldID, field.TypeInt))
	)
	if value, ok := bc.mutation.Price(); ok {
		_spec.SetField(bill.FieldPrice, field.TypeInt, value)
		_node.Price = value
	}
	if nodes := bc.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.ProjectTable,
			Columns: []string{bill.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.project_bills = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.SrcUserIDs(); len(nodes) > 0 {
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
		_node.user_src_bill = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.DstUserIDs(); len(nodes) > 0 {
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
		_node.user_dst_bill = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BillCreateBulk is the builder for creating many Bill entities in bulk.
type BillCreateBulk struct {
	config
	err      error
	builders []*BillCreate
}

// Save creates the Bill entities in the database.
func (bcb *BillCreateBulk) Save(ctx context.Context) ([]*Bill, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Bill, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BillMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BillCreateBulk) SaveX(ctx context.Context) []*Bill {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BillCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BillCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
