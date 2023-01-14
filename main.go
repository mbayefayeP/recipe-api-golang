package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type Recipe struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}


var recipes = []Recipe{
	{ID: "1",Title: "title1"},
	{ID: "1",Title: "title1"},
	{ID: "1",Title: "title1"},
}


func getRecipes(c *gin.Context){
   c.IndentedJSON(http.StatusOK,recipes)
}

func main(){
	router  := gin.Default()
    router.GET("/recipes",getRecipes)
    router.Run(":8080")
}