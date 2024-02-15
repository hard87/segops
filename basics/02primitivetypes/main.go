package main

import "fmt"

func main() {
	sumResult := sumTwoInts(2)
	var firstName, lastName string = "", ""
	myName := myNameIs(firstName, lastName)
	v1, v2 := 100.0, 3.0
	divisionResult := divisionTwoNumber(v1, v2)

	phrass := fmt.Sprintf("Alguns fatos: \n2 mais 2 é igual %d \nMeu nome é %s \n%d dividido por %d é igual a %d", sumResult, myName, v1, v2, divisionResult)

	fmt.Printf("%s\n", phrass)
}

func sumTwoInts(num int) int {
	return num + num
}

func myNameIs(firstName, lastName string) string {
	return firstName + " " + lastName
}

func divisionTwoNumber(v1, v2 float64) float64 {
	return v1 / v2
}
