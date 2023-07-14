package impl

import (
	"context"
	"log"
	"os"

	"github.com/sunflower10086/restful-api-demo/apps"
	"github.com/sunflower10086/restful-api-demo/apps/dao/db"
	"github.com/sunflower10086/restful-api-demo/apps/host"
)

var impl = &HostServiceImpl{}

/*
	之前都是手动把实现的服务注册到IOC层的
	apps.HostService = impl.NewHostServiceImpl()
*/

type HostServiceImpl struct {
	l *log.Logger
}

// 通过匿名引入可以动态注册我们实现的服务
func init() {
	apps.RegistryImpl(impl)
}

func (h *HostServiceImpl) Name() string {
	return apps.AppName
}

func (h *HostServiceImpl) Config() {
	h.l = log.New(os.Stderr, "  [Host] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		l: log.New(os.Stderr, "  [Host] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (h *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	//TODO implement me

	if err := db.CreateHost(ins); err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *HostServiceImpl) QueryHost(ctx context.Context, request *host.QueryHostRequest) (*host.HostSet, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HostServiceImpl) DescribeHost(ctx context.Context, request *host.DescribeHostRequest) (*host.HostSet, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HostServiceImpl) UpdateHost(ctx context.Context, request *host.UpdateHostRequest) (*host.HostSet, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HostServiceImpl) DeleteHost(ctx context.Context, request *host.DeleteHostRequest) (*host.HostSet, error) {
	//TODO implement me
	panic("implement me")
}
