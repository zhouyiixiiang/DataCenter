package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"service"
)

func GetMeetingInfo(c *gin.Context) {
	var service service.GetMeetingService
	fmt.Println(service)
}
