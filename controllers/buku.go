package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quiz-3/database"
	"quiz-3/repository"
	"quiz-3/structs"
	"strconv"
)

func GetBooks(c *gin.Context) {
	var (
		result gin.H
	)

	books, err := repository.GetBooks(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBooks(c *gin.Context) {
	var books structs.Book

	err := c.ShouldBindJSON(&books)
	if err != nil {
		panic(err)
	}

	err = repository.InsertsBooks(database.DbConnection, books)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Category",
	})
}

func UpdateBooks(c *gin.Context) {
	var books structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&books)
	if err != nil {
		panic(err)
	}

	books.ID = int64(id)

	err = repository.UpdateBooks(database.DbConnection, books)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update books",
	})
}

func DeleteBooks(c *gin.Context) {
	var books structs.Book
	id, err := strconv.Atoi(c.Param("id"))

	books.ID = int64(id)

	err = repository.DeleteBooks(database.DbConnection, books)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete book",
	})
}
