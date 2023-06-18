package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jcarlospontes/go-crud-api/initializers"
	"github.com/jcarlospontes/go-crud-api/models"
)

func Cadastrar(c *gin.Context) {

	var body struct {
		Musica string
		Autor  string
		Genero string
	}

	c.Bind(&body)

	music := models.Music{Musica: body.Musica, Autor: body.Autor, Genero: body.Genero}

	result := initializers.DB.Create(&music)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"music": music,
	})
}

func Visualizar(c *gin.Context) {
	var musics []models.Music
	initializers.DB.Find(&musics)

	c.JSON(200, musics)
}

func Selecionar(c *gin.Context) {

	id := c.Param("id")

	var music models.Music
	initializers.DB.First(&music, id)

	c.JSON(200, gin.H{
		"music": music,
	})
}

func Editar(c *gin.Context) {
	id := c.Param("id")

	var music models.Music
	if err := initializers.DB.First(&music, id).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "Música não encontrada",
		})
		return
	}

	var updatedMusic models.Music
	if err := c.ShouldBindJSON(&updatedMusic); err != nil {
		c.JSON(400, gin.H{
			"error": "Dados inválidos",
		})
		return
	}

	music.Musica = updatedMusic.Musica
	music.Autor = updatedMusic.Autor
	music.Genero = updatedMusic.Genero

	if err := initializers.DB.Save(&music).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Erro ao atualizar a música",
		})
		return
	}

	c.JSON(200, gin.H{
		"music": music,
	})
}

func Deletar(c *gin.Context) {

	id := c.Param("id")

	var music models.Music
	result := initializers.DB.First(&music, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Música não encontrada"}) // Retorna status 404 se a música não for encontrada
		return
	}

	initializers.DB.Delete(&music)

	c.JSON(204, gin.H{})
}
