package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/kyokomi/emoji/v2"
)

//todo schema object
type Todo struct {
	ID                string `json:"id"`
	Text              string `json:"text"`
	CompletedStatus   string `json:"completedStatus"`
	CreatedDateTime   string `json:"createdDateTime"`
	CompletedDateTime string `json:"completedDateTime"`
}

//todos slice
var todos []Todo

//return list of todos
func getTodos(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	log.Println("Working on endpoint METHOD:GET localhost:8080/api/todos")
	emojiFeature()
	json.NewEncoder(response).Encode(todos)

}

//return single todo by id
func getTodo(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	log.Println("Working on endpoint METHOD:GET localhost:8080/api/todos/" + params["id"])
	emojiFeature()

	for _, item := range todos {
		if item.ID == params["id"] {
			json.NewEncoder(response).Encode(item)
			return
		}

	}

}

// create new todo task
func createTodos(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	log.Println("Working on endpoint METHOD:POST localhost:8080/api/todos")
	emojiFeature()

	var newTodo Todo
	json.NewDecoder(request.Body).Decode(&newTodo)
	newTodo.ID = strconv.Itoa(len(todos) + 1)
	newTodo.CreatedDateTime = time.Now().Format(time.ANSIC)
	newTodo.CompletedStatus = "false"
	todos = append(todos, newTodo)
	json.NewEncoder(response).Encode(newTodo)

}

//make todo as complete
func makeTodoAsComplete(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	log.Println("Working on endpoint METHOD:POST localhost:8080/api/todos/complete/" + params["id"])
	emojiFeature()

	for i, item := range todos {
		if item.ID == params["id"] {
			todos = append(todos[:i], todos[i+1:]...)
			var makeTodoComplete Todo
			json.NewDecoder(request.Body).Decode(&makeTodoComplete)
			makeTodoComplete.ID = params["id"]
			makeTodoComplete.Text = item.Text
			makeTodoComplete.CreatedDateTime = item.CreatedDateTime
			makeTodoComplete.CompletedDateTime = time.Now().Format(time.ANSIC)
			makeTodoComplete.CompletedStatus = "true"
			todos = append(todos, makeTodoComplete)
			json.NewEncoder(response).Encode(makeTodoComplete)
			return
		}
	}

}

//delete specific todo
func deleteTodo(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	log.Println("Working on endpoint METHOD:DELETE localhost:8080/api/todos/" + params["id"])
	emojiFeature()

	for i, item := range todos {
		if item.ID == params["id"] {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	json.NewEncoder(response).Encode(todos)

}

//Implementation of emoji feature
func emojiFeature() {
	// Create a message using a random format.
	emoji.Println(randomFormat())
}

//References for randomizing message https://golang.org/doc/tutorial/random-greeting

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of emoji greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi. Welcome! :partying_face: ",
		"Great to see you! You should used :memo: app for todo task",
		"Cheers :beer: , Well met!",
		"Time to have :pizza: ",
		"Friend gifted you :fries: ",
		"Congrats!! You earned :moneybag: ",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

//main method to call the application functionality

func main() {
	todos = append(todos,
		Todo{ID: "1", Text: "Eating", CompletedStatus: "false", CreatedDateTime: time.Now().Format(time.ANSIC), CompletedDateTime: ""},
		Todo{ID: "2", Text: "Studying", CompletedStatus: "false", CreatedDateTime: time.Now().Format(time.ANSIC), CompletedDateTime: ""},
	)
	//initialize router
	router := mux.NewRouter()

	log.Println("Started listening on localhost:8080")
	emojiFeature()
	//endpoints
	router.HandleFunc("/api/todos", getTodos).Methods("GET")
	router.HandleFunc("/api/todos/{id}", getTodo).Methods("GET")
	router.HandleFunc("/api/todos", createTodos).Methods("POST")
	router.HandleFunc("/api/todos/complete/{id}", makeTodoAsComplete).Methods("POST")
	router.HandleFunc("/api/todos/{id}", deleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}
