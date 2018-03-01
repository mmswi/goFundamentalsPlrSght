package main

import (
	"fmt"
	"runtime"
)

var (
	name = "mihai"
	module float64
) // var sunt disponibile in intreg pachetul

func main() {
	fmt.Println("Hello from", runtime.GOOS)
	fmt.Println("name is ", name)
	fmt.Println("module is ", module)

	a := 10.0 // variabilele cu := sunt disponibile doar in functie si pot fi declarate asa doar in functie
	b := 3

	c:= int(a) + b

	ptr := &c // adresa de memorie

	fmt.Println(c)
	fmt.Println("memory address of *c var is ", ptr, "and the value of *c is", *ptr)

	course := "golang course"
	fmt.Println("My name is", name, "you're watching", course)

	changeCourse(&course) // aruncam adresa de memorie

	fmt.Println("you're now watching", course)

}

func changeCourse(courseNewName *string) string { //ii spunem ca valoarea de la adresa este de tip string
	*courseNewName = "This is the new course" // asignam pe valoare courseNewName

	fmt.Println("Changing the course to", *courseNewName) // printam valoarea

	return *courseNewName // returnam valoarea
}