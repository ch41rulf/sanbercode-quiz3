package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quiz-3/database"
	"quiz-3/repository"
	"quiz-3/structs"
	"strconv"
)

func GetAllCategories(c *gin.Context) {
	var (
		result gin.H
	)

	category, err := repository.GetAllCategories(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": category,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCategories(c *gin.Context) {
	var category structs.Category

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	err = repository.InsertCategories(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Category",
	})
}

func UpdateCategories(c *gin.Context) {
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	category.ID = int64(id)

	err = repository.UpdateCategories(database.DbConnection, category)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Category",
	})
}

func DeleteCategories(c *gin.Context) {
	var category structs.Category
	id, err := strconv.Atoi(c.Param("id"))

	category.ID = int64(id)

	err = repository.DeleteCategories(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Category",
	})
}

func GetBooksByCategoryIDHandler(c *gin.Context) {
	var (
		result gin.H
	)

	books, err := repository.GetBooksByCategoryIDFromDatabase(database.DbConnection)

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
