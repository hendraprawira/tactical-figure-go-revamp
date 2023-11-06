package models

import (
	"time"
)

type Point struct {
	Coordinates    []float64 `json:"coordinates"`
	Color          string    `json:"color"`
	Amplifications string    `json:"amplifications"`
	Opacity        int       `json:"opacity"`
	Altitude       float64   `json:"altitude"`
	IdUnique       string    `json:"idUnique"`
	SaveDB         bool      `json:"saveDb"`
}

type SingleLine struct {
	Coordinates    [][]float64 `json:"coordinates"`
	Color          string      `json:"color"`
	Amplifications string      `json:"amplifications"`
	Opacity        int         `json:"opacity"`
	Altitude       float64     `json:"altitude"`
	IdUnique       string      `json:"idUnique"`
	SaveDB         bool        `json:"saveDb"`
}

type MultiLine struct {
	Coordinates    [][][]float64 `json:"coordinates"`
	Color          string        `json:"color"`
	Amplifications string        `json:"amplifications"`
	Opacity        int           `json:"opacity"`
	Altitude       float64       `json:"altitude"`
	IdUnique       string        `json:"idUnique"`
	SaveDB         bool          `json:"saveDb"`
}

// Seeder
type TacticalFigure struct {
	Id             int       `json:"id"  gorm:"primaryKey;type:int;"`
	FigureType     string    `json:"figure_type" binding:"required" gorm:"not null"`
	Coordinates    string    `json:"coordinates" gorm:"not null" binding:"required"`
	Color          string    `json:"color" gorm:"not null" binding:"required"`
	Amplifications string    `json:"amplifications" gorm:"not null" binding:"required"`
	Opacity        int       `json:"opacity" gorm:"not null" binding:"required"`
	Altitude       float64   `json:"altitude" gorm:"not null" binding:"required"`
	UpdatedAt      time.Time `json:"updated_at"`
	IdUnique       string    `json:"id_unique"`
	IsDeleted      bool      `json:"id_deleted"`
}

type TacticalFigureInput struct {
	FigureType     string    `json:"figure_type"`
	Coordinates    string    `json:"coordinates"`
	Color          string    `json:"color" `
	Amplifications string    `json:"amplifications" `
	Opacity        int       `json:"opacity" `
	Altitude       float64   `json:"altitude" `
	UpdatedAt      time.Time `json:"updated_at"`
	IdUnique       string    `json:"id_unique"`
	IsDeleted      bool      `json:"id_deleted"`
}

func (TacticalFigure) TableName() string {
	return "tactical_figures"
}

func (TacticalFigureInput) TableName() string {
	return "tactical_figures"
}
