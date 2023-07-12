package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/restful-api-demo/apps"
	"github.com/sunflower10086/restful-api-demo/apps/host"
)

func NewHostHTTPHandler() *handler {
	return &handler{}
}

type handler struct {
	svc host.Service
}

func (h *handler) Config() error {
	if apps.HostService == nil {
		return errors.New("请注册HostService实例")
	}

	h.svc = apps.HostService
	return nil
}

func (h *handler) RouteRegistry(r *gin.Engine) {
	r.POST("/hosts", h.createHost)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello, world")
	})
}
