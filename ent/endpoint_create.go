// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/1nv8rzim/otn/ent/endpoint"
)

// EndpointCreate is the builder for creating a Endpoint entity.
type EndpointCreate struct {
	config
	mutation *EndpointMutation
	hooks    []Hook
}

// SetURI sets the "uri" field.
func (ec *EndpointCreate) SetURI(s string) *EndpointCreate {
	ec.mutation.SetURI(s)
	return ec
}

// SetContent sets the "content" field.
func (ec *EndpointCreate) SetContent(s string) *EndpointCreate {
	ec.mutation.SetContent(s)
	return ec
}

// SetCreatedAt sets the "created_at" field.
func (ec *EndpointCreate) SetCreatedAt(t time.Time) *EndpointCreate {
	ec.mutation.SetCreatedAt(t)
	return ec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ec *EndpointCreate) SetNillableCreatedAt(t *time.Time) *EndpointCreate {
	if t != nil {
		ec.SetCreatedAt(*t)
	}
	return ec
}

// Mutation returns the EndpointMutation object of the builder.
func (ec *EndpointCreate) Mutation() *EndpointMutation {
	return ec.mutation
}

// Save creates the Endpoint in the database.
func (ec *EndpointCreate) Save(ctx context.Context) (*Endpoint, error) {
	var (
		err  error
		node *Endpoint
	)
	ec.defaults()
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EndpointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			if node, err = ec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			if ec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ec.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ec.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Endpoint)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EndpointMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EndpointCreate) SaveX(ctx context.Context) *Endpoint {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EndpointCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EndpointCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EndpointCreate) defaults() {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		v := endpoint.DefaultCreatedAt()
		ec.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EndpointCreate) check() error {
	if _, ok := ec.mutation.URI(); !ok {
		return &ValidationError{Name: "uri", err: errors.New(`ent: missing required field "Endpoint.uri"`)}
	}
	if v, ok := ec.mutation.URI(); ok {
		if err := endpoint.URIValidator(v); err != nil {
			return &ValidationError{Name: "uri", err: fmt.Errorf(`ent: validator failed for field "Endpoint.uri": %w`, err)}
		}
	}
	if _, ok := ec.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Endpoint.content"`)}
	}
	if v, ok := ec.mutation.Content(); ok {
		if err := endpoint.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Endpoint.content": %w`, err)}
		}
	}
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Endpoint.created_at"`)}
	}
	return nil
}

func (ec *EndpointCreate) sqlSave(ctx context.Context) (*Endpoint, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ec *EndpointCreate) createSpec() (*Endpoint, *sqlgraph.CreateSpec) {
	var (
		_node = &Endpoint{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: endpoint.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: endpoint.FieldID,
			},
		}
	)
	if value, ok := ec.mutation.URI(); ok {
		_spec.SetField(endpoint.FieldURI, field.TypeString, value)
		_node.URI = value
	}
	if value, ok := ec.mutation.Content(); ok {
		_spec.SetField(endpoint.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.SetField(endpoint.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// EndpointCreateBulk is the builder for creating many Endpoint entities in bulk.
type EndpointCreateBulk struct {
	config
	builders []*EndpointCreate
}

// Save creates the Endpoint entities in the database.
func (ecb *EndpointCreateBulk) Save(ctx context.Context) ([]*Endpoint, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Endpoint, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EndpointMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EndpointCreateBulk) SaveX(ctx context.Context) []*Endpoint {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EndpointCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EndpointCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
