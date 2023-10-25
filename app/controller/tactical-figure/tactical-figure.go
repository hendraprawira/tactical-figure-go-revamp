package tacticalfigure

import (
	"be-tactical-figure/app/db"
	"be-tactical-figure/app/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	zmq4 "github.com/go-zeromq/zmq4"
)

// @Summary Get All Tactical Point Figure
// @Description  Get All Tactical Point Figure
// @Tags Tactical Figure
// @Success 200 {object} models.Point
// @Failure 404 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Router /figure/point [get]
func GetAllPoint(c *gin.Context) {
	var points []models.TacticalFigure
	//database process
	result := db.DB.Where("figure_type = ? AND is_deleted = ?", "Point", false).Find(&points)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": result.Error})
		return
	}
	c.JSON(http.StatusOK, points)

}

// @Summary Get All Tactical Multi Single Line Figure
// @Description  Get All Tactical Single Line Figure
// @Tags Tactical Figure
// @Success 200 {object} models.Point
// @Failure 404 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Router /figure/single [get]
func GetAllSingle(c *gin.Context) {
	var single []models.TacticalFigure
	//database process
	result := db.DB.Where("figure_type = ? AND is_deleted = ?", "Single", false).Find(&single)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": result.Error})
		return
	}
	c.JSON(http.StatusOK, single)

}

// @Summary Get All Tactical Multi Line Figure
// @Description  Get All Tactical Multi Line Figure
// @Tags Tactical Figure
// @Success 200 {object} models.Point
// @Failure 404 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Router /figure/multi [get]
func GetAllMulti(c *gin.Context) {
	var multi []models.TacticalFigure
	//database process
	result := db.DB.Where("figure_type = ? AND is_deleted = ?", "Multi", false).Find(&multi)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": result.Error})
		return
	}
	c.JSON(http.StatusOK, multi)

}

// SSE channel to broadcast messages
var SseChannel = make(chan string)

func CreatePoint(c *gin.Context, pub zmq4.Socket) {
	// Define a struct for the data you want to insert
	var newPoint models.Point
	mockID := os.Getenv("MOCK_ID")
	// Bind the request JSON or form data to the struct
	if err := c.ShouldBindJSON(&newPoint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Convert [] to string needed for insert to db
	newCoord := fmt.Sprintf("%v", newPoint.Coordinates)
	newSaveDb := fmt.Sprintf("%v", newPoint.SaveDB)
	// Insert the newPoint struct into the database
	if newPoint.SaveDB {
		result := db.DB.Create(&models.TacticalFigureInput{
			FigureType:     "Point",
			Coordinates:    newCoord,
			Color:          newPoint.Color,
			Amplifications: newPoint.Amplifications,
			Opacity:        newPoint.Opacity,
			Altitude:       newPoint.Altitude,
			UpdatedAt:      time.Now(),
			IdUnique:       newPoint.IdUnique,
			IsDeleted:      false,
		})

		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": result.Error})
			return
		}
	}

	// Encode the Point to a JSON []byte
	jsonData, err := json.Marshal(newPoint)
	if err != nil {
		fmt.Println("Error encoding Point to JSON:", err)
		return
	}

	// Publish Message
	msgA := zmq4.NewMsgFrom(
		[]byte("Point"),
		[]byte(jsonData),
		[]byte(mockID),
		[]byte(newSaveDb),
	)
	errs := pub.Send(msgA)
	if errs != nil {
		log.Fatal(errs)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Point created successfullyss", "data": newPoint})

}

func CreateSingleLine(c *gin.Context, pub zmq4.Socket) {
	// Define a struct for the data you want to insert
	var newPoint models.SingleLine
	mockID := os.Getenv("MOCK_ID")
	// Bind the request JSON or form data to the struct
	if err := c.ShouldBindJSON(&newPoint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Convert [] to string needed for insert to db
	newCoord := fmt.Sprintf("%v", newPoint.Coordinates)
	newSaveDb := fmt.Sprintf("%v", newPoint.SaveDB)
	// Insert the newPoint struct into the database
	if newPoint.SaveDB {
		result := db.DB.Create(&models.TacticalFigureInput{
			FigureType:     "Single",
			Coordinates:    newCoord,
			Color:          newPoint.Color,
			Amplifications: newPoint.Amplifications,
			Opacity:        newPoint.Opacity,
			Altitude:       newPoint.Altitude,
			UpdatedAt:      time.Now(),
			IdUnique:       newPoint.IdUnique,
			IsDeleted:      false,
		})

		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": result.Error})
			return
		}
	}

	// Encode the Point to a JSON []byte
	jsonData, err := json.Marshal(newPoint)
	if err != nil {
		fmt.Println("Error encoding Point to JSON:", err)
		return
	}

	// Publish Message
	msgA := zmq4.NewMsgFrom(
		[]byte("Single"),
		[]byte(jsonData),
		[]byte(mockID),
		[]byte(newSaveDb),
	)
	errs := pub.Send(msgA)
	if errs != nil {
		log.Fatal(errs)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Point created successfullyss", "data": newPoint})

}

func CreateMultiLine(c *gin.Context, pub zmq4.Socket) {
	// Define a struct for the data you want to insert
	var newPoint models.MultiLine
	mockID := os.Getenv("MOCK_ID")
	// Bind the request JSON or form data to the struct
	if err := c.ShouldBindJSON(&newPoint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Convert [] to string needed for insert to db
	newCoord := fmt.Sprintf("%v", newPoint.Coordinates)
	newSaveDb := fmt.Sprintf("%v", newPoint.SaveDB)
	// Insert the newPoint struct into the database
	if newPoint.SaveDB {
		result := db.DB.Create(&models.TacticalFigureInput{
			FigureType:     "Multi",
			Coordinates:    newCoord,
			Color:          newPoint.Color,
			Amplifications: newPoint.Amplifications,
			Opacity:        newPoint.Opacity,
			Altitude:       newPoint.Altitude,
			UpdatedAt:      time.Now(),
			IdUnique:       newPoint.IdUnique,
			IsDeleted:      false,
		})

		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": result.Error})
			return
		}
	}

	// Encode the Point to a JSON []byte
	jsonData, err := json.Marshal(newPoint)
	if err != nil {
		fmt.Println("Error encoding Point to JSON:", err)
		return
	}

	// Publish Message
	msgA := zmq4.NewMsgFrom(
		[]byte("Multi"),
		[]byte(jsonData),
		[]byte(mockID),
		[]byte(newSaveDb),
	)
	errs := pub.Send(msgA)
	if errs != nil {
		log.Fatal(errs)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Point created successfullyss", "data": newPoint})

}

func ClientSSE(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	sseWriter := c.Writer
	defer func() {
		SseChannel <- "Client disconnected"
	}()

	// Send a welcome message when a new client connects
	// sseWriter.WriteString("data: Welcome!\n\n")
	sseWriter.Flush()

	// Listen for incoming messages on the SSE channel
	for {
		message := <-SseChannel
		sseWriter.WriteString("data: " + message + "\n\n")
		sseWriter.Flush()
		time.Sleep(1 * time.Second)
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}
