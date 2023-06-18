package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jcarlospontes/go-crud-api/controllers"
	"github.com/jcarlospontes/go-crud-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}

	r.Use(cors.New(config))

	r.POST("/", controllers.Cadastrar)
	r.PUT("/", controllers.Editar)
	r.GET("/", controllers.Visualizar)
	r.GET("/:id", controllers.Selecionar)
	r.DELETE("/:id", controllers.Deletar)
	r.Run()
}
