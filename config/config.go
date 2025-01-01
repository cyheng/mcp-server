package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"log"
)

// Config 是一个封装了 koanf 的结构体
type Config struct {
	koanf *koanf.Koanf
}

// NewConfig 创建一个新的 Config 实例
func NewConfig() *Config {
	k := koanf.New(".")
	return &Config{koanf: k}
}

// Load 加载配置，支持从文件、环境变量和命令行参数加载
func (c *Config) Load(configFile string) {
	// 2. 从环境变量加载配置
	c.koanf.Load(env.Provider("APP_", ".", func(s string) string {
		return s
	}), nil)
	if err := c.koanf.Load(file.Provider(configFile), yaml.Parser()); err != nil {
		log.Fatalf("error loading config file: %v", err)
	}

}

// GetString 获取字符串类型的配置值
func (c *Config) GetString(key string) string {
	return c.koanf.String(key)
}

// GetInt 获取整数类型的配置值
func (c *Config) GetInt(key string) int {
	return c.koanf.Int(key)
}

// GetBool 获取布尔类型的配置值
func (c *Config) GetBool(key string) bool {
	return c.koanf.Bool(key)
}
