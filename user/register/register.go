package main

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"github.com/pborman/uuid"
	"os"
	"strconv"
)

func Register(consulHost, consulPort, svcHost, svcPort string, logger log.Logger) (registrar sd.Registrar) {
	var client consul.Client
	{
		consulCfg := api.DefaultConfig()
		consulCfg.Address = consulHost + ":" + consulPort
		consulClient, err := api.NewClient(consulCfg)
		if err != nil {
			logger.Log("create consul client error:", err)
			os.Exit(1)
		}

		client = consul.NewClient(consulClient)
	}

	check := api.AgentServiceCheck{
		HTTP:     "http://" + svcHost + ":" + svcPort + "/api/user/health",
		Interval: "10s",
		Timeout:  "1s",
		Notes:    "Consul check service health status.",
	}

	port, _ := strconv.Atoi(svcPort)

	reg := api.AgentServiceRegistration{
		ID:      "lr2021-user" + uuid.New(),
		Name:    "lr2021-user",
		Address: svcHost,
		Port:    port,
		Tags:    []string{"lr2021", "user"},
		Check:   &check,
	}

	// 执行注册
	registrar = consul.NewRegistrar(client, &reg, logger)
	return
}
