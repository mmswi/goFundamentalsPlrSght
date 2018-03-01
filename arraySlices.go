package main

import "fmt"

func main() {
	//making a slice -> make(<type>,<len>,<capacity>)
	//capacity is the maximum size of the slice
	//you can also create it by giving it a type and a length. the capacity will = the length
	myCourses := make([]string, 5, 10)
	// A slice creates an underlying array - a slice is a reference to that array

	myOtherCourses := []string{"Docker", "Puppet", "Python"}
	//the capacity and length will be = with the number of elements declared in the initialization (3)

	fmt.Printf("For myCourses Length is: %d \nCapacity is: %d", len(myCourses), cap(myCourses))
	fmt.Printf("\nFor myOtherCourses Length is: %d \nCapacity is: %d", len(myOtherCourses), cap(myOtherCourses))

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mySlice[1] = 0
	//slicing a slice
	sliceOfSlice := mySlice[2:5]
	//Note is you put [:5] - will begin from 0
	//[2:] - will go until the last index
	fmt.Println(sliceOfSlice)

	mySliceTwo := make([]int, 1, 4)
	fmt.Printf("Length is: %d \nCapacity is: %d", len(mySliceTwo), cap(mySliceTwo))

	for i:= 1; i<17; i++ {
		mySliceTwo = append(mySliceTwo, i)
		//once we hit the capacity, append will double the size of the underlying array
		// - it will create a new array with twice the length
		fmt.Printf("\nCapacity is: %d", cap(mySliceTwo))
	}
	// "for range" loops iterate slices
	// "for range" returns two values - index and data

	for _, i := range mySlice {
		fmt.Println("for range loop: ", i)
	}

	//can "append()" slices to slices with the elipses "..."

	newSlice := []int{10, 30, 40}
	mySlice = append(mySlice, newSlice...)
	fmt.Println(mySlice)
}
