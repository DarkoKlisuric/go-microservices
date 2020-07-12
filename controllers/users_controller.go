package controllers

import (
	"encoding/json"
	"github.com/DarkoKlisuric/go-microservices/services"
	"net/http"
	"strconv"
)

// curl localhost:8000/users?user_id=1
func GetUser(resp http.ResponseWriter, req *http.Request) {

	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10 ,64)

	if err != nil {
		// Just return the Bad Request to the client
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("user_ud must be a number"))
		return
	}

	user, err := services.GetUser(userId)

	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		//Handle the err and return to the client
		return
	}

	// return user to client
	jsonValue, _:= json.Marshal(user)
	resp.Write(jsonValue)
}