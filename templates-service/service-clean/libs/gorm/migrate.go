package gorm

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	ID          string    `gorm:"column:id;type:varchar;primaryKey;index:,type:btree;not null" json:"id"`
	Title       string    `gorm:"column:title;type:varchar;index:,type:btree;not null" json:"title"`
	Description string    `gorm:"column:description;type:varchar;not null" json:"description"`
	Status      string    `gorm:"column:status;type:varchar;not null" json:"status"`
	TodoDate    time.Time `gorm:"column:todo_date;type:timestamp WITH TIME ZONE;not null" json:"todoDate"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp WITH TIME ZONE;not null" json:"createdAt"`
}

func migrateTableSchema(gormDB *gorm.DB) error {
	if err := gormDB.AutoMigrate(
		&Todo{},
	); err != nil {
		return fmt.Errorf("failed to migrate db: %w", err)
	}

	return nil
}
