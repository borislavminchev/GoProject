package controller

import (
	"example/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPeople(c *gin.Context) {
	c.JSON(http.StatusOK, repository.PersonRepository().GetAllPeople())
}

func GetPerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("Error getting person")
	}
	c.JSON(http.StatusOK, repository.PersonRepository().GetPerson(id))
}