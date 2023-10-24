package gorm

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const (
	envDatabaseHost     = "DATABASE_HOST"
	envDatabaseUser     = "DATABASE_USER"
	envDatabasePassword = "DATABASE_PASSWORD"
	envDatabaseName     = "DATABASE_NAME"
	envDatabasePort     = "DATABASE_PORT"
	envDatabaseSSLMode  = "DATABASE_SSL_MODE"
	envDatabaseTimezone = "DATABASE_TIMEZONE"
)

func NewGormProvider(dialector gorm.Dialector) (*gorm.DB, error) {

	gormDB, err := gorm.Open(dialector)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate new gorm provider: %w", err)
	}

	if err = migrateTableSchema(gormDB); err != nil {
		return nil, err
	}

	return gormDB, nil
}

func GetDSNByEnv() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		viper.GetString(envDatabaseHost),
		viper.GetString(envDatabaseUser),
		viper.GetString(envDatabasePassword),
		viper.GetString(envDatabaseName),
		viper.GetString(envDatabasePort),
		viper.GetString(envDatabaseSSLMode),
		viper.GetString(envDatabaseTimezone),
	)
}
