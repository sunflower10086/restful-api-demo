package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

// 把conf映射成config对象

// LoadConfigFromToml 从Toml文件中加载
func LoadConfigFromToml(filePath string) error {
	config = NewDefaultConfig()
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load err: %s", err.Error())
	}
	return nil
}

// LoadConfigFromEnv 从环境变量加载
func LoadConfigFromEnv() error {
	config = NewDefaultConfig()
	return env.Parse(config)
}
