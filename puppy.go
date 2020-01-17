package main

import "github.com/gin-gonic/gin"

type Puppy struct {
	Name  string
	Breed string
	Age   int
	Male  bool
}

func handlePuppy(context *gin.Context) {
	puppy := getPuppy("Bixby")
	context.JSON(200, puppy)
}

func getPuppy(name string) Puppy {
	return Puppy{
		Name:  name,
		Breed: "retriever",
	}
}
