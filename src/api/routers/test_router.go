package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/milad-ghaderi-dev/golang-clean-web-api/api/handlers"
)

func TestRouter(r *gin.RouterGroup) {
	h := handlers.NewTestHandler()

	r.GET("/", h.Test)
	r.GET("/users", h.Users)
	r.GET("/user/:id", h.User)
	r.GET("/user/get-user-by-username/:username", h.UserByUsername)
	r.GET("/user/:id/accounts", h.UserAccounts)

	r.GET("/person", h.BodyBinder)
}
