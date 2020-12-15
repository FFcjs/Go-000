package main

import (
	"Go-000/Week04/internal/myservice/biz"
	"Go-000/Week04/internal/myservice/data"
	"github.com/google/wire"
)

//go:generate wire
func NewUserUsecase()(*biz.UserUsecase,func(),error){
	//userSet:=wire.NewSet(data.NewUserRepo,biz.NewUserUsecase)
	panic(wire.Build(biz.NewUserUsecase, data.NewUserRepo, data.DbSet))
	//return biz.UserUsecase{}
}
