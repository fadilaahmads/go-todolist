package controllers

import (
	"go-todolist/models"
	"go-todolist/services"
	"net/http"

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

}

func Create(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
