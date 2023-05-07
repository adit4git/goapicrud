package main

import (
	"fmt"

	"github.com/adit4git/go-crud/controllers"
	"github.com/adit4git/go-crud/initializers"
	"github.com/gin-gonic/gin"
)


func init() {
  
	initializers.LoadEnvVariables()
	initializers.ConnectoDB()
   
}

func main() {
	fmt.Println("Hello")

	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate )
	r.GET("/posts", controllers.PostIndex )
	r.GET("/posts/:id", controllers.PostShow )
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}