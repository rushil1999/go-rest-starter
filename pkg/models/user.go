package models

type Person struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	UserAccount string `json:"userAccount"`
}
var UserData []Person

type MainTesting struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	UserAccount string `json:"userAccount"`
}



