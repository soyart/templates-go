package config

type ConfPostgres struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DbName   string `mapstructure:"db_name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
