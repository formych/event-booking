package dao

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// Event 事件的字段
type Event struct {
	ID       int64     `db:"id" form:"id" json:"id"`
	Name     string    `db:"name" form:"name" json:"name"`
	Price    string    `db:"price" form:"price" json:"price"`
	Dates    string    `db:"dates" form:"dates" json:"dates"`
	Codes    string    `db:"codes" form:"codes" json:"codes"`
	Capacity int64     `db:"capacity" form:"capacity" json:"capacity"`
	Del      bool      `db:"del" form:"del" json:"del"`
	CreateBy int64     `db:"create_by" form:"create_by" json:"create_by"`
	UpdateBy int64     `db:"update_by" form:"update_by" json:"update+by"`
	CreateAt time.Time `db:"create_at" form:"create_at" json:"create_at"`
	UpdateAt time.Time `db:"update_at" form:"update_at" json:"update_at"`
}

type eventDAO struct {
	tName      string
	columns    string
	addColumns string
}

// EventDao ...
var EventDao = eventDAO{
	tName:      "event",
	columns:    "id, name, price, dates, codes, capacity, create_by, update_by, create_at, update_at",
	addColumns: "name, price, dates, codes, capacity, create_by, update_by",
}

func (e eventDAO) Add(event *Event) (err error) {
	exeSQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES (?, ?, ?, ?, ?, ?, ?)", e.tName, e.addColumns)
	_, err = DB.Exec(exeSQL, event.Name, event.Price, event.Dates, event.Codes, event.Capacity, event.CreateBy, event.UpdateBy)
	if err != nil {
		logrus.Errorf("Add event to db failed, sql:[%s], event:[%+v], err:[%v]", exeSQL, *event, err)
	}
	return
}

func (e eventDAO) GetAll() (events []Event, err error) {
	exeSQL := fmt.Sprintf("SELECT %s FROM %s AND del = false", e.columns, e.tName)
	err = DB.Select(&events, exeSQL)
	if err != nil {
		logrus.Errorf("Get all events failed, sql:[%s], err:[%v]", exeSQL, err)
	}
	return
}

func (e eventDAO) Delete(eventID int64) (err error) {
	exeSQL := fmt.Sprintf("UPDATE %s SET del = true WHERE event_id = ?", e.tName)
	_, err = DB.Exec(exeSQL, eventID)
	if err != nil {
		logrus.Errorf("update event del status failed, sql:[%s], eventID:[%d], del:[%t], err:[%v]", exeSQL, eventID, err)
	}
	return
}
