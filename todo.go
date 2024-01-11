package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var tasks = []Task{
	{ID: "1", Title: "go to grociery"},
	{ID: "2", Title: "go to barber"},
	{ID: "3", Title: "get the keys"},
}

func getAllTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}
func main() {
	router := gin.Default()
	
	router.GET("/tasks", getAllTasks)

	router.Run("localhost:8080")
}
