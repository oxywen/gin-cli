package config

//根config
type Application struct {
	*System  `mapstructure:"system"`
	*Captcha `mapstructure:"captcha"`
	*MySQL   `mapstructure:"mysql"`
	*Redis   `mapstructure:"redis"`
	*Zap     `mapstructure:"zap"`
}
