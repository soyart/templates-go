package database

import (
	"context"
	"example.com/service-clean/module/todo/domain"
	"gorm.io/gorm"
)

type Adapter interface {
	CreateTodo(context.Context, domain.Todo) error
	GetTodo(context.Context, string) (*domain.Todo, error)
	GetTodos(context.Context) ([]domain.Todo, error)
	UpdateTodo(context.Context, domain.Todo) error
	DeleteTodo(context.Context, string) error
}

type adapterImpl struct {
	db GormWrapper
}

func New(db *gorm.DB) Adapter {
	return &adapterImpl{db: newGormWrapper(db)}
}

type Query interface {
	GetDBQuery(db *gorm.DB) *gorm.DB
}

type Model interface {
	TableName() string
}

type GormWrapper interface {
	Create(context.Context, Model) error
	Delete(context.Context, Query) error
	First(context.Context, Query, Model) error
	Find(context.Context, Query, interface{}) error
	Update(context.Context, Query, Model) error
}

type gormWrapper struct {
	db *gorm.DB
}

func (g gormWrapper) Create(ctx context.Context, model Model) error {
	return g.db.WithContext(ctx).Create(model).Error
}

func (g gormWrapper) Delete(ctx context.Context, query Query) error {
	return g.db.WithContext(ctx).Delete(query).Error
}

func (g gormWrapper) First(ctx context.Context, query Query, model Model) error {
	return query.GetDBQuery(g.db.WithContext(ctx)).First(model).Error
}

func (g gormWrapper) Find(ctx context.Context, query Query, model interface{}) error {
	db := g.db.WithContext(ctx)
	if query != nil {
		return query.GetDBQuery(db).Find(model).Error
	}
	return db.Find(model).Error
}

func (g gormWrapper) Update(ctx context.Context, query Query, model Model) error {
	if err := query.GetDBQuery(g.db.WithContext(ctx)).Updates(model).Error; err != nil {
		return err
	}
	if g.db.RowsAffected == 0 {
		return errNotFoundTodo
	}
	return nil
}

func newGormWrapper(db *gorm.DB) GormWrapper {
	return &gormWrapper{
		db: db,
	}
}
