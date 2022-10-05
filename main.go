package main

import (
	"fga-practice-rest-api/config"
	"fga-practice-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	var db = config.DBInit()
	var inDB = controllers.InDB{DB: db}

	var router = gin.Default()
	router.GET("/person/:id", inDB.GetPerson)
	router.GET("/persons", inDB.GetPersons)
	router.POST("/person", inDB.CreatePerson)
	router.PUT("/person", inDB.UpdatePerson)
	router.DELETE("/person/:id", inDB.DeletePerson)

	router.Run(":8000")
}
