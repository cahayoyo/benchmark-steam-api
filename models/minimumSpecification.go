package models

import "time"

type MinimumSpecification struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	OS        string    `json:"os"`
	Processor string    `json:"processor"`
	Memory    string    `json:"memory"`
	Graphics  string    `json:"graphics"`
	DirectX   string    `json:"directx"`
	Network   string    `json:"network"`
	Storage   string    `json:"storage"`
	SoundCard string    `json:"soundcard"`
	GameID    uint      `json:"gameID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Game      Game      `json:"-"`
}
