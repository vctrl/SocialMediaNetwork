package main

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"github.com/vctrl/social-media-network/api/internal/api"
	"github.com/vctrl/social-media-network/api/internal/config"
	"github.com/vctrl/social-media-network/api/internal/db/mysql"
	"github.com/vctrl/social-media-network/api/internal/model"
	"github.com/vctrl/social-media-network/api/internal/password"
	"github.com/vctrl/social-media-network/api/internal/session"
	"log"
	"os"
	"os/signal"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("failed to init config: %v", err)
	}

	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}

	logger, err := cfg.Logging.Build()
	if err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}

	if err := run(&cfg); err != nil {
		logger.Fatal("failed to start server")
	}
}

func run(cfg *config.Config) error {
	sqlDB, err := mysql.FromConfig(cfg)

	if err != nil {
		return errors.WithMessage(err, "create sql connection from config")
	}

	users := mysql.NewUsersMySQL(sqlDB)
	profiles := mysql.NewProfilesMySQL(sqlDB)
	friends := mysql.NewFriendsMySQL(sqlDB)
	friendRequests := mysql.NewFriendRequestsMySQL(sqlDB)

	sm, err := session.FromConfig(cfg)
	if err != nil {
		return errors.WithMessage(err, "create session manager from config")
	}

	ph := password.NewPasswordHasher()

	model := model.New(users, profiles, friends, friendRequests, sm, ph)

	service := api.New(model)

	router := service.RegisterHTTPEndpoints()

	server := fasthttp.Server{
		Handler:      router.HandleRequest,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
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
