package config

import (
	"github.com/spf13/viper"
	"log"
)

const (
	GrpcPortKey         = "GRPC_PORT"
	DatabaseHostKey     = "DATABASE_HOST"
	DatabasePortKey     = "DATABASE_PORT"
	DatabaseTimezoneKey = "DATABASE_TIMEZONE"
	DatabaseSSLModeKey  = "DATABASE_SSL_MODE"
	DatabaseNameKey     = "DATABASE_NAME"
	DatabaseUserKey     = "DATABASE_USER"
	DatabasePasswordKey = "DATABASE_PASSWORD"
)

var envList = []string{
	GrpcPortKey,
	DatabaseHostKey,
	DatabasePortKey,
	DatabaseTimezoneKey,
	DatabaseSSLModeKey,
	DatabaseNameKey,
	DatabaseUserKey,
	DatabasePasswordKey,
}

func BindEnv(envList []string) {
	var err error
	for i := 0; i < len(envList); i++ {
		if err = viper.BindEnv(envList[i]); err != nil {
			log.Fatalf("%s: %v", "failed to binding app environment", err)
		}
	}
}

func EnvList() []string {
	return envList
}
