package service

import (
	"fmt"
	"model"
)

type GetMeetingService struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (service *GetMeetingService) GetMeeting() (meeting model.Meeting, err error) {
	meeting, err = model.GetMeetingByName(service.Name)
	if err != nil {
		fmt.Println("get meeting err")
		return meeting, err
	}
	return meeting, err
}
