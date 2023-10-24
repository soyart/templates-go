package model

import (
	"example.com/service-clean/module/todo/common"
	"fmt"
	"gorm.io/gorm"

	"time"
)

type Status int

const (
	StatusTodo Status = iota
	StatusInProgress
	StatusDone
)

type Todo struct {
	ID          string    `gorm:"column:id;type:varchar;primaryKey;index:,type:btree;not null" json:"id"`
	Title       string    `gorm:"column:title;type:varchar;index:,type:btree;not null" json:"title"`
	Description string    `gorm:"column:description;type:varchar;not null" json:"description"`
	Status      Status    `gorm:"column:status;type:varchar;not null" json:"status"`
	TodoDate    time.Time `gorm:"column:todo_date;type:timestamp WITH TIME ZONE;not null" json:"todoDate"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp WITH TIME ZONE;not null" json:"createdAt"`
}

func (t *Todo) GetDBQuery(db *gorm.DB) *gorm.DB {
	if t.ID != "" {
		db = db.Where(fmt.Sprintf("%s = ?", common.ColNameID), t.ID)
	}
	if t.Title != "" {
		db = db.Where(fmt.Sprintf("%s = ?", common.ColNameTitle), t.Title)
	}
	return nil
}

func (t *Todo) TableName() string {
	return common.TableNameTodo
}
