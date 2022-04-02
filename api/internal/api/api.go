package api

import (
	"encoding/json"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/vctrl/SocialNetworkHighload/api/internal/model"
)

func New(m model.Model) *service {
	return &service{
		model: m,
	}
}

type service struct {
	model model.Model
}

func (s *service) HandleLogin(ctx *routing.Context) error {
	return nil
}

func (s *service) HandleRegister(ctx *routing.Context) error {
	r := &model.RegisterRequest{}

	err := json.Unmarshal(ctx.PostBody(), r)

	if err != nil {
		return err
	}

	err = s.model.Register(ctx, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) HandleGetUser(ctx *routing.Context) error {
	return nil
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
