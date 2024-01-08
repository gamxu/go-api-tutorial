package main

import (
	"github/com/gamxu/go-api-tutorial/controllers"
	initializers "github/com/gamxu/go-api-tutorial/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVar()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.PostShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)
	r.Run() // listen and serve on 0.0.0.0:3000
}