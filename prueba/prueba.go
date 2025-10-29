package prueba

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string `json:"id"`
	TITLE  string `json:"title"`
	YEAR   string `json:"year"`
	AUTHOR string `json:"author"`
}

var albums = []album{
	{ID: "1", TITLE: "Blue Train", YEAR: "1957", AUTHOR: "John Coltrane"},
	{ID: "2", TITLE: "Jeru", YEAR: "1954", AUTHOR: "Gerry Mulligan"},
	{ID: "3", TITLE: "Sarah Vaughan and Clifford Brown", YEAR: "1954", AUTHOR: "Sarah Vaughan"},
}

// c *gin.Context captura la petición del cliente
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
func postAlbums(c *gin.Context) {
	var newAlbum album
	//bind json recibe el json y lo mapea a newAlbum
	c.BindJSON(&newAlbum)
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}
