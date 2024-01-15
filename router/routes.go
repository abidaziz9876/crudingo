package router

import (
	"gomongo/config"
	"gomongo/controller"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	mongoDB := config.MongoDB
	postgresDB := config.PostGresDB
	collectionName := "data"
	collection := config.GetCollection(mongoDB, collectionName)
	router.GET("/getnames", controller.GetAllDataHandler(collection))
	router.GET("/getname/:id", controller.GetOneDataHandler(collection))
	router.POST("/insert", controller.InsertDataHandler(collection))
	router.PUT("/update/:id", controller.UpdateDataHandler(collection))
	router.DELETE("/delete/:id", controller.DeleteDataByIdHandler(collection))
	router.GET("getpostgres", controller.GetFromPostgresHandler(postgresDB))
	router.POST("/postinpostgres", controller.PostInPostgresHandler(postgresDB))
	router.GET("/getfrompostgres",controller.GetFromPostgresHandler(postgresDB))
}
