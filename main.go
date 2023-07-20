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

func createTask(c *gin.Context) {
	var newTask Task

	// Bind the request body to the newTask variable
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Assign a new unique ID for the task
	newTask.ID = len(tasks) + 1

	// Append the new task to the tasks slice
	tasks = append(tasks, newTask)

	c.JSON(201, newTask)
}

func updateTask(c *gin.Context) {
	idParam := c.Param("id")

	// Convert the idParam to an integer
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	// Find the task with the given ID
	var foundTask *Task
	for i, task := range tasks {
		if task.ID == id {
			foundTask = &tasks[i]
			break
		}
	}

	if foundTask == nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	// Bind the request body to the foundTask variable for update
	if err := c.ShouldBindJSON(foundTask); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	c.JSON(200, foundTask)
}

func deleteTask(c *gin.Context) {
	idParam := c.Param("id")

	// Convert the idParam to an integer
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	// Find the task with the given ID
	var foundTaskIndex = -1
	for i, task := range tasks {
		if task.ID == id {
			foundTaskIndex = i
			break
		}
	}

	if foundTaskIndex == -1 {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	// Remove the task from the slice
	tasks = append(tasks[:foundTaskIndex], tasks[foundTaskIndex+1:]...)

	c.JSON(200, gin.H{"message": "Task deleted successfully"})
}




func main() {
	router := gin.Default()

	// Define the "GET" route to fetch data by ID
	router.GET("/tasks/:id", getTaskByID)

	// Define the "POST" route to create a new task
	router.POST("/tasks", createTask)

	// Define the "PUT" route to update an existing task
	router.PUT("/tasks/:id", updateTask)

	// Define the "DELETE" route to remove a task
	router.DELETE("/tasks/:id", deleteTask)

	router.Run(":8080")
}
