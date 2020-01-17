package main

import "github.com/gin-gonic/gin"

// structs are objects/classes

func main() {
	// capital D means it's exported (public method)
	// := is the go funky syntax for creating a new variable
	// only inside a function
	// go likes single letter variables
	router := gin.Default()
	router.GET("/puppy", handlePuppy)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
