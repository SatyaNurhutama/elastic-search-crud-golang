package route

import (
	"github.com/gin-gonic/gin"
	"github.com/satyanurhutama/elastic-search-crud-golang/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/students/create", handler.CreateStudent)
	router.GET("/students/detail/:id", handler.GetStudent)
	router.PUT("/students/update", handler.UpdateStudent)
	router.DELETE("/students/delete/:id", handler.DeleteStudent)

	return router
}
