package route

import (
	"github.com/gin-gonic/gin"
)

var (
	r *gin.Engine
)

// NewRoute new a router
func NewRoute() *gin.Engine {
	r = gin.New()
	return r
}
