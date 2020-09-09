package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type SubMeeting struct {
	gorm.Model
	MeetingId uint
	Title     string
	SignInTime time.Time
	ShouldArrivalNum int
	ArrivalShowWords string
	  string
}

func subMeetingtableName() string {
	return "sub_meetings"
}

func (item *SubMeeting) Create() error {
	return DB.Table(subMeetingtableName()).Create(item).Error
}

func GetSubMeetingsByMeetingId(id uint) (list []SubMeeting, err error) {
	err = DB.Table(subMeetingtableName()).Where("meeting_id=?", id).Find(&list).Error
	return list, err
}

func (item *SubMeeting)DeleteById(id uint) ( err error) {
	err = DB.Table(subMeetingtableName()).Where("id=?", id).Delete(&item).Error
	return  err
}


func GetSubMeeting(Id interface{}) (meeting SubMeeting, err error) {
	result := DB.First(&meeting, Id)
	return meeting, result.Error
}
func GetSubMeetingByName(name string) (meeting SubMeeting, err error) {
	err = DB.Where("name=?", name).First(&meeting).Error
	return meeting, err
}
