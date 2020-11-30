package config

type System struct {
	Env           string `mapstructure:"env"`
	Port          int    `mapstructure:"port"`
	DbType        string `mapstructure:"db-type"`
	OssType       string `mapstructure:"oss-type"`
	ResourcePath  string `mapstructure:"resource-path"`
	OssDir        string `mapstructure:"oss-dir"`
	TemplateDir   string `mapstructure:"template-dir"`
	UseMultipoint bool   `mapstructure:"use-multipoint"`
	Version       string `mapstructure:"version"`
	MachineID     int64  `mapstructure:"machine_id"`
	StartTime     string `mapstructure:"start_time"`
}
