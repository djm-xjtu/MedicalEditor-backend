package handlers

import (
	"editor-backend/internal/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecordInfo struct {
	PatientCdno string `json:"patientCdno"`
	RecordType  string `json:"recordType"`
	Record      string `json:"record"`
	RecordNo    int    `json:"recordNo"`
}

func UpdateMedicalRecord(c *gin.Context) {
	recordInfo := RecordInfo{}
	c.BindJSON(&recordInfo)
	log.Printf("recordInfo %+v\n", recordInfo)
	err := services.UpdateMedicalRecord(recordInfo.Record, recordInfo.RecordNo)

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
	patientCdno := c.Query("patientCdno")
	recordType := c.Query("recordType")
	fmt.Printf("patientId: %s recordType: %s\n", patientCdno, recordType)

	record, isRecord, err := services.GetMedicalRecord(patientCdno, recordType)

	if err != nil {
		log.Printf("err: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"record": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"record":   record,
			"isRecord": isRecord,
		})
	}
}

func InsertMedicalRecord(c *gin.Context) {
	recordInfo := RecordInfo{}
	c.BindJSON(&recordInfo)
	err := services.InsertMedicalRecord(recordInfo.PatientCdno, recordInfo.RecordType, recordInfo.Record)

	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "insert fail",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	}
}
