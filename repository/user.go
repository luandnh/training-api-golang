package repository

import (
	"context"
	"database/sql"
	"training-api-golang/common/model"
)

type (
	IUser interface {
		SelectUsers(c context.Context, limit, offset int) (*[]model.User, int, error)
	}
	User struct {
	}
)

var UserRepo IUser

func NewUserRepo() IUser {
	if _, err := Db.NewCreateTable().Model((*model.User)(nil)).IfNotExists().Exec(context.Background()); err != nil {
		panic(err)
	}

	return &User{}
}

func (u *User) SelectUsers(c context.Context, limit, offset int) (*[]model.User, int, error) {
	users := new([]model.User)
	query := Db.NewSelect().Model(users).Limit(limit).Offset(offset)
	total, err := query.ScanAndCount(c)
	if err == sql.ErrNoRows {
		return users, 0, nil
	}
	return users, total, err
}
