package services

import (
	"errors"
	"fmt"

	"github.com/nickUseGitHub/go-microservices/mvc/domain"
)

var (
	users = map[int64]*domain.User{
		123: {Id: 123, Title: "Mr.", Name: "Anakin", Lastname: "skywalker"},
	}
)

func GetUser(userId int64) (*domain.User, error) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, errors.New(fmt.Sprintf("UserId %d not found", userId))
}
