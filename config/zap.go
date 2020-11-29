package config

type Zap struct {
	Level         string `mapstructure:"level"`
	Format        string `mapstructure:"format"`
	Prefix        string `mapstructure:"prefix"`
	Director      string `mapstructure:"director"`
	LinkName      string `mapstructure:"link-name"`
	ShowLine      bool   `mapstructure:"show-line"`
	EncodeLevel   string `mapstructure:"encode-level"`
	StacktraceKey string `mapstructure:"stacktrace-key"`
	LogInConsole  bool   `mapstructure:"log-in-console"`
	RotateType    string `mapstructure:"rotate-type"`
	FileName      string `mapstructure:"file-name"`
	MaxSize       int    `mapstructure:"max-size"`
	MaxBackups    int    `mapstructure:"max-backups"`
	MaxAge        int    `mapstructure:"max-age"`
}
