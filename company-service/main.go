package main

import (
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
		log.Fatalf("error retrieving client for push notification\n%v", errr)
	}
	authService := services.NewAuthService(authRepo, conf)

	s := &server.Server{
		Config:         conf,
		AuthRepository: authRepo,
		AuthService:    authService,
	}
	s.Start()
}
