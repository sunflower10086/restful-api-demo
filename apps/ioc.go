package apps

import "github.com/sunflower10086/restful-api-demo/apps/host"

// IOC 容器层

//  1. HostService的实例必须注册过来，HostService才会有具体的实例

// HostService HTTP模块依赖于IOC中的HostService
var (
	HostService host.Service
)
