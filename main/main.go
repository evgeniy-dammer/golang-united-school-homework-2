package main

import (
	"fmt"

	solution "github.com/evgeniy-dammer/golang-united-school-homework-2"
)

func main() {
	fmt.Println(solution.CalcSquare(10.0, solution.SidesTriangle))
	fmt.Println(solution.CalcSquare(10.0, solution.SidesSquare))
	fmt.Println(solution.CalcSquare(10.0, solution.SidesCircle))
}
