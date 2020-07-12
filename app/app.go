package app

import (
	"github.com/DarkoKlisuric/go-microservices/controllers"
	"net/http"
)

func Start() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}