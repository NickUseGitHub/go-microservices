package controllers

import (
	"fmt"
	"net/http"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {
	userId := req.URL.Query().Get("userId")
	res.Write([]byte(fmt.Sprintf("User Id: %s", userId)))
}
