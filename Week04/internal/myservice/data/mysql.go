package data

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

var DbSet = wire.NewSet(DefaultConnectionOpt, NewDb)

type ConnectionOpt struct {
	Drive string
	DNS   string
}

/// DefaultConnectionOpt 提供默认的连接选项
func DefaultConnectionOpt() *ConnectionOpt {
	return &ConnectionOpt{
		Drive: "mysql",
		DNS:   "root:123456@tcp(127.0.0.1:3306)/ADS?charset=utf8mb4&parseTime=True",
	}
}

// NewDb 提供一个Db对象
func NewDb(opt *ConnectionOpt) (*sqlx.DB, func(), error) {
	conn, err := sqlx.Connect(opt.Drive, opt.DNS)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		fmt.Println("cleanup mysql")
		if err := conn.Close(); err != nil {
			log.Println(err)
		}
	}
	return conn, cleanup, nil
}
