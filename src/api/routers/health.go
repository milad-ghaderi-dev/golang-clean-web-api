package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/milad-ghaderi-dev/golang-clean-web-api/handlers"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()
	r.GET("", handler.Health)
}
