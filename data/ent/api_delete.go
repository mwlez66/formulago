// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"formulago/data/ent/api"
	"formulago/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// APIDelete is the builder for deleting a API entity.
type APIDelete struct {
	config
	hooks    []Hook
	mutation *APIMutation
}

// Where appends a list predicates to the APIDelete builder.
func (ad *APIDelete) Where(ps ...predicate.API) *APIDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *APIDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, APIMutation](ctx, ad.sqlExec, ad.mutation, ad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *APIDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *APIDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: api.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: api.FieldID,
			},
		},
	}
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ad.mutation.done = true
	return affected, err
}

// APIDeleteOne is the builder for deleting a single API entity.
type APIDeleteOne struct {
	ad *APIDelete
}

// Exec executes the deletion query.
func (ado *APIDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{api.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *APIDeleteOne) ExecX(ctx context.Context) {
	ado.ad.ExecX(ctx)
}