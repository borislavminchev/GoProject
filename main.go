package main

import (
	"example/controller"
	"example/database"

	"github.com/gin-gonic/gin"
)

func init() {
	database.LoadDB()
}

func main() {
	r := gin.Default()

	r.GET("/person", controller.GetAllPeople)
	r.GET("/person/:id", controller.GetPerson)

	r.Run()
}