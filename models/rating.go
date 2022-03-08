package models

import "time"

type Rating struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Rating    int       `json:"rating"`
	GameID    uint      `json:"gameID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Game      Game      `json:"-"`
}
