package discover

import (
	"context"
	"flag"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"os"
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
}
