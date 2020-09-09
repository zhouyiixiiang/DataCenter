package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserId string
	GroupId  int
	UserType  int
	Password  string
}

func usertableName() string {
	return "sub_meetings"
}

func (item *User) Create() error {
	return DB.Table(usertableName()).Create(item).Error
}

func GetUser(id interface{}) (item User, err error) {
	err = DB.Table(subMeetingtableName()).Where("user_id=?", id).Find(&item).Error
	return item, err
}
