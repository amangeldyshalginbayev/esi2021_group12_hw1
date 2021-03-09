package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var API_URL = "http://172.25.0.2:8080/api/todos"

//var API_URL = "http://localhost:8080/api/todos"

func listTodos() {
	fmt.Println("Listing Todos")
	// response, err := http.Get("http://localhost:8080/api/todos")
	response, err := http.Get(API_URL)
	if err != nil {
		fmt.Printf("The HTTP error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
func createTodo() error {
	var input string
	fmt.Scanln(&input)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type your todo item: ")
	var text string
	if scanner.Scan() {
		text = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return err
	} else {
		fmt.Println("Creating todo Item ", text)
		jsonData := map[string]string{"Text": text}
		jsonValue, _ := json.Marshal(jsonData)
		// var response, err = http.Post("http://localhost:8080/api/todos", "application/json", bytes.NewBuffer(jsonValue))
		var response, err = http.Post(API_URL, "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
	}
	return nil
}

func getTodo() error {
	var input string
	fmt.Scanln(&input)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type your todo item id: ")
	var id string
	if scanner.Scan() {
		id = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return err
	} else {
		// response, err := http.Get("http://localhost:8080/api/todos/" + id)
		response, err := http.Get(API_URL + "/" + id)
		if err != nil {
			fmt.Printf("The HTTP error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
	}
	return nil
}

func completeTodo() error {
	var input string
	fmt.Scanln(&input)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type your completed todo item id: ")
	var id string
	if scanner.Scan() {
		id = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return err
	} else {
		jsonData := map[string]string{"ID": id}
		jsonValue, _ := json.Marshal(jsonData)
		// var response, err = http.Post("http://localhost:8080/api/todos/complete/"+id, "application/json", bytes.NewBuffer(jsonValue))
		var response, err = http.Post(API_URL+"/complete/"+id, "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println("Todo Item " + id + " marked as completed")
			fmt.Println(string(data))
		}
	}
	return nil
}

func deleteTodo() error {
	var input string
	fmt.Scanln(&input)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type todo item id: ")
	var id string
	if scanner.Scan() {
		id = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return err
	} else {
		jsonData := map[string]string{"ID": id}
		jsonValue, _ := json.Marshal(jsonData)
		// var request, err = http.NewRequest(http.MethodDelete, "http://localhost:8080/api/todos/"+id, bytes.NewBuffer(jsonValue))
		var request, err = http.NewRequest(http.MethodDelete, API_URL+"/"+id, bytes.NewBuffer(jsonValue))
		client := &http.Client{}

		response, err := client.Do(request)
		if err != nil {
			fmt.Printf("The HTTP error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println("Deleted Item Id is " + id + ". Remaining todo items are : ")
			fmt.Println(string(data))
		}
	}
	return nil
}

func main() {
	// which api should we call
	fmt.Println("Todo API Console")
	fmt.Println("---------------------")
	fmt.Println("Create Todo -> Type 1")
	fmt.Println("List Todos   -> Type 2")
	fmt.Println("Get Todo   -> Type 3")
	fmt.Println("Complete Todo   -> Type 4")
	fmt.Println("Delete Todo   -> Type 5")

	for {
		fmt.Print("-> ")
		var i int
		fmt.Scan(&i)
		if i == 1 {
			createTodo()
		} else if i == 2 {

			listTodos()
		} else if i == 3 {
			getTodo()
		} else if i == 4 {
			completeTodo()
		} else if i == 5 {
			deleteTodo()
		}

	}

}
