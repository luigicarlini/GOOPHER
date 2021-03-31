package services

import (
	"MICROSERVICES/mvc1/domain"
	"MICROSERVICES/mvc1/utils"
)

//we want to get this user from the database in domain package

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
