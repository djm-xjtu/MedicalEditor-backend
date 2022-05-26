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

	template := router.Group("/record-template")
	{
		template.POST("", handlers.CreateRecordTemplate)
		template.GET("", handlers.GetRecordTemplate)
	}

	record := router.Group("/medical-record")
	{
		record.POST("/update", handlers.UpdateMedicalRecord)
		record.POST("/insert", handlers.InsertMedicalRecord)
	}

	outpatient := router.Group("/outpatient")
	{
		outpatient.GET("/login", handlers.Login)
		outpatient.GET("/logout", handlers.Logout)
	}



	router.GET("/medical-record", handlers.GetMedicalRecord)
	return router
}
