// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent/group"
	"github.com/hay-kot/homebox/backend/internal/data/ent/notifier"
	"github.com/hay-kot/homebox/backend/internal/data/ent/user"
)

// NotifierCreate is the builder for creating a Notifier entity.
type NotifierCreate struct {
	config
	mutation *NotifierMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (nc *NotifierCreate) SetCreatedAt(t time.Time) *NotifierCreate {
	nc.mutation.SetCreatedAt(t)
	return nc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nc *NotifierCreate) SetNillableCreatedAt(t *time.Time) *NotifierCreate {
	if t != nil {
		nc.SetCreatedAt(*t)
	}
	return nc
}

// SetUpdatedAt sets the "updated_at" field.
func (nc *NotifierCreate) SetUpdatedAt(t time.Time) *NotifierCreate {
	nc.mutation.SetUpdatedAt(t)
	return nc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nc *NotifierCreate) SetNillableUpdatedAt(t *time.Time) *NotifierCreate {
	if t != nil {
		nc.SetUpdatedAt(*t)
	}
	return nc
}

// SetGroupID sets the "group_id" field.
func (nc *NotifierCreate) SetGroupID(u uuid.UUID) *NotifierCreate {
	nc.mutation.SetGroupID(u)
	return nc
}

// SetUserID sets the "user_id" field.
func (nc *NotifierCreate) SetUserID(u uuid.UUID) *NotifierCreate {
	nc.mutation.SetUserID(u)
	return nc
}

// SetName sets the "name" field.
func (nc *NotifierCreate) SetName(s string) *NotifierCreate {
	nc.mutation.SetName(s)
	return nc
}

// SetURL sets the "url" field.
func (nc *NotifierCreate) SetURL(s string) *NotifierCreate {
	nc.mutation.SetURL(s)
	return nc
}

// SetIsActive sets the "is_active" field.
func (nc *NotifierCreate) SetIsActive(b bool) *NotifierCreate {
	nc.mutation.SetIsActive(b)
	return nc
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (nc *NotifierCreate) SetNillableIsActive(b *bool) *NotifierCreate {
	if b != nil {
		nc.SetIsActive(*b)
	}
	return nc
}

// SetID sets the "id" field.
func (nc *NotifierCreate) SetID(u uuid.UUID) *NotifierCreate {
	nc.mutation.SetID(u)
	return nc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (nc *NotifierCreate) SetNillableID(u *uuid.UUID) *NotifierCreate {
	if u != nil {
		nc.SetID(*u)
	}
	return nc
}

// SetGroup sets the "group" edge to the Group entity.
func (nc *NotifierCreate) SetGroup(g *Group) *NotifierCreate {
	return nc.SetGroupID(g.ID)
}

// SetUser sets the "user" edge to the User entity.
func (nc *NotifierCreate) SetUser(u *User) *NotifierCreate {
	return nc.SetUserID(u.ID)
}

// Mutation returns the NotifierMutation object of the builder.
func (nc *NotifierCreate) Mutation() *NotifierMutation {
	return nc.mutation
}

// Save creates the Notifier in the database.
func (nc *NotifierCreate) Save(ctx context.Context) (*Notifier, error) {
	nc.defaults()
	return withHooks(ctx, nc.sqlSave, nc.mutation, nc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NotifierCreate) SaveX(ctx context.Context) *Notifier {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NotifierCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NotifierCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NotifierCreate) defaults() {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		v := notifier.DefaultCreatedAt()
		nc.mutation.SetCreatedAt(v)
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		v := notifier.DefaultUpdatedAt()
		nc.mutation.SetUpdatedAt(v)
	}
	if _, ok := nc.mutation.IsActive(); !ok {
		v := notifier.DefaultIsActive
		nc.mutation.SetIsActive(v)
	}
	if _, ok := nc.mutation.ID(); !ok {
		v := notifier.DefaultID()
		nc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NotifierCreate) check() error {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Notifier.created_at"`)}
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Notifier.updated_at"`)}
	}
	if _, ok := nc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group_id", err: errors.New(`ent: missing required field "Notifier.group_id"`)}
	}
	if _, ok := nc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Notifier.user_id"`)}
	}
	if _, ok := nc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Notifier.name"`)}
	}
	if v, ok := nc.mutation.Name(); ok {
		if err := notifier.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Notifier.name": %w`, err)}
		}
	}
	if _, ok := nc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Notifier.url"`)}
	}
	if v, ok := nc.mutation.URL(); ok {
		if err := notifier.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Notifier.url": %w`, err)}
		}
	}
	if _, ok := nc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "Notifier.is_active"`)}
	}
	if _, ok := nc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required edge "Notifier.group"`)}
	}
	if _, ok := nc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Notifier.user"`)}
	}
	return nil
}

func (nc *NotifierCreate) sqlSave(ctx context.Context) (*Notifier, error) {
	if err := nc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
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
	nc.mutation.id = &_node.ID
	nc.mutation.done = true
	return _node, nil
}

func (nc *NotifierCreate) createSpec() (*Notifier, *sqlgraph.CreateSpec) {
	var (
		_node = &Notifier{config: nc.config}
		_spec = sqlgraph.NewCreateSpec(notifier.Table, sqlgraph.NewFieldSpec(notifier.FieldID, field.TypeUUID))
	)
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := nc.mutation.CreatedAt(); ok {
		_spec.SetField(notifier.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := nc.mutation.UpdatedAt(); ok {
		_spec.SetField(notifier.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := nc.mutation.Name(); ok {
		_spec.SetField(notifier.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := nc.mutation.URL(); ok {
		_spec.SetField(notifier.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := nc.mutation.IsActive(); ok {
		_spec.SetField(notifier.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if nodes := nc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifier.GroupTable,
			Columns: []string{notifier.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.GroupID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifier.UserTable,
			Columns: []string{notifier.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NotifierCreateBulk is the builder for creating many Notifier entities in bulk.
type NotifierCreateBulk struct {
	config
	builders []*NotifierCreate
}

// Save creates the Notifier entities in the database.
func (ncb *NotifierCreateBulk) Save(ctx context.Context) ([]*Notifier, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Notifier, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NotifierMutation)
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
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NotifierCreateBulk) SaveX(ctx context.Context) []*Notifier {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NotifierCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NotifierCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}
