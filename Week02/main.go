package main

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Dao struct {
	engine *gorm.DB
}
type Service struct {
	dao *Dao
}

func NewDao(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}

var DBEngine *gorm.DB

func NewService() Service {
	svc := Service{NewDao(DBEngine)}
	return svc
}

// User info
type User struct {
	Id     uint64   `json:"id"`
	Extras struct{} `json:"extras"`
}

//认为时异常的service
func (s *Service) GetUserInfo404(id uint64) (*User, error) {
	user, err := s.dao.QueryUserInfoById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//不认为时异常的service
func (s *Service) GetUserInfo(id uint64) (*User, error) {
	user, err := s.dao.QueryUserInfoById(id)
	//不认为是异常
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Dao query user info
func (d *Dao) QueryUserInfoById(id uint64) (*User, error) {
	//access DB...
	err := sql.ErrNoRows
	return nil, errors.Wrapf(err, "Not Found User:%d", id)
}

func main() {
	svc := NewService()
	//认为是异常
	_, err := svc.GetUserInfo404(1)
	if errors.Is(err, sql.ErrNoRows) {
		//返回404
		fmt.Printf("404 User Not Found %+v\n", err)
	} else {
		//返回500
		fmt.Printf("500 Server Error %+v\n", err)
	}

	//不认为是异常
	_, err = svc.GetUserInfo(1)
	if err != nil {
		//返回500
		fmt.Printf("500 Server Error %+v\n", err)
	}
	//正常返回
	fmt.Printf("200 Success")
}
