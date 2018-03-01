package main

import "fmt"

func main() {
	coursesInProg := []string{"curs 1", "curs 2", "curs 3"}
	coursesCompleted := []string{"curs 1", "curs 5"}

	for _,i := range coursesInProg{
		fmt.Println(i)
		for _, j := range coursesCompleted {
			if i == j {
				fmt.Println(i, "is in both lists")
			}
		}
	}
}
