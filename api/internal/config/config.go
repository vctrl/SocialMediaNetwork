package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)

type Config struct {
	Server  *ServerCfg  `json:"server"`
	MySQL   *MySQLCfg   `json:"my_sql"`
	Token   *TokenCfg   `json:"token"`
	Logging *zap.Config `json:"logging"`
}

type ServerCfg struct {
	ListenAddr   string        `json:"listen_addr"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
}

type MySQLCfg struct {
	Conn string `json:"conn"`
}

type TokenCfg struct {
	PrivateKeyPath string        `json:"private_key_path"`
	PublicKeyPath  string        `json:"public_key_path"`
	AccessTTL      time.Duration `json:"access_ttl"`
}

func Init() error {
	viper.AddConfigPath("./config.json")
	viper.SetConfigName("config")

	viper.AutomaticEnv()

	return viper.ReadInConfig()
}
