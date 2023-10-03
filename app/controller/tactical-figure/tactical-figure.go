package tacticalfigure

import (
	"be-tactical-figure/app/db"
	"be-tactical-figure/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
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

type ErrorResponse struct {
	Message string `json:"message"`
}
