package config

type MySQL struct {
	Addr         string `mapstructure:"addr"`
	Config       string `mapstructure:"config"`
	DbName       string `mapstructure:"db-name"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	MaxIdleConns int    `mapstructure:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode"`
	LogZap       bool   `mapstructure:"log-zap"`
}
