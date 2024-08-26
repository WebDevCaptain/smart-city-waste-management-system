package models

type SensorData struct {
	SensorID  string `json:"sensor_id" binding:"required"`
	FillLevel int    `json:"fill_level" binding:"required,gte=0,lte=100"`
	Timestamp int64  `json:"timestamp" binding:"required"`
	Location  string `json:"location"`
}
