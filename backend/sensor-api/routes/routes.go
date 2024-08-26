package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/webdevcaptain/scwm-sensor-api/middlewares"
)

func Register(app *gin.Engine) {
	sensor := app.Group("/sensor-data")
	sensor.Use(middlewares.SampleMiddleware)
	sensor.POST("/", handleSensorData)
	sensor.POST("/bulk", handleBulkSensorData)
	sensor.GET("/health", healthCheck)
}
