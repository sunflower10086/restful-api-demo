package host

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/sunflower10086/restful-api-demo/apps/models"
)

// Service Host 接口定义
type Service interface {
	// CreateHost 创建一个主机实例
	CreateHost(context.Context, *Host) (*Host, error)
	// QueryHost 查询主机列表
	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
	// DescribeHost 查询主机详情
	DescribeHost(context.Context, *DescribeHostRequest) (*HostSet, error)
	// UpdateHost 更新主机信息
	UpdateHost(context.Context, *UpdateHostRequest) (*HostSet, error)
	// DeleteHost 删除一个主机, 比如前端需要展示当前删除主机的信息，这时就需要返回这个对象
	DeleteHost(context.Context, *DeleteHostRequest) (*HostSet, error)
}

type HostSet struct {
	Items *[]Host
	Total int
}

// Host 模型定义
type Host struct {
	// 资源公共属性
	*models.Resource
	// 资源独有属性
	*models.Describe
	validate *validator.Validate
}

func (h *Host) Validate() error {
	return h.validate.Struct(h)
}

func NewHost() *Host {
	return &Host{
		&models.Resource{},
		&models.Describe{},
		validator.New(),
	}
}

type QueryHostRequest struct {
}

type DescribeHostRequest struct {
}

type UpdateHostRequest struct {
	*models.Describe
}

type DeleteHostRequest struct {
	Id string
}
