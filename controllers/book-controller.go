package controllers

import (
	"net/http"
	"test/db"
	"test/entity"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(ctx *gin.Context) {

	booksCollection := db.GetDBCollection("books")

	cursor, err := booksCollection.Find(ctx, gin.H{})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching all books"})
		return
	}

	var allBooks []entity.Book

	if err = cursor.All(ctx, &allBooks); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while decoding all books"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"books": allBooks,
	})
}

func CreateBook(ctx *gin.Context) {

	var book entity.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	booksCollection := db.GetDBCollection("books")

	result, err := booksCollection.InsertOne(ctx, book)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting book"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book inserted successfully",
		"book":    result,
	})
}
