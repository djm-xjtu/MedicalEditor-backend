package handlers

import (
	"editor-backend/internal/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecordInfo struct {
	Mzghxh      string `json:"mzghxh"`
	PatientCdno string `json:"patientCdno"`
	RecordType  string `json:"recordType"`
	Record      string `json:"record"`
	Xm          string `json:"xm"`
	Xb          string `json:"xb"`
	Cssj        string `json:"cssj"`
	Jzks        string `json:"jzks"`
	Tel         string `json:"tel"`
	UpdateBy    string `json:"updateBy"`
	UpdateTime  string `json:"updateTime"`
	ChangeLog   string `json:"changeLog"`
	RecordXml   string `json:"recordXml"`
}

func UpdateMedicalRecord(c *gin.Context) {
	recordInfo := RecordInfo{}
	c.BindJSON(&recordInfo)
	log.Printf("recordInfo %+v\n", recordInfo)
	err := services.UpdateMedicalRecord(recordInfo.Record, recordInfo.Mzghxh, recordInfo.UpdateBy, recordInfo.UpdateTime, recordInfo.ChangeLog, recordInfo.RecordXml)

	if err != nil {
		log.Println(err)
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
	log.Println(recordInfo)
	err := services.InsertMedicalRecord(recordInfo.Mzghxh, recordInfo.PatientCdno, recordInfo.RecordType, recordInfo.Record, recordInfo.Xm, recordInfo.Xb, recordInfo.Cssj, recordInfo.Jzks, recordInfo.Tel, recordInfo.UpdateBy, recordInfo.UpdateTime, recordInfo.ChangeLog, recordInfo.RecordXml)
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
