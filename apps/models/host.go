package models

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

func (*Resource) TableName() string {
	return "resource"
}

type Describe struct {
	ResourceId   string `json:"resource_id" validate:"required"` // 外键名称
	CPU          int    `json:"cpu" validate:"required"`         // 核数
	Memory       int    `json:"memory" validate:"required"`      // 内存
	GPUAmount    int    `json:"gpu_amount"`                      // GPU数量
	GPUSpec      string `json:"gpu_spec"`                        // GPU类型
	OSType       string `json:"os_type"`                         // 操作系统类型，分为Windows和Linux
	OSName       string `json:"os_name"`                         // 操作系统名称
	SerialNumber string `json:"serial_number"`                   // 序列号
}

func (*Describe) TableName() string {
	return "host"
}
