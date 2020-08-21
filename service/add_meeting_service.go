package service

import (
	"fmt"
	"model"
	"serializer"
	"time"
)

type AddMeetingService struct {
	Title     string `form:"title" json:"title" binding:"required,min=2,max=50"`
	StartTime string `form:"start_time" json:"start_time"`
	EndTime   string `form:"end_time" json:"end_time"`
	GroupId   uint   `form:"group_id" json:"group_id"`
}

// AddMeeting 添加会议
func (service *AddMeetingService) AddMeeting() serializer.Response {
	startDay, err := time.ParseInLocation(STANDARD_TIME_FMT, service.StartTime, time.Local)
	fmt.Println(startDay)
	if err != nil {
		fmt.Println("parse meeting startTime err: ", err)
	}
	endDay, err := time.ParseInLocation(STANDARD_TIME_FMT, service.EndTime, time.Local)
	if err != nil {
		fmt.Println("parse meeting startTime err: ", err)
	}
	meeting := model.Meeting{
		Title:     service.Title,
		StartTime: startDay,
		EndTime:   endDay,
		GroupId:   service.GroupId,
	}
	err = meeting.Create()
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Data:  nil,
			Msg:   "会议添加失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildMeeting(meeting),
	}

}
