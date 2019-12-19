package controllers

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("userId"), 10, 64)
	if err != nil {
		res.Write([]byte("Error woy"))
		return
	}

	res.Write([]byte(fmt.Sprintf("User Id naja: %d", userId)))
}
