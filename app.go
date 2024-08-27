package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func Index(rw http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("templates/index.html")

	template.Execute(rw, nil)

}

func main() {
	http.HandleFunc("/", Index)

	// creando el servidor
	fmt.Println("Johansito El servidor esta corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000")
	http.ListenAndServe("localhost:3000", nil)
}
