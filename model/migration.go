package model

func migration(){
	DB.AutoMigrate(&Meeting{})
}
