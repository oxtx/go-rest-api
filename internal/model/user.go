package model

import "time"

type User struct {
	ID        string
	Email     string
	Name      string
	CreatedAt time.Time
}
