package models

import "time"

type User struct {
	ID   int32
	Name string
	Dob  time.Time
}
