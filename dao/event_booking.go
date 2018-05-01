package dao

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// EventBooking ...
type EventBooking struct {
	ID       int64     `db:"id" form:"id" json:"id"`
	EventID  int64     `db:"event_id" form:"event_id" json:"event_id"`
	UserID   int64     `db:"user_id" form:"user_id" json:"user_id"`
	Status   int8      `db:"status" form:"status" json:"status"`
	CreateAt time.Time `db:"create_at" form:"create_at" json:"create_at"`
	UpdateAt time.Time `db:"update_at" form:"update_at" json:"update_at"`
}

type eventBookingDAO struct {
	tName      string
	columns    string
	addColumns string
}

// EventBookingDao ...
var EventBookingDao = eventBookingDAO{
	tName:      "event_booking",
	columns:    "id, event_id, user_id, `status`, create_at, update_at",
	addColumns: "event_id, user_id, `status`",
}

func (e eventBookingDAO) Update(booking *EventBooking) (err error) {
	exeSQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES (?, ?, ?) on duplicate key UPDATE `status` = %d", e.tName, e.addColumns, booking.Status)
	_, err = DB.Exec(exeSQL, booking.EventID, booking.UserID, booking.Status)
	if err != nil {
		logrus.Errorf("Update booking status failed, sql:[%s], EventBooking:[%d], err:[%v]", exeSQL, booking, err)
	}
	return
}

func (e eventBookingDAO) GetAll(userID int64) (eventBooings []*EventBooking, err error) {
	exeSQL := fmt.Sprintf("SELECT %s FROM %s WHERE user_id = ?", e.columns, e.tName)
	err = DB.Select(&eventBooings, exeSQL, userID)
	if err != nil {
		logrus.Errorf("Get all event booking by user id failed, sql:[%s], user_id:[%d], err:[%s]", exeSQL, userID, err.Error())
	}
	return
}
