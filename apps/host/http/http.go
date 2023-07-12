package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/restful-api-demo/apps"
	"github.com/sunflower10086/restful-api-demo/apps/host"
)

var handler = &Handler{}

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
	if apps.HostService == nil {
		return errors.New("请注册HostService实例")
	}

	h.svc = apps.HostService
	return nil
}

func (h *Handler) RouteRegistry(r *gin.Engine) {
	r.POST("/hosts", h.createHost)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello, world")
	})
}
