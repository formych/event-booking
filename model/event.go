package model

import (
	"net/http"
	"strconv"

	"github.com/formych/event-booking/dao"
	"github.com/gin-gonic/gin"
)

// EventList get all event
func EventList(c *gin.Context) {
	events, err := dao.EventDao.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "internal error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success", "data": events})
}

// EventAdd add a new event
func EventAdd(c *gin.Context) {
	event := &dao.Event{}
	c.Bind(event)
	err := dao.EventDao.Add(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "add failed"})
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success"})
}

// EventDelete del a event
func EventDelete(c *gin.Context) {
	id := c.Param("id")
	eventID, _ := strconv.ParseInt(id, 10, 64)
	err := dao.EventDao.Delete(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "add failed"})
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success"})
}

// func EventStart
