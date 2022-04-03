package main

import (
	"database/sql"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"github.com/vctrl/SocialNetworkHighload/api/internal/api"
	"github.com/vctrl/SocialNetworkHighload/api/internal/db"
	"github.com/vctrl/SocialNetworkHighload/api/internal/model"
	"github.com/vctrl/SocialNetworkHighload/api/internal/password"
	"github.com/vctrl/SocialNetworkHighload/api/internal/session"
	"io/ioutil"
	"log"
)

const (
	conn           = "root:qwer1234@tcp(localhost:3306)/social-network"
	publicKeyPath  = ""
	privateKeyPath = ""
)

func main() {
	dbConn, err := sql.Open("mysql", conn)

	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	publicKey, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatalf("failed to read public key: %v", err)
	}

	privateKey, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("failed to read private key: %v", err)
	}

	users := db.NewUsersMySQL(dbConn)
	profiles := db.NewProfilesMySQL(dbConn)
	friends := db.NewFriendsMySQL(dbConn)
	friendRequests := db.NewFriendRequestsMySQL(dbConn)

	sm, err := session.NewSessionsJWTManager(publicKey, privateKey)
	ph := password.NewPasswordHasher()

	if err != nil {
		log.Fatalf("failed to init session manager: %v", err)
	}

	model, err := model.New(users, profiles, friends, friendRequests, sm, ph)

	service := api.New(model)

	router := routing.New()

	router.Post("/register", service.HandleRegister)

	router.Post("/login", service.HandleLogin)

	router.Get("/users/<ids>", service.HandleGetUsers)

	router.Put("/users/<id>", service.HandleUpdateUser)

	router.Delete("/users/<id>", service.HandleDeleteUser)

	router.Get("/friends", service.HandleGetFriends)

	router.Get("/friends/requests/sent", service.HandleGetSentRequests)

	router.Get("/friends/requests", service.HandleGetIncomeRequests)

	router.Post("/friends/requests/<id>", service.HandleSendFriendRequest)

	router.Post("/friends/<id>/accept", service.HandleAcceptFriendRequest)

	router.Delete("/friends/<id>", service.HandleDeleteFriend)

	router.Delete("/friends/requests/{id}", service.HandleDeleteFriendRequest)

	router.Group("/api")

	if err = fasthttp.ListenAndServe(":8080", router.HandleRequest); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
