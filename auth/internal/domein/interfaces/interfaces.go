package interfaces

import "github.com/gin-gonic/gin"

type Auth interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	Permissions(c *gin.Context)
}
