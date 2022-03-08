package models

import "time"

type Game struct {
	ID                   uint                   `json:"id" gorm:"primary_key"`
	Title                string                 `json:"title"`
	About                string                 `json:"about"`
	Year                 int                    `json:"year"`
	DeveloperID          uint                   `json:"developer_id"`
	PublisherID          uint                   `json:"publisher_id"`
	GenreID              uint                   `json:"genre_id"`
	AgeRatingCategoryID  uint                   `json:"age_rating_category_id"`
	CreatedAt            time.Time              `json:"created_at"`
	UpdatedAt            time.Time              `json:"updated_at"`
	Developer            Developer              `json:"-"`
	Publisher            Publisher              `json:"-"`
	Genre                Genre                  `json:"-"`
	AgeRatingCategory    AgeRatingCategory      `json:"-"`
	MinimumSpecification []MinimumSpecification `json:"-"`
	MaximumSpecification []MaximumSpecification `json:"-"`
	Review               []Review               `json:"-"`
	Rating               []Rating               `json:"-"`
}
