package controllers

import (
	"strconv"

	"github.com/Fabricio-Lima/api-golang/database"
	"github.com/Fabricio-Lima/api-golang/models"
	"github.com/gin-gonic/gin"
)

//listar game por id
func ShowGame(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	var game models.Game

	err = db.First(&game, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find game: " + err.Error(),
		})

		return
	}

	c.JSON(200, game)
}

//adicionar um novo game
func CreateGame(c *gin.Context) {
	db := database.GetDatabase()

	var game models.Game

	err := c.ShouldBindJSON(&game)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind game: " + err.Error(),
		})

		return
	}

	err = db.Create(&game).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create game: " + err.Error(),
		})

		return
	}

	c.JSON(200, game)
}

//listar todos os games
func ShowGames(c *gin.Context) {
	db := database.GetDatabase()

	var games []models.Game
	err := db.Find(&games).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list game: " + err.Error(),
		})

		return
	}

	c.JSON(200, games)
}

func UpdateGame(c *gin.Context) {

	db := database.GetDatabase()

	var game models.Game

	err := c.ShouldBindJSON(&game)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind game: " + err.Error(),
		})

		return
	}

	err = db.Save(&game).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot update game: " + err.Error(),
		})

		return
	}

	c.JSON(200, game)
}

func DeleteGame(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Game{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete book " + err.Error(),
		})

		return
	}

	c.Status(204)
}
