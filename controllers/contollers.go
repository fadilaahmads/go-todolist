package controllers

import (
	"fmt"
	"go-todolist/models"
	"go-todolist/services"
	"go-todolist/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	Service := services.GetAll(c)
	MakeWebResponse := models.WebResponse{
		Message: http.StatusOK,
		Status:  "OK",
		Data:    Service,
	}
	c.JSON(http.StatusOK, MakeWebResponse)
}

func GetOne(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Param Id: ", id)
	idInt, err := strconv.Atoi(id)
	utils.PanicIfError(err)
	Service := services.GetOne(c, idInt)
	MakeWebResponse := models.WebResponse{
		Message: http.StatusOK,
		Status:  "OK",
		Data:    Service,
	}
	c.JSON(http.StatusOK, MakeWebResponse)
}

func Create(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
