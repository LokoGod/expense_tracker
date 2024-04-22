package main

// import "fmt"
import (

	"github.com/LokoGod/expense_tracker/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	router := gin.Default()

	router.GET("/api/v1/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // Listen and serve on 0.0.0.0:8080
}