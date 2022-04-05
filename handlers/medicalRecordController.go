package handlers

import (
	"editor-backend/services"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RecordInfo struct {
	PatientId  int    `json:"patientId"`
	RecordType string `json:"recordType"`
	Record     string `json:"record"`
}

func UpdateMedicalRecord(c *gin.Context) {
	recordInfo := RecordInfo{}
	c.BindJSON(&recordInfo)
	err := services.UpdateOrInsertMedicalRecord(recordInfo.PatientId, recordInfo.RecordType, recordInfo.Record)

	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "update fail",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	}
}

func GetMedicalRecord(c *gin.Context) {
	patientId, _ := strconv.Atoi(c.Query("patientId"))
	recordType := c.Query("recordType")
	fmt.Printf("patientId: %d recordType: %s\n", patientId, recordType)

	record, err := services.GetMedicalRecord(patientId, recordType)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"record": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"record": record,
		})
	}
}
