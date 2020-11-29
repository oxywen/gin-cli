package config

type Captcha struct {
	KeyLong   int `mapstructure:"key-long"`
	ImgWidth  int `mapstructure:"img-width"`
	ImgHeight int `mapstructure:"img-height"`
}
