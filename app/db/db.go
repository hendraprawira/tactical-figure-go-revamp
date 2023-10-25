package db

import (
	"be-tactical-figure/app/models"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	//define var from env

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbType := os.Getenv("DB_TYPE")
	//switch state for dynamic connection presistent db
	switch dbType {
	case "mysql":
		sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
		db, err = gorm.Open(mysql.Open(sqlInfo), &gorm.Config{})
	case "postgres":
		sqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)
		db, err = gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dbname), &gorm.Config{})
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
	if err != nil {
		return nil, err
	}
	DB = db
	//auto migrate or auto create table to database connection
	DB.AutoMigrate(&models.TacticalFigure{})
	log.Print("Database Connected")
	return db, nil
}

func InsertDBPoint(data *models.Point) {
	jsonData, _ := json.Marshal(data.Coordinates)
	jsonStr := string(jsonData)
	datas := models.TacticalFigureInput{
		FigureType:     "Point",
		IdUnique:       data.IdUnique,
		Coordinates:    jsonStr,
		Color:          data.Color,
		Amplifications: data.Amplifications,
		Opacity:        data.Opacity,
		Altitude:       data.Altitude,
		IsDeleted:      false,
		UpdatedAt:      time.Now(),
	}
	result := DB.Create(datas)
	if result.Error != nil {
		log.Fatal(result.Error)
		return
	}
}

func InsertDBSingle(data *models.SingleLine) {
	jsonData, _ := json.Marshal(data.Coordinates)
	jsonStr := string(jsonData)
	datas := models.TacticalFigure{
		FigureType:     "Single",
		IdUnique:       data.IdUnique,
		Coordinates:    jsonStr,
		Color:          data.Color,
		Amplifications: data.Amplifications,
		Opacity:        data.Opacity,
		Altitude:       data.Altitude,
		IsDeleted:      false,
		UpdatedAt:      time.Now(),
	}
	result := DB.Create(datas)
	if result.Error != nil {
		log.Fatal(result.Error)
		return
	}
}

func InsertDBMulti(data *models.MultiLine) {
	jsonData, _ := json.Marshal(data.Coordinates)
	jsonStr := string(jsonData)
	tacticalFigure := models.TacticalFigure{
		FigureType:     "Multi",
		IdUnique:       data.IdUnique,
		Coordinates:    jsonStr,
		Color:          data.Color,
		Amplifications: data.Amplifications,
		Opacity:        data.Opacity,
		Altitude:       data.Altitude,
		IsDeleted:      false,
		UpdatedAt:      time.Now(),
	}
	result := DB.Create(tacticalFigure)
	if result.Error != nil {
		log.Fatal(result.Error)
		return
	}

}
