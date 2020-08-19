package serializer

import (
	"model"
	"time"
)

//Meeting 会议序列化器
type Meeting struct {
	ID         uint `json:"id"`
	Title      string `json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	GroupId   uint    `json:"group_id"`
}

// BuildMeeting 序列化会议
func BuildMeeting(item model.Meeting) Meeting {
	return Meeting{
		ID: item.ID,
		Title: item.Title,
		StartTime: item.StartTime,
		EndTime: item.EndTime,
		GroupId: item.GroupId,
	}
}
//BUildMeetings 序列化会议列表
func BuildMeetings(items []model.Meeting)(meetings []Meeting){
	for _,item:=range items {
		meeting:=BuildMeeting(item)
		meetings=append(meetings, meeting)
	}
	return meetings
}
