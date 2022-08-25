package pkg

import (
	"github.com/spf13/viper"
	"log"
)

// LoadConfig 加载服务配置
func LoadConfig(path string, c any) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Failed to load a conf file: %s", err)
	}

	if err := v.Unmarshal(c); err != nil {
		log.Fatalf("Failed to unmarshal conf: %s", err)
	}
}
