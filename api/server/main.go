package main

import (
	"database/sql"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/vctrl/SocialNetworkHighload/api/internal/api"
	"github.com/vctrl/SocialNetworkHighload/api/internal/db"
	"github.com/vctrl/SocialNetworkHighload/api/internal/model"
	"io/ioutil"
	"log"
)

const (
	conn           = "username:password@tcp(127.0.0.1:3306)/test"
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
	model, err := model.New(users, publicKey, privateKey)

	service := api.New(model)

	router := routing.New()

	router.Post("/users/register", service.HandleRegister)

	// TODO
	//router.Get("/users/<ids>", service.HandleGetUsers)
	//
	//router.Get("/users/<id>", service.HandleGetUser)
	//
	//router.Post("/users/login", service.HandleLogin)
	//
	//router.Put("/users/<id>", service.HandleUpdateUser)
	//
	//router.Delete("/users/<id>", service.HandleDeleteUser)
}
