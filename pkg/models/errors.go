package models


type CustomError struct {
	Message string `json"message"`
	DebugMessage string `json:"-"`
	HttpCode int `json:"-"`
}


func (customError CustomError) Error() string { //Implementing the error response in Golang
	return customError.Message
}


