package domain

import "time"

type Position struct {
	Lat       float32   `json:"lat"`
	Lon       float32   `json:"lon"`
	UserId    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}