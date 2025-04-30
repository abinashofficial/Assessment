package vendor

import(
	"github.com/gin-gonic/gin"
)

type Handler interface {
	CreateVendor(c *gin.Context)
}