package apps

import (
	"fmt"

	"github.com/sunflower10086/restful-api-demo/apps/host"
)

// IOC 容器层

//  1. HostService的实例必须注册过来，HostService才会有具体的实例

// HostService HTTP模块依赖于IOC中的HostService
var (
	HostService host.Service
	services    = map[string]Service{}
)

// Registry 注册服务到Ioc中心
func Registry(service Service) {
	if _, ok := services[service.Name()]; ok {
		panic("Service is registered")
	}

	fmt.Println(service.Name(), service)

	services[service.Name()] = service

	if v, ok := service.(host.Service); ok {
		HostService = v
	}
}

// Init Ioc初始化所有服务
func Init() {
	for _, service := range services {
		service.Config()
	}
}

type Service interface {
	Config()
	Name() string
}
