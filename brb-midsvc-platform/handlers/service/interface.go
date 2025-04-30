package service

import(
	"github.com/gin-gonic/gin"
)

type Handler interface {
	CreateService(c *gin.Context)
	UpdateService(c *gin.Context)
	ToggleService(c *gin.Context)
}