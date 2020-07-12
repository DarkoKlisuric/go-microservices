package services

import (
	"github.com/DarkoKlisuric/go-microservices/domain"
	"github.com/DarkoKlisuric/go-microservices/utils"
)

func GetUser(userId int64) (*domain.User, *utils.AppError) {
	return domain.GetUser(userId)
}