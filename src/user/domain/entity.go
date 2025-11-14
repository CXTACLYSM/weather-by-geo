package domain

import "time"

type User struct {
	id        string
	name      string
	createdAt time.Time
}
