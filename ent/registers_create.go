// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/AnirudhAgnihotri2902/jwt-using-go-and-postgres/ent/registers"
)

// RegistersCreate is the builder for creating a Registers entity.
type RegistersCreate struct {
	config
	mutation *RegistersMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (rc *RegistersCreate) SetUsername(s string) *RegistersCreate {
	rc.mutation.SetUsername(s)
	return rc
}

// SetPassword sets the "password" field.
func (rc *RegistersCreate) SetPassword(s string) *RegistersCreate {
	rc.mutation.SetPassword(s)
	return rc
}

// Mutation returns the RegistersMutation object of the builder.
func (rc *RegistersCreate) Mutation() *RegistersMutation {
	return rc.mutation
}

// Save creates the Registers in the database.
func (rc *RegistersCreate) Save(ctx context.Context) (*Registers, error) {
	return withHooks[*Registers, RegistersMutation](ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RegistersCreate) SaveX(ctx context.Context) *Registers {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RegistersCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RegistersCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RegistersCreate) check() error {
	if _, ok := rc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "Registers.username"`)}
	}
	if _, ok := rc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Registers.password"`)}
	}
	return nil
}

func (rc *RegistersCreate) sqlSave(ctx context.Context) (*Registers, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RegistersCreate) createSpec() (*Registers, *sqlgraph.CreateSpec) {
	var (
		_node = &Registers{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: registers.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: registers.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.Username(); ok {
		_spec.SetField(registers.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := rc.mutation.Password(); ok {
		_spec.SetField(registers.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	return _node, _spec
}

// RegistersCreateBulk is the builder for creating many Registers entities in bulk.
type RegistersCreateBulk struct {
	config
	builders []*RegistersCreate
}

// Save creates the Registers entities in the database.
func (rcb *RegistersCreateBulk) Save(ctx context.Context) ([]*Registers, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Registers, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RegistersMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RegistersCreateBulk) SaveX(ctx context.Context) []*Registers {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RegistersCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RegistersCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}