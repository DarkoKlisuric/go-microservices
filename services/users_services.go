package services

import "github.com/DarkoKlisuric/go-microservices/domain"

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}