package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func getAlbums(c *gin.Context) {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return
	}

	var albums []album
	result := db.Find(&albums)

	if result.Error != nil {
		return
	}

	c.IndentedJSON(200, albums)

}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return
	}

	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	db.Create(&newAlbum)
	c.IndentedJSON(http.StatusOK, gin.H{"mensaje": "AGREGADO EXITOSA"})

}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return
	}

	var album album

	err = db.First(&album, id).Error

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, &album)
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return
	}

	var album album
	err = db.Delete(&album, id).Error
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"mensaje": "Eliminacion EXITOSA"})

}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return
	}

	db.AutoMigrate(&album{})

	r := gin.Default()
	r.GET("/albums", getAlbums)
	r.GET("/albums/:id", getAlbumByID)
	r.POST("/albums", postAlbums)
	r.DELETE("/albums/:id", deleteAlbumByID)
	r.Run("localhost:8085") // listen and serve on 0.0.0.0:8080
}
