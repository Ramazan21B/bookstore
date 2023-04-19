package controller

import (
	"Assignment3/model"
	"Assignment3/repository"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Controller struct {
	DB *repository.Postgres
}

func (c *Controller) GetBooks(w http.ResponseWriter, r *http.Request) {
	books := c.DB.GetBooks()
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
func (c *Controller) GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	book, err := c.DB.GetbyId(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "applixation/json")
	json.NewEncoder(w).Encode(book)
}
func (c *Controller) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := c.DB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}
func (c *Controller) AddBook(w http.ResponseWriter, r *http.Request) {
	var book model.Books
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = c.DB.AddBook(&book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}
func (c *Controller) UpdateBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var book model.Books
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	updatedBook, err := c.DB.UpdateById(id, book.Title, book.Description)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedBook)
}
func (c *Controller) SearchTitle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	books, err := c.DB.SearchByTitle(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func (c *Controller) GetBooksByCostDescending(w http.ResponseWriter, r *http.Request) {
	books := c.DB.GetBooksByCostDescending()
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
func (c *Controller) GetBooksByCostAscending(w http.ResponseWriter, r *http.Request) {
	books := c.DB.GetBooksByCostAscending()
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
