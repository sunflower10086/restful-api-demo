package host

import "context"

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
	*Resource
	// 资源独有属性
	*Describe
}

type Vendor int

const (
	PPRIVATE_IDC Vendor = iota
	ALIYUN
	TXYUN
)

type Resource struct {
	Id          string            `json:"id"  validate:"required"`     // 全局唯一Id
	Vendor      Vendor            `json:"vendor"`                      // 厂商
	Region      string            `json:"region"  validate:"required"` // 地域
	CreateAt    int64             `json:"create_at"`                   // 创建时间
	ExpireAt    int64             `json:"expire_at"`                   // 过期时间
	Type        string            `json:"type"  validate:"required"`   // 规格
	Name        string            `json:"name"  validate:"required"`   // 名称
	Description string            `json:"description"`                 // 描述
	Status      string            `json:"status"`                      // 服务商中的状态
	Tags        map[string]string `json:"tags"`                        // 标签
	UpdateAt    int64             `json:"update_at"`                   // 更新时间
	SyncAt      int64             `json:"sync_at"`                     // 同步时间
	Account     string            `json:"account"`                     // 资源的所属账号
	PublicIP    string            `json:"public_ip"`                   // 公网IP
	PrivateIP   string            `json:"private_ip"`                  // 内网IP
}

type Describe struct {
	CPU          int    `json:"cpu" validate:"required"`    // 核数
	Memory       int    `json:"memory" validate:"required"` // 内存
	GPUAmount    int    `json:"gpu_amount"`                 // GPU数量
	GPUSpec      string `json:"gpu_spec"`                   // GPU类型
	OSType       string `json:"os_type"`                    // 操作系统类型，分为Windows和Linux
	OSName       string `json:"os_name"`                    // 操作系统名称
	SerialNumber string `json:"serial_number"`              // 序列号
}

type QueryHostRequest struct {
}

type DescribeHostRequest struct {
}

type UpdateHostRequest struct {
	*Describe
}

type DeleteHostRequest struct {
	Id string
}
