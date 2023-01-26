package main

import (
	"go-todolist/controllers"
	"go-todolist/models"
	"go-todolist/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.WebResponse{
			Message: http.StatusOK,
			Status:  "OK",
			Data: gin.H{
				"title": "Hello World",
			},
		})
	})
	r.GET("/getall", controllers.GetAll)
	r.GET("/getone/:id")
	r.POST("/create")
	r.PATCH("/update/:id")
	r.DELETE("/delete/:id")

	err := r.Run()
	utils.PanicIfError(err)
}
