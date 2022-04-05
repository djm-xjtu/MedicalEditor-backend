package routers

import (
	"editor-backend/handlers"
	"editor-backend/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.Cors())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works")
	})

	router.GET("/departments", handlers.GetDepartmentList)
	router.POST("/record-template", handlers.CreateRecordTemplate)
	router.POST("/medical-record", handlers.UpdateMedicalRecord)
	router.GET("/medical-record", handlers.GetMedicalRecord)

	return router
}
