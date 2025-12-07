package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	// TODO refactor grouping with controller package
	"github.com/developerasun/SignalDash/server/api/health"
	"github.com/developerasun/SignalDash/server/api/indicator"

	"github.com/developerasun/SignalDash/server/config"
	docs "github.com/developerasun/SignalDash/server/docs"
	"github.com/developerasun/SignalDash/server/instance"
	"github.com/developerasun/SignalDash/server/models"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title SignalDash API
// @version 1.0
// @description SignalDash backend API documentation
// @BasePath /
func main() {
	// TODO refactor with `NewXXX` constructor pattern
	options := &config.ViperOptions{
		Filename:  "options",
		ConfigDir: "config",
	}
	v := options.InitConfig()
	port := v.GetString("server.port")

	restApi := gin.Default()
	restApi.SetTrustedProxies(nil)

	db, oErr := gorm.Open(sqlite.Open("sd_app.db"), &gorm.Config{})
	if oErr != nil {
		log.Fatalf("main.go: failed to open sqlite")
	}
	db.AutoMigrate(&models.Indicator{})
	log.Println("main.go: database migrated and opened")

	docs.SwaggerInfo.BasePath = ""
	restApi.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// TODO move to middleware
	restApi.Use(ErrorHandler())
	restApi.Use(gin.Recovery())

	api := restApi.Group("/api")
	api.GET("/health", health.Health)
	api.GET("/indicator", func(ctx *gin.Context) {
		indicator.GetIndicator(ctx, db)
	})

	cronWorker := instance.NewCronWorker(v)
	cronWorker.Run()

	restApi.Run(":" + port)
	log.Println("main.go: router started")
}

// ErrorHandler captures errors and returns a consistent JSON error response
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Step1: Process the request first.

		// Step2: Check if any errors were added to the context
		if len(c.Errors) > 0 {
			// Step3: Use the last error
			err := c.Errors.Last().Err

			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "error!",
			})
		}

		// Any other steps if no errors are found
	}
}
