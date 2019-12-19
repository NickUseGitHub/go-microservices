package services

import "github.com/nickUseGitHub/go-microservices/mvc/domain"

func GetUser(userId int64) (domain.User, error) {
	return domain.User{
		Id:       userId,
		Title:    "Ngong",
		Name:     "John",
		Lastname: "Doe",
	}, nil
}
