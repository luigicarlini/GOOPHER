package domain

import (
	"MICROSERVICES/mvc1/utils"
	"fmt"
	"net/http"
)

var (
	users = map[int64]*User{123: {Id: 123, FirstName: "Luigi", SecondName: "Carlini", Email: "myemail@gmail.com"},
		456: {Id: 456, FirstName: "Joe", SecondName: "Fritzland", Email: "fritzemail@gmail.com"},
		789: {Id: 789, FirstName: "Jimmy", SecondName: "Fontana", Email: "fontanamail@gmail.com"}}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	//return nil, errors.New(fmt.Sprintf("user %v was not found", userId))
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}
}
