package dao

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// EventBooking ...
type EventBooking struct {
	ID       int64     `db:"id" form:"id" json:"id"`
	EventID  int64     `db:"event_id" form:"event:id" json:"event_id"`
	UserID   int64     `db:"user_id" form:"user_id" json:"user_id"`
	Booking  bool      `db:"status" form:"status" json:"status"`
	Del      bool      `db:"del" form:"del" json:"del"`
	CreateAt time.Time `db:"create_at" form:"create_at" json:"create_at"`
	UpdateAt time.Time `db:"update_at" form:"update_at" json:"update_at"`
}

type eventBookingDAO struct {
	tName      string
	columns    string
	addColumns string
}

var eventBookingDao = eventBookingDAO{
	tName:      "event_booking",
	columns:    "id, event_id, user_id, booking, del, create_at, update_at",
	addColumns: "event_id, user_id, booking",
}

func (e eventBookingDAO) Update(booking EventBooking) (err error) {
	exeSQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES (?, ?, ?) on duplicate key UPDATE booking = ?", e.tName, e.addColumns)
	_, err = DB.Exec(exeSQL, booking.EventID, booking.UserID, booking.Booking)
	if err != nil {
		logrus.Printf("Update booking status failed, sql:[%s], EventBooking:[%s], err:[%v]", exeSQL, booking, err)
	}
	return
}
