package store

import "Assessment/tapcontext"

type Repository interface {
	GetAllForms(ctx tapcontext.TContext) error
	//FindFormByID(ctx tapcontext.TContext, id string) (*model.Form, error)
	//InsertForm(ctx tapcontext.TContext, form *model.Form) error
	//UpdateForm(ctx tapcontext.TContext, id string, form *model.Form) error
}
