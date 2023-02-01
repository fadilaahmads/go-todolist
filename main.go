package main

import (
	"go-todolist/controllers"
	"go-todolist/models"
	"go-todolist/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.SetTrustedProxies([]string{"127.0.0.1"})

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
	r.GET("/getone/:id", controllers.GetOne)
	r.POST("/create")
	r.PATCH("/update/:id")
	r.DELETE("/delete/:id")

	err := r.Run()
	utils.PanicIfError(err)
}
