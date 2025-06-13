package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Monitor    *Monitor     `toml:"monitor" mapstructure:"monitor" json:"monitor"`
	Log        logx.LogConf `toml:"log" mapstructure:"log" json:"log"`
	ProjectCfg Project      `toml:"project" mapstructure:"project" json:"project"`
	DB         DB           `toml:"db" mapstructure:"db" json:"db"`
	Redis      Redis        `toml:"redis" mapstructure:"redis" json:"redis"`
}

type Monitor struct {
	PprofEnable bool  `toml:"pprof_enable" mapstructure:"pprof_enable" json:"pprof_enable"`
	PprofPort   int64 `toml:"pprof_port" mapstructure:"pprof_port" json:"pprof_port"`
}

type Project struct {
	Name string `toml:"name" mapstructure:"name" json:"name"`
}

type DB struct {
	DSN string `toml:"dsn" mapstructure:"dsn" json:"dsn"`
}

type Redis struct {
	Host string `toml:"host" mapstructure:"host" json:"host"`
	Port int64  `toml:"port" mapstructure:"port" json:"port"`
	DB   int    `toml:"db" mapstructure:"db" json:"db"`
}

// UnmarshalConfig unmarshal conifg file
// @params path: the path of config dir
func UnmarshalConfig(configFilePath string) (*Config, error) {
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType("toml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CNFT")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

// UnmarshalCmdConfig unmarshal conifg file
// @params path: the path of config dir
func UnmarshalCmdConfig() (*Config, error) {
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var c Config

	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
