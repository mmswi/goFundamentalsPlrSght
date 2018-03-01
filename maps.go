package main

import "fmt"

func main() {
	//map[keytype]valueType
	leagueTitles := make(map[string]int)
	leagueTitles["Sunderland"] = 6
	leagueTitles["Newcastle"] = 4

	//another way of initializing a map
	recentHead2Head := map[string]int {
		"Sunderland": 5,
		"Newcastle": 0,
	}

	fmt.Printf("\nLeague titles: %v\nRecent head to heads: %v\n", leagueTitles, recentHead2Head)

	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}

	testMap["F"] = 6

	for key, value := range testMap{
		//if you run the file multiple times, you will see different starting keys
		fmt.Printf("\nKey is: %v Value is: %v", key, value)
	}

	delete(testMap, "F")
}
