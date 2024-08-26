package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/webdevcaptain/scwm-sensor-api/models"
	"github.com/webdevcaptain/scwm-sensor-api/queues"
)

func handleSensorData(context *gin.Context) {
	var data models.SensorData

	// Bind and validate request body
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Marshal the sensor data to json
	message, err := json.Marshal(data)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to marshall JSON",
		})
		return
	}

	// Publish the message to MQ
	if err := queues.PublishMessage(message); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to publish message to MQ",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "Sensor data received sucessfully",
	})
}

func handleBulkSensorData(context *gin.Context) {
	var data []models.SensorData

	// Bind and validate request body
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, sensorData := range data {
		// Marshal each sensor datapoint to JSON
		msg, err := json.Marshal(sensorData)

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to marshall JSON",
			})
			return
		}

		// Publish the message to MQ
		if err := queues.PublishMessage(msg); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to publish message to MQ",
			})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "Bulk sensor data received sucessfully",
	})
}

func healthCheck(context *gin.Context) {
	reqId := context.GetInt("reqId")

	if queues.Ch.IsClosed() {
		context.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "API is down",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "Sensor API running properly",
		"reqId":  reqId,
	})
}
