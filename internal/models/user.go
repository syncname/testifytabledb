package models

import "time"

type User struct {
	Email     string
	Name      string
	CreatedAt time.Time
}
