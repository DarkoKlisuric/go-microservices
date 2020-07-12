package services

import (
	"github.com/DarkoKlisuric/go-microservices/mvc/domain"
	"github.com/DarkoKlisuric/go-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.AppError) {
	return domain.GetUser(userId)
}