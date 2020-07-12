package controllers

import (
	"encoding/json"
	"github.com/DarkoKlisuric/go-microservices/mvc/services"
	"github.com/DarkoKlisuric/go-microservices/mvc/utils"
	"net/http"
	"strconv"
)

// curl localhost:8000/users?user_id=1
func GetUser(resp http.ResponseWriter, req *http.Request) {

	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10 ,64)

	if err != nil {

		apiErr := &utils.AppError{
			Message: "user_id must be a number",
			Status:  http.StatusBadRequest,
			Code:    "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)

	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.Status)
		resp.Write(jsonValue)
		return
	}

	jsonValue, _:= json.Marshal(user)
	resp.Write(jsonValue)
}