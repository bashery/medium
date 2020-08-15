package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Dog struct {
	Name string
	Age  int
}

func main() {

	http.HandleFunc("/", dogFunc)

	http.ListenAndServe(":8000", nil)
}

func dogFunc(w http.ResponseWriter, r *http.Request) {
	dog := &Dog{"boby", 3}
	temp, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Errorf("Error %s", err)
	}
	temp.Execute(os.Stdout, dog)
	temp.Execute(w, dog)

}
