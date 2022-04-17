package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapViperConfig is the copy of zap.Config but only with fields which we use
type ZapViperConfig struct {
	Level            zap.AtomicLevel        `json:"level"`
	Encoding         string                 `json:"encoding"`
	EncoderConfig    zapcore.EncoderConfig  `json:"encoderConfig"`
	OutputPaths      []string               `json:"outputPaths"`
	ErrorOutputPaths []string               `json:"errorOutputPaths"`
	InitialFields    map[string]interface{} `json:"initialFields"`
}

type Config struct {
	Server  *ServerCfg     `json:"server"`
	MySQL   *MySQLCfg      `json:"mysql"`
	Token   *TokenCfg      `json:"token"`
	Logging ZapViperConfig `json:"logging"`
}

type ServerCfg struct {
	ListenAddr   string `json:"listen_addr"`
	ReadTimeout  string `json:"read_timeout"`
	WriteTimeout string `json:"write_timeout"`
}

type MySQLCfg struct {
	Conn string `json:"conn"`
}

type TokenCfg struct {
	PrivateKeyPath string `json:"private_key_path"`
	PublicKeyPath  string `json:"public_key_path"`
	AccessTTL      string `json:"access_ttl"`
}

func Init() error {
	viper.AddConfigPath("./internal/config")
	viper.SetConfigName("config")

	viper.AutomaticEnv()

	return viper.ReadInConfig()
}
