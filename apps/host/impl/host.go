package impl

import (
	"context"

	"github.com/sunflower10086/restful-api-demo/apps/host"
)

func (h *HostServiceImpl) CreateHost(ctx context.Context, h2 *host.Host) (*host.Host, error) {
	//TODO implement me
	h.l.Println("create func")
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
