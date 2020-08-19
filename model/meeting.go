package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Meeting struct {
	gorm.Model
	Title      string
	StartTime time.Time
	EndTime   time.Time
	GroupId   uint
}
func (item *Meeting)tableName() string {
	return "meeting"
}

func (item *Meeting)Create() error{
	return DB.Table(item.tableName()).Create(item).Error
}

func (item *Meeting)GetByGroupId(id int64)(list []Meeting,err error){
	err =DB.Table(item.tableName()).Where("group_id=?",id).Find(&list).Error
	return
}

func GetMeeting(Id interface{}) (meeting Meeting, err error) {
	result := DB.First(&meeting, Id)
	return meeting, result.Error
}
func GetMeetingByName(name string) (meeting Meeting, err error) {
	err = DB.Where("name=?", name).First(&meeting).Error
	return meeting, err
}
