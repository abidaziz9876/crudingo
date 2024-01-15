package main

import (
	"gomongo/config"
	"gomongo/router"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectToMongoDB()
	config.ConnecToPostresDB()
	config.SetupPostgresSchemaAndTables()

	
	r := gin.Default()
	router.Routes(r)
	log.Fatal(http.ListenAndServe("0.0.0.0:4000",r))
}
