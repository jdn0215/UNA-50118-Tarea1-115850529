package main

import (
	"fmt"
	"os"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	fmt.Println(os.Getenv("PORT"))//-- Resultado daba blanco
	http.HandleFunc("/", handler)
	//http.ListenAndServe(":8080", nil) 
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)

}