package model

import (
	"net/http"
	"strconv"

	"github.com/formych/event-booking/dao"
	"github.com/gin-gonic/gin"
)

// EventBookingList get all by user id
func EventBookingList(c *gin.Context) {
	userIDstr := c.PostForm("user_id")
	userID, _ := strconv.ParseInt(userIDstr, 10, 64)
	eventlist, err := dao.EventBookingDao.GetAll(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": eventlist,
	})
}

// EventBookingAdd add an event booking
func EventBookingAdd(c *gin.Context) {
	userIDStr, _ := c.Get("id")
	userID, _ := strconv.ParseInt(userIDStr.(string), 10, 64)
	eventID, _ := strconv.ParseInt(c.PostForm("event_id"), 10, 64)
	status, _ := strconv.ParseInt(c.PostForm("status"), 10, 64)

	booking := dao.EventBooking{
		UserID:  userID,
		EventID: eventID,
		Status:  int8(status),
	}
	dao.EventBookingDao.Update(&booking)
	c.JSON(200, gin.H{"code": http.StatusOK})
}
