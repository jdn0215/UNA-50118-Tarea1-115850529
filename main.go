package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	var err error
	readData("books.csv")
	switch request.Method {
	case "GET":
		fmt.Println("GET....")
		err = handleGet(writer, request)
	case "POST":
		fmt.Println("POST....")
		err = handlePost(writer, request)
	case "PUT":
		fmt.Println("PUT....")
		err = handlePut(writer, request)
	case "DELETE":
		fmt.Println("DELETE....")
		err = handleDelete(writer, request)
	}
	writeData("books.csv")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	fmt.Println("Servidor iniciado....")
	http.HandleFunc("/book/", handler)
	http.ListenAndServe(":8080", nil)
}
