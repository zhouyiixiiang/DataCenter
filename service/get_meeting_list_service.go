package service

import (
	"model"
	"serializer"
)

type GetMeetingListService struct {
	GroupId uint `form:"group_id" json:"group_id"`
}

// AddMeeting 添加会议
func (service *GetMeetingListService) GetMeetingList() serializer.Response {
	meetings, err := model.GetByGroupId(service.GroupId)
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Data:  nil,
			Msg:   "会议列表获取失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildMeetings(meetings),
	}

}
