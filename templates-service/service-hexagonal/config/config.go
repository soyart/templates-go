package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Rest ConfRest `mapstructure:"rest"`
	Grpc ConfGrpc `mapstructure:"grpc"`

	Postgres ConfPostgres `mapstructure:"postgres"`
	Redis    ConfRedis    `mapstructure:"redis"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.SetDefault("rest.address", "127.0.0.1:8000")
	viper.SetDefault("rest.jwt_secret", []byte("servicehex-secret"))

	viper.SetDefault("postgres.host", "127.0.0.1")
	viper.SetDefault("postgres.port", "5432")
	viper.SetDefault("postgres.password", "postgres")
	viper.SetDefault("postgres.name", "postgres")

	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("redis.db", 5)

	viper.AutomaticEnv()
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
