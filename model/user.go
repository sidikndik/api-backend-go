package model

type User struct{
	Base
	Name 		string `json:"name" gorm:"type:varchar(256)"`
	Age 		int	   `json:"age" gorm:"type:smallint"`
	Address 	string `json:"address" gorm:"type:varchar(256)"`
	Phone 		string `json:"phone" gorm:"type:varchar(20)"`
}
