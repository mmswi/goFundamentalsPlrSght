package main

import "fmt"

func main() {
	type courseMeta struct {
		Author string
		Level string
		Rating float64
	}

	// both declarations will create an instance of our courseMeta type
	// these declarations will initialize the properties with 0 or empty strings
	//var DockerDeepDive courseMeta
	//DockerDeepDive := new(courseMeta)

	//initializing with values
	DockerDeepDive := courseMeta{
		Author: "Mihai M",
		Level: "Advanced",
		Rating: 10,
	}

	fmt.Print(DockerDeepDive)

	fmt.Println("\n Course Author is: ", DockerDeepDive.Author)
	DockerDeepDive.Rating = 11
	fmt.Println("\n Course Rating is: ", DockerDeepDive.Rating)
}
