package common

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	Uid           int        `json:"uid"`
	Mail          string     `json:"mail,omitempty"`
	Pass          string     `json:"-"`
	EncryptedPass string     `json:"-"`
	Fname         string     `json:"fname"`
	Lname         string     `json:"lname"`
	Birth         CustomDate `json:"birth"`
	Age           int        `json:"age,omitempty"`
	Gender        string     `json:"gender,omitempty"`
	Orientation   string     `json:"orientation,omitempty"`
	Bio           string     `json:"bio,omitempty"`
	AvaID         int        `json:"avaID,omitempty"`
	Latitude      float32    `json:"latitude,omitempty"`
	Longitude     float32    `json:"longitude,omitempty"`
	Interests     []string   `json:"interests,omitempty"`
	Status        string     `json:"-"`
	Rating        int        `json:"rating"`
}

type Notif struct {
	Nid         int    `json:"nid"`
	UidSender   int    `json:"uidSender"`
	UidReceiver int    `json:"uidReceiver"`
	Body        string `json:"body"`
}

type Message struct {
	Mid         int    `json:"nid"`
	UidSender   int    `json:"uidSender"`
	UidReceiver int    `json:"uidReceiver"`
	Body        string `json:"body"`
}

type Device struct {
	Id     int    `json:"id"`
	Uid    int    `json:"uid"`
	Device string `json:"device"`
}

type Interest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Photo struct {
	Pid int    `json:"pid"`
	Uid int    `json:"uid"`
	Src string `json:"src"`
}

type CustomDate struct {
	Time *time.Time
}

func (d *CustomDate) UnmarshalJSON(b []byte) (err error) {
	layout := "2006-01-02"
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return
	}
	date, err := time.Parse(layout, s)
	d.Time = &date
	return err
}

func (d CustomDate) MarshalJSON() ([]byte, error) {
	layout := "2006-01-02"
	if d.Time == nil {
		return []byte("null"), nil
	}
	if d.Time.IsZero() {
		println("zero")
		return []byte(fmt.Sprintf(`"%s"`, time.Now().Format(layout))), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, d.Time.Format(layout))), nil
}
