package model

type User struct{
	Base
	Name string `json:"name"`
	Age int		`json:"age"`
	Address string `json:"address"`
	Phone string `json:"phone"`
}
