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
	}

	e := r.Group("/api/event")
	e.Use(model.Authentication())
	{
		e.POST("/list", model.EventList)
		e.POST("/list/:id", model.EventOne)
		e.POST("/add", model.EventAdd)
		e.POST("/update", model.EventUpdate)
		e.DELETE("/list/:id", model.EventDelete)
	}

	eb := r.Group("/api/eventbooking")
	eb.Use(model.Authentication())
	{
		eb.POST("/add", model.EventBookingAdd)
	}

	r.Run(":8080")

	defer dao.DB.Close()
}
