package entities

import "time"

type Category struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdateAt time.Time
}