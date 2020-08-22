package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Meeting struct {
	gorm.Model
	Title     string
	StartTime time.Time
	EndTime   time.Time
	GroupId   uint
}

func tableName() string {
	return "meetings"
}

func (item *Meeting) Create() error {
	return DB.Table(tableName()).Create(item).Error
}

func GetByGroupId(id uint) (list []Meeting, err error) {
	err = DB.Table(tableName()).Where("group_id=?", id).Find(&list).Error
	return list, err
}

func (item *Meeting)DeleteById(id uint) ( err error) {
	err = DB.Table(tableName()).Where("id=?", id).Delete(&item).Error
	return  err
}


func GetMeeting(Id interface{}) (meeting Meeting, err error) {
	result := DB.First(&meeting, Id)
	return meeting, result.Error
}
func GetMeetingByName(name string) (meeting Meeting, err error) {
	err = DB.Where("name=?", name).First(&meeting).Error
	return meeting, err
}
