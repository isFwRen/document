package config

type Zap struct {
	MaxSize    int  `mapstructure:"max-size" json:"max-size" yaml:"max-size"`
	MaxBackups int  `mapstructure:"max-backups" json:"max-backups" yaml:"max-backups"`
	MaxAge     int  `mapstructure:"max-age" json:"max-age" yaml:"max-age"`
	Develop    bool `mapstructure:"develop" json:"develop" yaml:"develop"`
}
