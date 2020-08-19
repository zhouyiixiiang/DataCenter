package service

import (
	"model"
	"serializer"
	"time"
)

type AddMeetingService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=50"`
	StartTime time.Time `form:"start_time" json:"start_time"`
	EndTime time.Time `form:"end_time" json:"end_time"`
	GroupId int64 `form:"group_id" json:"group_id"`
}

// AddMeeting 添加会议
func (service *AddMeetingService) AddMeeting() serializer.Response {
	meeting :=model.Meeting{
		Title: service.Title,
		StartTime: service.StartTime,
		EndTime: service.EndTime,
		GroupId: service.GroupId,
	}
	err :=meeting.Create()
	if err!=nil{
		return serializer.Response{
			Code:50001,
			Data: nil,
			Msg:"会议添加失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data:
	}

}

