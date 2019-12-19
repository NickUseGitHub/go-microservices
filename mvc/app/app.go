package app

import (
	"fmt"
	"net/http"

	"github.com/nickUseGitHub/go-microservices/mvc/controllers"
)

var port int = 8080

func Start() {
	http.HandleFunc("/users", controllers.GetUsers)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
