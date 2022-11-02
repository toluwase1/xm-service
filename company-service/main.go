package main

import (
	"company-service/config"
	"company-service/db"
	"company-service/server"
	"company-service/services"
	"log"
	"net/http"
	"time"
)

func main() {
	http.DefaultClient.Timeout = time.Second * 10
	conf, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	gormDB := db.GetDB(conf)
	authRepo := db.NewAuthRepo(gormDB)
	if err != nil {
		log.Fatalf("error retrieving client for push notification\n%v", err)
	}
	authService := services.NewCompanyService(authRepo, conf)
	s := &server.Server{
		Config:         conf,
		AuthRepository: authRepo,
		AuthService:    authService,
	}
	s.Start()
}
