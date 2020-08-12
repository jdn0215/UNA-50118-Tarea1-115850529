package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
)

func find(x string) int {
	for i, book := range books {
		if x == book.Id {
			return i
		}
	}
	return -1
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	if id == "book" {
		return handleAll(w, r)
	}
	i := find(id)
	if i == -1 {
		return
	}
	dataJSON, err := json.Marshal(books[i])
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJSON)
	return
}

func handleAll(w http.ResponseWriter, r *http.Request) (err error) {
	dataJSON, err := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJSON)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	fmt.Println("recibiendo put")
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	book := Book{}
	json.Unmarshal(body, &book)
	books = append(books, book)
	w.WriteHeader(200)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	fmt.Println("recibiendo POST")
	//Buscando el libro segun id
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	if i == -1 {
		return
	}
	//Consiguinedo los nuevos valores del libro
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	book := Book{}
	json.Unmarshal(body, &book)
	fmt.Println(book)
	//revisando posiciones
	book2 := books[i]
	if book.Id != "" {
		book2.Id = book.Id
	}
	if book.Title != "" {
		book2.Title = book.Title
	}
	if book.Edition != "" {
		book2.Edition = book.Edition
	}
	if book.Copyright != "" {
		book2.Copyright = book.Copyright
	}
	if book.Language != "" {
		book2.Language = book.Language
	}
	if book.Pages != "" {
		book2.Pages = book.Pages
	}
	if book.Author != "" {
		book2.Author = book.Author
	}
	if book.Publisher != "" {
		book2.Publisher = book.Publisher
	}
	//actualizando libro
	books[i] = book2
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	fmt.Println("recibiendo delete")
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	if i == -1 {
		return
	}
	books = append(books[:i], books[i+1:]...)
	w.WriteHeader(200)
	return
}
