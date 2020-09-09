package model

func migration(){
	DB.AutoMigrate(&Meeting{})
	DB.AutoMigrate(&SubMeeting{})
	DB.AutoMigrate(&User{})
}
