package structs

import "time"

type Category struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"create-at"`
	UpdateAt time.Time `json:"update-at"`
}
