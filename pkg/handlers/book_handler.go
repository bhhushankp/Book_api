package handlers

import (
	"Book_api/pkg/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db = models.Db

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid request Body", http.StatusBadRequest)
		return
		//respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}
	result := db.Create(&book)
	if result.Error != nil {
		log.Printf("error for creating book %v", result.Error)
		http.Error(w, "failled to create book", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, book)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	id := mux.Vars(r)["id"]
	result := db.Find(&book, id)
	if result.Error != nil {
		log.Printf("error for fetching book %v", result.Error)
		http.Error(w, "failled to fetch book", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, book)
}

func GetAllBook(w http.ResponseWriter, r *http.Request) {
	books := []models.Book{}

	result := db.Find(&books)
	if result.Error != nil {
		log.Printf("error for fetching all books %v", result.Error)
		http.Error(w, "failled to fetch all books", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, books)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updatebook := models.Book{}
	id := mux.Vars(r)["id"]

	err := json.NewDecoder(r.Body).Decode(&updatebook)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}
	existingbook := models.Book{}
	result := db.First(&existingbook, id)
	if result.Error != nil {
		log.Printf("error for fetching book for update %v", result.Error)
		http.Error(w, "failled to fetch book for update", http.StatusInternalServerError)
		return
	}
	if updatebook.Name != "" {
		existingbook.Name = updatebook.Name
	}
	if updatebook.Author != "" {
		existingbook.Author = updatebook.Author
	}
	if updatebook.Publication != "" {
		existingbook.Publication = updatebook.Publication
	}

	result = db.Save(&existingbook)
	if result.Error != nil {
		log.Printf("error for updating book %v", result.Error)
		http.Error(w, "failed to update book", http.StatusInternalServerError)

	}

	respondWithJSON(w, http.StatusOK, existingbook)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	id := mux.Vars(r)["id"]
	result := db.First(&book, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}
		log.Printf("erorr checking if the book is exit %v", result.Error)
		http.Error(w, "Failled to check if the book is exits", http.StatusInternalServerError)
		return
	}
	result = db.Delete(&book)
	if result.Error != nil {
		log.Printf("Error deleting book: %v", result.Error)
		http.Error(w, "Failed to delete the book", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Book deleted Succesfullly"})

}

// func respondWithError(w http.ResponseWriter, statusCode int, errorMessage string) {
// 	w.WriteHeader(statusCode)
// 	response := map[string]string{"error": errorMessage}
// 	json.NewEncoder(w).Encode(response)
// }

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
