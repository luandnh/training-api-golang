package service

import (
	"context"
	"net/http"
	"training-api-golang/repository"
)

type (
	IUser interface {
		GetUsers(c context.Context, limit, offset int) (int, interface{})
	}
	User struct {
	}
)

func NewUser() IUser {
	return &User{}
}

func (u *User) GetUsers(c context.Context, limit, offset int) (int, interface{}) {
	users, total, err := repository.UserRepo.SelectUsers(c, limit, offset)
	if err != nil {
		return http.StatusServiceUnavailable, map[string]interface{}{
			"error": err,
		}
	}
	return http.StatusOK, map[string]interface{}{
		"data":  users,
		"total": total,
	}
}
