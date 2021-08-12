package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/lr2021/recruit-backend/user/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"github.com/lr2021/recruit-backend/user/endpoint"
)

func main(){
	addr := flag.String("http", ":8090", "http listen address")
	flag.Parse()

	ctx := context.Background()

	srv := service.NewService()
	errChan := make(chan error)

	go func(){
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	endpoints := endpoint.Endpoints{
		Login: endpoint.Login(srv),
		Register: endpoint.Register(srv),
		Logout: endpoint.Logout(srv),
		GetUserInformation: endpoint.GetUserInformation(srv),
		GetUserSolves: endpoint.GetUserSolves(srv),
		UpdateUserProfile: endpoint.UpdateUserProfile(srv),
		GetAllUserProfile: endpoint.GetAllUserProfile(srv),
	}

	go func(){
		log.Println("user service is running on port: ", *addr)
		// TODO: check if database is ready
		handler := NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*addr, handler)
	}()

	log.Fatalln(<-errChan)
}