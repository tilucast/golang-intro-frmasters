package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func home(w http.ResponseWriter, r * http.Request){
	fmt.Println("I am home boi")
}

type PageData  struct{
	PageTitle string
	PageTodos []Todo
}

type Todo struct{
	Title string
	Content string
}

var todos []Todo

func getTodos(w http.ResponseWriter, r * http.Request){
	pageData := PageData{
		PageTitle: "Get todos",
		PageTodos: todos,
	}

	t, err := template.ParseFiles("todos.html")

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, pageData)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/todos", getTodos)
	fmt.Println("Server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
