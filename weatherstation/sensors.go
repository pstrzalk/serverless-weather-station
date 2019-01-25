package weatherstation

import (
	"github.com/gin-gonic/gin"
)

func HumiditySensorState() gin.H {
	return gin.H{"value": "23", "unit": "percent"}
}

func TemperatureSensorState() gin.H {
	return gin.H{"value": "50", "unit": "celsius"}
}
