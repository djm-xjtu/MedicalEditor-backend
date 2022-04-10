package handlers

import (
	"editor-backend/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecordTemplate struct {
	RecordType string `json:"recordType"`
	Template   string `json:"template"`
}

func CreateRecordTemplate(c *gin.Context) {
	recordTemplate := RecordTemplate{}
	c.BindJSON(&recordTemplate)
	fmt.Println(recordTemplate.RecordType)
	fmt.Println(recordTemplate.Template)

	err := services.InsertMedicalRecordTemplate(recordTemplate.RecordType, recordTemplate.Template)

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

func GetRecordTemplate(c *gin.Context)  {
	recordType := c.Query("recordType")

	template, err := services.GetMedicalRecordTemplate(recordType)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"template": template,
			"ok": true,
		})
	}
}


