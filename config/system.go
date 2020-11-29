package config

type System struct {
	Env           string `mapstructure:"env"`
	Port          int    `mapstructure:"port"`
	DbType        string `mapstructure:"db-type"`
	OssType       string `mapstructure:"oss-type"`
	UseMultipoint bool   `mapstructure:"use-multipoint"`
	Version       string `mapstructure:"version"`
	MachineID     int64  `mapstructure:"machine_id"`
	StartTime     string `mapstructure:"start_time"`
}
