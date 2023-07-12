package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/restful-api-demo/apps"
	"github.com/sunflower10086/restful-api-demo/apps/host"
)

var handler = &Handler{}

// 自注册到Ioc层
func init() {
	apps.RegistryGin(handler)
}

//func NewHostHTTPHandler() *Handler {
//	return &Handler{}
//}

type Handler struct {
	svc host.Service
}

func (h *Handler) Name() string {
	return apps.AppName
}

func (h *Handler) Config() error {
	hostService := apps.GetImpl(apps.AppName)
	if v, ok := hostService.(host.Service); ok {
		h.svc = v
		return nil
	}

	return fmt.Errorf("%s does not implement the %s Service interface", apps.AppName, apps.AppName)
}

func (h *Handler) RouteRegistry(r *gin.Engine) {
	r.POST("/hosts", h.createHost)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello, world")
	})
}
