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
	if len(events) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success", "data": "[]"})
		return
	}
	userIDstring, _ := c.Get("id")
	userID, _ := strconv.ParseInt(userIDstring.(string), 10, 64)
	booking, err := dao.EventBookingDao.GetAll(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "internal error"})
		return
	}
	for _, ev := range events {
		for _, v := range booking {
			if (v.EventID == ev.ID) && (v.Status == 1) {
				ev.Status = 1
			}
		}
		ev.CreateAt = ev.CreateAt.UTC()
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success", "data": events})
}

// EventAdd add a new event
func EventAdd(c *gin.Context) {
	userIDstring, _ := c.Get("id")
	userID, _ := strconv.ParseInt(userIDstring.(string), 10, 64)
	event := &dao.Event{
		CreateBy: userID,
		UpdateBy: userID,
	}
	c.Bind(event)
	err := dao.EventDao.Add(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "add failed"})
		return
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
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success"})
}

// EventUpdate ...
func EventUpdate(c *gin.Context) {
	userIDstring, _ := c.Get("id")
	userID, _ := strconv.ParseInt(userIDstring.(string), 10, 64)
	event := &dao.Event{
		UpdateBy: userID,
	}
	c.Bind(event)
	err := dao.EventDao.Update(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "add failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success"})
}

// EventOne ...
func EventOne(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	event, err := dao.EventDao.GetOne(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": "add failed"})
		return
	}
	if event == nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "no rows"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success", "event": event})
}
