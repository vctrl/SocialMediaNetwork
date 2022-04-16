package api

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/vctrl/social-media-network/api/internal/model"
	"net/http"
)

func New(m model.Model) *service {
	return &service{
		model: m,
	}
}

type service struct {
	model model.Model
}

type LoginResponse struct {
	Token string      `json:"token"`
	User  *model.User `json:"user"`
}

type RegisterResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

func (s *service) HandleLogin(ctx *routing.Context) error {
	r := &model.LoginRequest{}
	err := json.Unmarshal(ctx.PostBody(), r)
	if err != nil {
		return errors.Wrap(err, "unmarshal login request")
	}

	user, token, err := s.model.Login(ctx, r)
	if err != nil {
		return errors.WithMessage(err, "login model")
	}

	data, err := json.Marshal(&LoginResponse{
		Token: token,
		User:  user,
	})

	if err != nil {
		return errors.Wrap(err, "marshal login response")
	}

	ctx.Response.SetBody(data)
	ctx.SetStatusCode(http.StatusOK)
	return nil
}

func (s *service) HandleRegister(ctx *routing.Context) error {
	r := &model.RegisterRequest{}
	err := json.Unmarshal(ctx.PostBody(), r)
	if err != nil {
		return errors.Wrap(err, "unmarshal register request")
	}

	id, token, err := s.model.Register(ctx, r)
	if err != nil {
		return errors.WithMessage(err, "register model")
	}

	data, err := json.Marshal(&RegisterResponse{
		ID:    id,
		Token: token,
	})

	if err != nil {
		return errors.Wrap(err, "marshal register response")
	}

	ctx.Response.SetBody(data)
	ctx.SetStatusCode(http.StatusCreated)
	return nil
}

func (s *service) RegisterHTTPEndpoints() *routing.Router {
	router := routing.New()

	router.Post("/register", s.HandleRegister)

	router.Post("/login", s.HandleLogin)

	router.Get("/users/<ids>", s.HandleGetUsers)

	router.Put("/users/<id>", s.HandleUpdateUser)

	router.Delete("/users/<id>", s.HandleDeleteUser)

	router.Get("/friends", s.HandleGetFriends)

	router.Get("/friends/requests/sent", s.HandleGetSentRequests)

	router.Get("/friends/requests", s.HandleGetIncomeRequests)

	router.Post("/friends/requests/<id>", s.HandleSendFriendRequest)

	router.Post("/friends/<id>/accept", s.HandleAcceptFriendRequest)

	router.Delete("/friends/<id>", s.HandleDeleteFriend)

	router.Delete("/friends/requests/{id}", s.HandleDeleteFriendRequest)

	router.Group("/api")

	return router
}

func (s *service) HandleGetUsers(ctx *routing.Context) error {
	return nil
}

func (s *service) HandleUpdateUser(ctx *routing.Context) error {
	return nil
}

func (s *service) HandleDeleteUser(ctx *routing.Context) error {
	return nil
}

func (s *service) HandleGetFriends(ctx *routing.Context) error {
	return nil
}

func (s *service) HandleGetSentRequests(ctx *routing.Context) error {
	return nil
}

func (s *service) HandleGetIncomeRequests(ctx *routing.Context) error {
	return nil
}

func (s *service) HandleSendFriendRequest(ctx *routing.Context) error {
	return nil
}

func (s *service) HandleAcceptFriendRequest(ctx *routing.Context) error {
	return nil
}

func (s *service) HandleDeleteFriend(ctx *routing.Context) error {
	return nil
}

func (s *service) HandleDeleteFriendRequest(ctx *routing.Context) error {
	return nil
}
