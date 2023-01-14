package main

import (
	"time"

	"github.com/gin-gonic/gin"
)


type Recipe struct {
	Name        string      `json:"name"`
	Tags        []string    `json:"tags"`
	Ingredients []string    `json:"ingredients"`
	Intructions []string    `json:"intructions"`
	PublishAt   time.Time   `json:"publishedAt"`
	UpdateAt    time.Time   `json:"updateAt"`
}




func getRecipes(c *gin.Context){}

func main(){
	router  := gin.Default()
    router.GET("/recipes",getRecipes)
    router.Run(":8080")
}