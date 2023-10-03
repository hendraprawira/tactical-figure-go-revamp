package models

import (
	"time"
)

type Point struct {
	Coordinates    []float64 `json:"coordinates"`
	Color          string    `json:"color"`
	Amplifications string    `json:"amplifications"`
	Opacity        int       `json:"opacity"`
	Altitude       int       `json:"altitude"`
}

type SingleLine struct {
	Coordinates    [][]float64 `json:"coordinates"`
	Color          string      `json:"color"`
	Amplifications string      `json:"amplifications"`
	Opacity        int         `json:"opacity"`
	Altitude       int         `json:"altitude"`
}

type MultiLine struct {
	Coordinates    [][][]float64 `json:"coordinates"`
	Color          string        `json:"color"`
	Amplifications string        `json:"amplifications"`
	Opacity        int           `json:"opacity"`
	Altitude       int           `json:"altitude"`
}

// Seeder
type TacticalFigure struct {
	FigureType     string    `json:"figure_type" binding:"required" gorm:"not null"`
	Coordinates    string    `json:"coordinates" gorm:"not null" binding:"required"`
	Color          string    `json:"color" gorm:"not null" binding:"required"`
	Amplifications string    `json:"amplifications" gorm:"not null" binding:"required"`
	Opacity        int       `json:"opacity" gorm:"not null" binding:"required"`
	Altitude       int       `json:"altitude" gorm:"not null" binding:"required"`
	UpdatedAt      time.Time `json:"updated_at"`
	IsDeleted      bool      `json:"id_deleted"`
}

func (TacticalFigure) TableName() string {
	return "tactical_figures"
}
