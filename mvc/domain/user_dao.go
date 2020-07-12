package domain

import (
	"fmt"
	"github.com/DarkoKlisuric/go-microservices/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		1: &User{
			Id:        1,
			FirstName: "Darko",
			LastName:  "Klisurić",
			Email:     "klisuric1995@gmail.com",
		},
	}
)

func GetUser(userId int64) (*User, *utils.AppError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.AppError{
		Message: fmt.Sprintf("user %v was not found", userId),
		Status:  http.StatusNotFound,
		Code:    "not_found",
	}
}
