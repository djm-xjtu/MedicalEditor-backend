package handlers

import (
	"editor-backend/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRecordTemplate(c *gin.Context) {
	recordTemplate := services.RecordTemplate{}
	c.BindJSON(&recordTemplate)
	log.Println(recordTemplate)
	err := services.InsertMedicalRecordTemplate(recordTemplate)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "insert fail",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	}
}

func GetRecordTemplate(c *gin.Context) {
	recordType := c.Query("recordType")

	template, err := services.GetMedicalRecordTemplate(recordType)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"template": template,
			"ok":       true,
		})
	}
}
