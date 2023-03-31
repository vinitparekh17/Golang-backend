package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	Router := gin.Default()
	Router.GET("/", getAlbums)
	Router.POST("/add", addAlbums)
	Router.GET("/:id", getById)
	Router.Run()
}

// A handler function to return the items as json format - http-type GET
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// A handler function to add new items - http-type POST
func addAlbums(c *gin.Context) {
	var newItem album
	if err := c.BindJSON(&newItem); err != nil {
		return
	}
	albums = append(albums, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}

// get by id function - http-type GET
func getById(c *gin.Context) {
	id := c.Param("id")
	for _, i := range albums {
		if i.ID == id {
			c.IndentedJSON(http.StatusOK, i)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
