package dao

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type userDAO struct {
	tableName  string
	columns    string
	addColumns string
}

// User 表字段
type User struct {
	ID       int64  `db:"id"`
	Name     string `db:"name" form:"name"`
	Tel      string `db:"tel"`
	Email    string `db:"email" form:"email"`
	Password string `db:"password" form:"password"`
	Del      bool   `db:"del" form:"del" json:"del"`
	CreateAt string `db:"create_at"`
	UpdateAt string `db:"update_at"`
}

// UserDao 外部调用的
var UserDao = userDAO{
	tableName:  "user",
	columns:    "id, name, tel, email, password, create_at, update_at",
	addColumns: "name, email, password",
}

func (u userDAO) Add(a *User) (err error) {
	exeSQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES (?, ?, ?)", u.tableName, u.addColumns)
	_, err = DB.Exec(exeSQL, a.Name, a.Email, a.Password)
	if err != nil {
		logrus.Printf("添加用户信息失败, err:[%v]", err)
	}
	return
}

func (u userDAO) GetByEmail(email string) (user User, err error) {
	exeSQL := fmt.Sprintf("SELECT %s FROM %s WHERE email = ?", u.columns, u.tableName)
	//为什么不能能用指针，后续看
	err = DB.Get(&user, exeSQL, email)
	if err != nil {
		logrus.Printf("查询用户信息失败, sql:[%s], email:[%s], err:[%v]", exeSQL, email, err)
	}
	return
}

func (u userDAO) GetByID(id int64) (user User, err error) {
	exeSQL := fmt.Sprintf("SELECT %s FROM %s WHERE id = ?", u.columns, u.tableName)
	_, err = DB.Exec(exeSQL, id)
	if err != nil {
		logrus.Printf("Get user info by id failed, sql:[%s], id:[%d], err:[%v]", exeSQL, id, err)
	}
	return
}

func (u userDAO) GetByName(name string) (user User, err error) {
	exeSQL := fmt.Sprintf("SELECT %s FROM %s WHERE name = ?", u.columns, u.tableName)
	_, err = DB.Exec(exeSQL, name)
	if err != nil {
		logrus.Printf("Get user info by name failed, sql:[%s], name:[%d], err:[%v]", exeSQL, name, err)
	}
	return
}

// func (u userDAO) UpdateByEmail(email string) (user *User, err error) {
// 	exeSQL := fmt.Printf("UPDATE %s SET WHERE id = ?")
// }
