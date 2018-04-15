package main

import (
	"github.com/formych/event-booking/dao"
	"github.com/formych/event-booking/model"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Static("/", "public")

	u := r.Group("/api/user")
	{
		u.POST("/signin", model.SignIn)
		u.POST("/signup", model.SignUp)
		u.POST("/signout", model.SignOut)
	}

	e := r.Group("/api/event")
	e.Use(model.Authentication())
	{
		e.POST("/add", model.EventAdd)
		e.DELETE("/delete", model.EventDelete)
	}

	r.Run(":8080")

	defer dao.DB.Close()
}
