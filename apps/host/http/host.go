package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/restful-api-demo/apps/host"
)

func (h *handler) createHost(c *gin.Context) {
	ins := host.NewHost()
	if err := c.ShouldBind(&ins); err != nil {
		c.JSON(-1, "error")
		return
	}

	_, err := h.svc.CreateHost(c.Request.Context(), ins)
	if err != nil {
		return
	}

	return
}
