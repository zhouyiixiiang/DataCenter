package router

import (
	"github.com/gin-gonic/gin"
)

func GroupRouter(engine *gin.Engine) {
	RouteFront(engine)
}
