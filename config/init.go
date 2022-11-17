package config

import (
	"errors"
	"github.com/spf13/viper"
)

// Conf 配置
var Conf *viper.Viper

// 配置类型
const (
	INI  = "ini"
	JSON = "json"
	YAML = "yaml"
)

// NewIniConf 实例化INI类型的配置文件
func NewIniConf(path, fileName string) error {
	return NewConf(path, fileName, INI)
}

// NewJsonConf 实例化JSON类型的配置文件
func NewJsonConf(path, fileName string) error {
	return NewConf(path, fileName, JSON)
}

// NewYamlConf 实例化YAML类型的配置文件
func NewYamlConf(path, fileName string) error {
	return NewConf(path, fileName, YAML)
}

// NewConf 实力化配置文件
func NewConf(path, fileName, style string) (err error) {
	Conf = viper.New()
	Conf.AddConfigPath(path)
	Conf.SetConfigFile(fileName)
	Conf.SetConfigType(style)
	if err = Conf.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = errors.New("profile not found")
			return
		} else {
			err = errors.New("error in profile")
			return
		}
	}
	return
}
