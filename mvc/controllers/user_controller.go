package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/nickUseGitHub/go-microservices/mvc/services"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("userId"), 10, 64)
	if err != nil {
		res.Write([]byte("Error woy"))
		return
	}

	log.Println(services.GetUser(userId))
	res.Write([]byte(fmt.Sprintf("User Id naja: %d", userId)))
}
