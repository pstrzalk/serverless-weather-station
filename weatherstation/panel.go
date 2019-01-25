package weatherstation

import (
	"github.com/gin-gonic/gin"
)

func PanelState() gin.H {
	return gin.H{"status": "processing"}
}
