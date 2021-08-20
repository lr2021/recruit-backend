package discover

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"github.com/lr2021/recruit-backend/user/discover/endpoints"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	var (
		consulHost = flag.String("consul.host", "", "consul server ip address")
		consulPort = flag.String("consul.port", "", "consul server port")
	)
	flag.Parse()

	//创建日志组件
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	//创建consul客户端对象
	var client consul.Client
	{
		consulConfig := api.DefaultConfig()

		consulConfig.Address = "http://" + *consulHost + ":" + *consulPort
		consulClient, err := api.NewClient(consulConfig)

		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		client = consul.NewClient(consulClient)
	}

	ctx := context.Background()

	discoverEndpoint := endpoints.MakeDiscoverEndpoint(ctx, client, logger)

	// TODO: create discoverEndpoint for all services
	r := NewHTTPHandler(ctx, discoverEndpoint)

	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	//开始监听
	go func() {
		fmt.Println("user client running on port: 8001")
		errc <- http.ListenAndServe(":8001", r)
	}()

	// 开始运行，等待结束
	fmt.Println(<-errc)
}
