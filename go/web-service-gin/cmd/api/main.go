package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// album slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	// Initialize a Gin router using Default.
	router := gin.Default()
	// Use the GET function to associate the GET HTTP method and /albums
	// path with a handler function.
	// Note that you’re passing the name of the getAlbums function. This is
	// different from passing the result of the function, which you would
	// do by passing getAlbums() (note the parenthesis).
	router.GET("/albums", getAlbums)
	// Associate the /albums/:id path with the getAlbumByID function. In
	// Gin, the colon preceding an item in the path signifies that the item
	// is a path parameter.
	router.GET("/albums/:id", getAlbumByID)
	// Associate the POST method at the /albums path with the postAlbums
	// function.
	router.POST("/albums", postAlbums)

	// With Gin, you can associate a handler with an HTTP method-and-path
	// combination. In this way, you can separately route requests sent to
	// a single path based on the method the client is using.

	// Use the Run function to attach the router to an http.Server and
	// start the server.
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON and takes
// a gin.Context parameter.
// gin.Context is the most important part of Gin. It carries request
// details, validates and serializes JSON, and more.
func getAlbums(c *gin.Context) {
	// Call Context.IndentedJSON to serialize the struct into JSON and add
	// it to the response.
	// The function’s first argument is the HTTP status code you want to
	// send to the client. Here, you’re passing the StatusOK constant from
	// the net/http package to indicate 200 OK.
	c.IndentedJSON(http.StatusOK, albums)

	// Note that you can replace Context.IndentedJSON with a call to
	// Context.JSON to send more compact JSON. In practice, the indented
	// form is much easier to work with when debugging and the size
	// difference is usually small.
}

// postAlbums adds and album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum
	// Use Context.BindJSON to bind the request body to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	// Append the album struct initialized from the JSON to the albums
	// slice.
	albums = append(albums, newAlbum)
	// Add a 201 status code to the response, along with JSON representing
	// the album you added.
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID values matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	// Use Context.Param to retrieve the id path parameter from the URL.
	// When you map this handler to a path, you’ll include a placeholder
	// for the parameter in the path.
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	// Loop over the album structs in the slice, looking for one whose ID
	// field value matches the id parameter value. If it’s found, you
	// serialize that album struct to JSON and return it as a response with
	// a 200 OK HTTP code.
	for _, a := range albums {
		// As mentioned above, a real-world service would likely use a
		// database query to perform this lookup.
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	// Return an HTTP 404 error with http.StatusNotFound if the album isn’t
	// found.
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
