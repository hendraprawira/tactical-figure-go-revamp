package router

import (
	tacticalfigure "be-tactical-figure/app/controller/tactical-figure"
	"be-tactical-figure/app/db"

	docs "be-tactical-figure/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// "github.com/go-zeromq/zmq4"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// // Make Connection to DB
	db.ConnectDatabase()
	// // Make Connection to Memcached
	// db.ConnectMemcached()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "OPTIONS", "GET", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	docs.SwaggerInfo.Title = "BE Tactical Figure API "
	docs.SwaggerInfo.Description = "Tactical Figure"
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Define a route handler for handling HTTP requests
	apiUri := r.Group("/v1")
	apiUri.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	tacticalFigure := apiUri.Group("/figure")
	{
		tacticalFigure.GET("/point", tacticalfigure.GetAllPoint)
		tacticalFigure.GET("/single", tacticalfigure.GetAllSingle)
		tacticalFigure.GET("/multi", tacticalfigure.GetAllMulti)

		tacticalFigure.POST("/point", tacticalfigure.CreatePoint)
		tacticalFigure.POST("/single", tacticalfigure.CreateSingleLine)
		tacticalFigure.POST("/multi", tacticalfigure.CreateMultiLine)

		tacticalFigure.GET("/pointSSE", tacticalfigure.ClientSSE)
	}

	return r
}
