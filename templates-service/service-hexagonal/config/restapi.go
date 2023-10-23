package config

type ConfRest struct {
	Address   string `mapstructure:"address"`
	JwtSecret []byte `mapstructure:"jwt_secret"`
}
