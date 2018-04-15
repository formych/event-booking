package model

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"

	"github.com/formych/event-booking/dao"
	"github.com/gin-gonic/gin"
)

var hmacSampleSecret = []byte("hello world")

// SignUp 注册
func SignUp(c *gin.Context) {
	user := &dao.User{}
	err := c.Bind(user)
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "添加新用户失败"})
		return
	}
	if user.Password == "" || user.Email == "" || user.Name == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "缺少参数"})
		return
	}
	// 注册前确认是否已经有了该name和email
	dbUser, err := dao.UserDao.GetByEmail(user.Email)
	if err != nil {
		c.JSON(501, gin.H{"code": 501, "msg": "查询用户失败"})
		return
	}
	if dbUser.Email != "" {
		c.JSON(502, gin.H{"code": 502, "msg": "用户已存在"})
		return
	}
	dbUser, err = dao.UserDao.GetByName(user.Name)
	if err != nil {
		c.JSON(501, gin.H{"code": 501, "msg": "查询用户失败"})
		return
	}
	if dbUser.Email != "" {
		c.JSON(500, gin.H{"code": 500, "msg": "用户已存在"})
		return
	}

	err = dao.UserDao.Add(user)
	if err != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "添加新用户失败"})
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

	dbuser, err := dao.UserDao.GetByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 501, "msg": "网络错误"})
		return
	}
	if user == nil || user.Password != dbuser.Password {
		c.JSON(http.StatusOK, gin.H{"code": 402, "msg": "用户不存在或密码错误"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":  user.Name,
		"email": user.Email,
		"nbf":   time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"exp":   time.Now().Add(2 * time.Minute).Unix(),
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

// Authentication 认证
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Token")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSampleSecret, nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("name", claims["name"])
			c.Set("email", claims["email"])
			fmt.Println(c)
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 2001, "msg": err.Error()})
			c.Abort()
			return
		}
	}
}

// code:
// 	2001    token失效
// 	2002
