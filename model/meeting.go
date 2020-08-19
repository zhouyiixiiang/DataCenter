package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Meeting struct {
	gorm.Model
	Name      string
	StartTime time.Time
	EndTime   time.Time
	Comments  string
}

func GetMeeting(Id interface{}) (meeting Meeting, err error) {
	result := DB.First(&meeting, Id)
	return meeting, result.Error
}
func GetMeetingByName(name string) (meeting Meeting, err error) {
	err = DB.Where("name=?", name).First(&meeting).Error
	return meeting, err
}
