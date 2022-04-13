package handlers

import (
	"editor-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetData(c *gin.Context) {
	mzghxh := c.Query("mzghxh")
	data, err := services.GetData(mzghxh)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
