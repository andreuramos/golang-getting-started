package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID		string 	`json:"id"`
	Title	string	`json:"title"`
	Artist	string	`json:"artist"`
	Price	float64	`json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Raining Blood", Artist: "Slayer", Price: 49.99},
	{ID: "2", Title: "Toxicity", Artist: "System of a Down", Price: 25.00},
	{ID: "3", Title: "Master of Puppets", Artist: "Metallica", Price: 39.95},
	{ID: "4", Title: "Battle of Los Angeles", Artist: "Rage Against the Machine", Price: 56.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)

	router.Run("0.0.0.0:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}