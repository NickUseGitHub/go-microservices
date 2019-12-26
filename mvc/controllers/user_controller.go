package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nickUseGitHub/go-microservices/mvc/services"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("userId"), 10, 64)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("userId mus be number"))
		return
	}

	user, err := services.GetUser(userId)

	if user == nil || err != nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(err.Error()))
		return
	}

	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
