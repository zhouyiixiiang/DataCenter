package api

import (
	"github.com/gin-gonic/gin"
	"serializer"
	"service"
)

func BackendAddMeeting(c *gin.Context) {
	var service service.AddMeetingService
	if err := c.ShouldBind(&service); err == nil {
		res := service.AddMeeting()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ParamErr("", err))
	}
}

func BackendDeleteMeeting(c *gin.Context) {
	var service service.DeleteMeetingService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DeleteMeeting()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ParamErr("", err))
	}
}

func BackendGetMettingList(c *gin.Context) {
	var service service.GetMeetingListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetMeetingList()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ParamErr("", err))
	}
}

func BackendGetMeeting(c *gin.Context) {
	var service service.GetMeetingService
	if err := c.ShouldBind(&service); err == nil {
		_,res := service.GetMeeting()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ParamErr("", err))
	}
}
