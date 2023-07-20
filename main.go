package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Task struct represents a simple task with an ID, Title, and Status
type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// Mock data for this example (you may replace it with a database later)
var tasks = []Task{
	{ID: 1, Title: "Buy groceries", Status: "Pending"},
	{ID: 2, Title: "Write a blog post", Status: "In Progress"},
	{ID: 3, Title: "Walk the dog", Status: "Completed"},
}


func getTaskByID(c *gin.Context) {
	idParam := c.Param("id")

	// Convert the idParam to an integer
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	// Find the task with the given ID
	var foundTask *Task
	for _, task := range tasks {
		if task.ID == id {
			foundTask = &task
			break
		}
	}

	if foundTask == nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(200, foundTask)
}



func main() {
	router := gin.Default()

	// Define the "GET" route to fetch data by ID
	router.GET("/tasks/:id", getTaskByID)

	router.Run(":8080")
}
