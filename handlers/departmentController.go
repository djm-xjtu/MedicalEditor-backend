package handlers

import (
	"editor-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDepartmentList(c *gin.Context) {
	departments, err := services.GetDepartments()
	
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"departments": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"departments": departments,
		})
	}
}