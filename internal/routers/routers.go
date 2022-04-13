package routers

import (
	"editor-backend/internal/handlers"

	"editor-backend/internal/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.Cors())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works")
	})

	router.GET("/patientInfos", handlers.GetPatientInfoList)
	router.GET("/patientInfo", handlers.GetPatientInfo)

	router.POST("/record-template", handlers.CreateRecordTemplate)
	router.GET("/record-template", handlers.GetRecordTemplate)
	
	router.POST("/medical-record/update", handlers.UpdateMedicalRecord)
	router.POST("/medical-record/insert", handlers.InsertMedicalRecord)
	router.GET("/medical-record", handlers.GetMedicalRecord)

	router.GET("/outpatient", handlers.GetData)
	return router
}
