package models

import "time"

type Publisher struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Game      []Game    `json:"-"`
}
