package dao

import (
	"os"

	_ "github.com/go-sql-driver/mysql" // 初始化
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// DB 全局连接
var DB *sqlx.DB

func init() {
	db, err := sqlx.Connect("mysql", "remote:qwe1234@tcp(112.74.32.54:3306)/test")
	if err != nil {
		logrus.Infof("建立数据库连接失败", err)
		os.Exit(-1)
	}
	DB = db
}
