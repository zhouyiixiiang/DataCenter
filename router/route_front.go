package router

import (
	"api"
	"github.com/gin-gonic/gin"
)

func RouteFront(router *gin.Engine) {
	group_router := router.Group("/data_center")
	group_router.POST("meeting_info", api.GetMeetingInfo)
}
