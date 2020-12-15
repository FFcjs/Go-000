package data

import (
	"fmt"

	"Go-000/Week04/internal/myservice/biz"
	"github.com/jmoiron/sqlx"
)

var _ biz.UserRepo = (biz.UserRepo)(nil)

func NewUserRepo(db *sqlx.DB) biz.UserRepo {
	return  &UserRepo{db:db}
}

type UserRepo struct{
	db *sqlx.DB
}

func (ur *UserRepo) CreateUser(u *biz.User){
	fmt.Printf("create user:%s success",u.Name)
}


