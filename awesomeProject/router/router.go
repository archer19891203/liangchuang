package router
import (
	. "awesomeProject/api"
	"github.com/gin-gonic/gin"
)
func InitRouter() *gin.Engine {
	router := gin.Default()
	//IndexApi为一个Handler
	router.GET("/", IndexApi)
	router.POST("/person", AddPersonApi)
	router.GET("/persons", GetPersonsApi)
	router.GET("/person/:id", GetPersonApi)
	router.PUT("/person/:id", ModPersonApi)
	router.DELETE("/person/:id", DelPersonApi)
	return router
}