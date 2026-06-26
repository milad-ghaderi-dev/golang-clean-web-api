package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandler struct{}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Test handler is working",
	})
}

func (h *TestHandler) Users(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Users handler is working",
	})
}

func (h *TestHandler) User(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "User handler is working",
	})
}

func (h *TestHandler) UserByUsername(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "User by username handler is working",
	})
}

func (h *TestHandler) UserAccounts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "User accounts handler is working",
	})
}

