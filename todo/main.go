package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("Up and Running")
	router := gin.Default() // Create the server
	router.GET("/todo", getTodos)
	router.DELETE("/delete/:id", deleteTodo)
	router.PATCH("/todo/:id", updateToDo)
	router.GET("/todo/:id", getToDo) //stored in parameter called id in the context
	router.POST("/addtodo", AddTodo)
	router.Run("localhost:9090")
}

func AddTodo(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil { // it takes whatever json inside our request body and bind it to this var
		// it will return an error if this json does not fit our structure (todo format)
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos) // transform to json
}

type todo struct {
	ID        string `json: "id"`
	Item      string `json: "item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func getTodoByID(id string) (*todo, error) {
	for i, val := range todos {
		if val.ID == id {
			return &todos[i], nil
		}

	}
	return nil, errors.New("todo Not found")
}
func getIndex(id string) int {
	for i, val := range todos {
		if val.ID == id {
			return i + 1
		}
	}
	return -1
}

func deleteTodo(context *gin.Context) {
	id := context.Param("id")
	n := getIndex(id)
	if n == -1 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo Not found"})
	}
	remove(todos, n)
	context.IndentedJSON(http.StatusOK, todos)
}

func remove(s []todo, i int) []todo {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func updateToDo(context *gin.Context) {

	id := context.Param("id")
	n, err := getTodoByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo with this ID not found"})
	}
	n.Completed = !n.Completed
	context.IndentedJSON(http.StatusOK, n)

}

func getToDo(context *gin.Context) {
	id := context.Param("id") // where the id parameter from url is stored
	n, err := getTodoByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo with this ID not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, n)

}
