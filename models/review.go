package models

import "time"

type Review struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Comment   string    `json:"comment"`
	GameID    uint      `json:"gameID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Game      Game      `json:"-"`
}
