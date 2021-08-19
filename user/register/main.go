package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	redis "github.com/lr2021/recruit-backend/general/db/cache"
	db "github.com/lr2021/recruit-backend/general/db/mysql"
	"github.com/lr2021/recruit-backend/user/register/endpoint"
	"github.com/lr2021/recruit-backend/user/register/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	var (
		consulHost  = flag.String("consul.host", "", "consul ip address")
		consulPort  = flag.String("consul.port", "", "consul port")
		serviceHost = flag.String("service.host", "", "service ip address")
		servicePort = flag.String("service.port", "", "service port")
	)
	flag.Parse()

	ctx := context.Background()
	srv := service.NewService()
	errChan := make(chan error)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	go func(){
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()


	endpoints := endpoint.Endpoints{
		Login: endpoint.Login(srv),
		Register: endpoint.Register(srv),
		Logout: endpoint.Logout(srv),
		GetUserProfile: endpoint.GetUserProfile(srv),
		GetUserSolved: endpoint.GetUserSolved(srv),
		UpdateUserProfile: endpoint.UpdateUserProfile(srv),
		GetUserRank: endpoint.GetUserRank(srv),
		HealthCheck: endpoint.HealthCheck(srv),
	}

	registrar := Register(*consulHost, *consulPort, *serviceHost, *servicePort, logger)

	go func(){
		fmt.Println("user service is running on port: ", *servicePort)
		if db.HasInit() {
			db.Init()
		}
		if redis.HasInit() {
			redis.Init()
		}
		handler := NewHTTPServer(ctx, endpoints)
		registrar.Register()
		errChan <- http.ListenAndServe(*servicePort, handler)
	}()

	fmt.Println(<-errChan)
	registrar.Deregister()
}