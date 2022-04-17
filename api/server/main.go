package main

import (
	"bytes"
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"github.com/vctrl/social-media-network/api/internal/api"
	"github.com/vctrl/social-media-network/api/internal/config"
	"github.com/vctrl/social-media-network/api/internal/db/mysql"
	"github.com/vctrl/social-media-network/api/internal/model"
	"github.com/vctrl/social-media-network/api/internal/password"
	"github.com/vctrl/social-media-network/api/internal/session"
	"github.com/xhit/go-str2duration/v2"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"reflect"
)

type server fasthttp.Server

func (s *server) fromConfig(handler func(ctx *fasthttp.RequestCtx), cfg *config.Config) (*fasthttp.Server, error) {
	s.Handler = handler
	wt, err := str2duration.ParseDuration(cfg.Server.WriteTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse write timeout")
	}

	s.WriteTimeout = wt
	rt, err := str2duration.ParseDuration(cfg.Server.ReadTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse read timeout")
	}

	s.ReadTimeout = rt
	srv := fasthttp.Server(*s)
	return &srv, nil
}

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("failed to init config: %v", err)
	}

	var cfg config.Config
	err := viper.Unmarshal(&cfg, func(cfg *mapstructure.DecoderConfig) { cfg.DecodeHook = StandardJSONUnmarshalHookFunc() })

	if err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}

	zapCfg := zap.Config{
		Level:            cfg.Logging.Level,
		Encoding:         cfg.Logging.Encoding,
		EncoderConfig:    cfg.Logging.EncoderConfig,
		OutputPaths:      cfg.Logging.OutputPaths,
		ErrorOutputPaths: cfg.Logging.ErrorOutputPaths,
		InitialFields:    cfg.Logging.InitialFields,
	}

	logger, err := zapCfg.Build()
	if err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}

	if err = run(&cfg); err != nil {
		logger.Fatal("failed to start server", zap.String("error", err.Error()))
	}
}

func run(cfg *config.Config) error {
	sqlDB, err := mysql.FromConfig(cfg)

	if err != nil {
		return errors.WithMessage(err, "invalid mysql config section")
	}

	users := mysql.NewUsersMySQL(sqlDB)
	profiles := mysql.NewProfilesMySQL(sqlDB)
	friends := mysql.NewFriendsMySQL(sqlDB)
	friendRequests := mysql.NewFriendRequestsMySQL(sqlDB)

	sm, err := session.FromConfig(cfg)
	if err != nil {
		return errors.WithMessage(err, "invalid token config section")
	}

	ph := password.NewPasswordHasher()

	model := model.New(users, profiles, friends, friendRequests, sm, ph)

	service := api.New(model)

	router := service.RegisterHTTPEndpoints()

	server, err := (&server{}).fromConfig(router.HandleRequest, cfg)
	if err != nil {
		return errors.WithMessage(err, "invalid server config section")
	}

	errCh := make(chan error)
	go func() {
		if err := server.ListenAndServe(cfg.Server.ListenAddr); err != nil {
			errCh <- errors.WithMessage(err, "server listen and serve")
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	select {
	case <-c:
		server.Shutdown()
	case err = <-errCh:
		server.Shutdown()
		return err
	}

	return nil
}

// StandardJSONUnmarshalHookFunc replacing viper decode with mapstructure on standard JSON decode
// mapstructure doesn't work with anonymous types and types which implement json.Unmarshal
func StandardJSONUnmarshalHookFunc() mapstructure.DecodeHookFuncType {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		raw, err := json.Marshal(data) // кодируем map[string]interface{} в json
		if err != nil {
			return nil, errors.Wrap(err, "marshal")
		}

		out := reflect.New(t).Interface()
		dec := json.NewDecoder(bytes.NewReader(raw))
		dec.DisallowUnknownFields()
		if err = dec.Decode(out); err != nil {
			return nil, errors.Wrap(err, "unmarshal config file content")
		}

		return out, nil
	}
}
