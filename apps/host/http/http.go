package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/restful-api-demo/apps/host"
)

func NewHostHTTPHandler(svc host.Service) *handler {
	return &handler{
		svc: svc,
	}
}

type handler struct {
	svc host.Service
}

func (h *handler) Register(r *gin.Engine) {
	r.POST("/hosts", h.createHost)
}
