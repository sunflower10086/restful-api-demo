package impl

import (
	"context"
	"testing"

	host2 "github.com/sunflower10086/restful-api-demo/apps/host"
)

var server *HostServiceImpl

func init() {
	s := NewHostServiceImpl()
	server = s
}

func TestCreate(t *testing.T) {
	host := host2.NewHost()
	host.Name = "test"
	server.CreateHost(context.Background(), host)
}
