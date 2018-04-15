package model

import (
	"fmt"
	"net/http"
	"time"

	"github.com/formych/event-booking/dao"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var hmacSampleSecret = []byte("hello world")

// Authentication 认证
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Token")
		if tokenString == "" {
			c.JSON(http.StatusOK, gin.H{"code": 2002, "msg": "token is null"})
			c.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSampleSecret, nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("name", claims["name"])
			c.Set("email", claims["email"])
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 2001, "msg": err.Error()})
			c.Abort()
			return
		}
	}
}

// SignUp 注册
func SignUp(c *gin.Context) {
	user := &dao.User{}
	err := c.Bind(user)
	dao.UserDao.GetByID(1)
	if err != nil {
		c.JSON(200, gin.H{"code": 2003, "msg": "参数转换失败"})
		return
	}
	if user.Password == "" || user.Email == "" || user.Name == "" {
		c.JSON(2004, gin.H{"code": 2004, "msg": "缺少参数"})
		return
	}
	// 注册前确认是否已经有了该name和email
	dbUser, err := dao.UserDao.GetByEmail(user.Email)
	if err != nil {
		c.JSON(200, gin.H{"code": 2005, "msg": "查询用户失败"})
		return
	}
	if len(dbUser) > 0 {
		c.JSON(200, gin.H{"code": 2006, "msg": "user email exists"})
		return
	}

	dbUser, err = dao.UserDao.GetByName(user.Name)
	if err != nil {
		c.JSON(200, gin.H{"code": 2005, "msg": "查询用户失败"})
		return
	}
	if len(dbUser) > 0 {
		c.JSON(200, gin.H{"code": 2007, "msg": "user name exists"})
		return
	}

	err = dao.UserDao.Add(user)
	if err != nil {
		c.JSON(200, gin.H{"code": 20007, "msg": "添加新用户失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "添加成功"})
}

// SignIn 登录
func SignIn(c *gin.Context) {
	user := &dao.User{}
	err := c.Bind(user)
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "网络错误"})
		return
	}

	if user.Password == "" || user.Email == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "缺少邮箱地址或者密码"})
		return
	}

	// dbuser, err := dao.UserDao.GetByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 501, "msg": "网络错误"})
		return
	}
	// if user == nil || user.Password != dbuser.Password {
	// 	c.JSON(http.StatusOK, gin.H{"code": 402, "msg": "用户不存在或密码错误"})
	// 	return
	// }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":  user.Name,
		"email": user.Email,
		"nbf":   time.Date(2018, 01, 01, 00, 0, 0, 0, time.UTC).Unix(),
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		logrus.Printf("get token string failed, err:[%v]", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功", "token": tokenString})
}

// SignOut 退出
func SignOut(c *gin.Context) {
	//jwt的话，前端注销该token
}

// code:
// 	2001    token失效
// 	2002	token为空
// 	2003	用户已存在
//	2004
