package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Author struct (model)
type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Books struct (model)
type Books struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// func (b Books) findByID(id string) Books {
// 	for _, item := range books {
// 		if item.ID == id {
// 			return item
// 		}
// 	}

// 	return Books{}
// }

// init a variable
var books []Books

func main() {
	// init router
	r := mux.NewRouter()

	// Mock Data @todo: add database connection
	books = append(books, Books{"1", "2323H32S", "The 5 AM Club", &Author{"Robin", "Sharma"}})
	books = append(books, Books{"2", "256G213I", "The 6 AM Club", &Author{"John", "Wick"}})

	// router handler / endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id:[0-9]+}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id:[0-9]+}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id:[0-9]+}", deleteBook).Methods("DELETE")

	// Run and serve the apis
	fmt.Println("Server running on 3000.......")
	log.Fatal(http.ListenAndServe(":3000", r))
}

// get All books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(books)
}

// get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Books{})
}

// create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var book Books
	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(1000))
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

// update a single book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Books
			_ = json.NewDecoder(r.Body).Decode(&book)

			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(books)
			return
		}
	}

	json.NewEncoder(w).Encode(&Books{})
}

// delete a single book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}
