package service

import (
	"model"
	"serializer"
)

type DeleteMeetingService struct {
	Id uint `form:"id" json:"id"`
}

// DeleteMeeting 添加会议
func (service *DeleteMeetingService) DeleteMeeting() serializer.Response {
	meeting := model.Meeting{}
	err :=meeting.DeleteById(service.Id)
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Data:  nil,
			Msg:   "删除会议获取失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Msg:"ok",
	}

}
