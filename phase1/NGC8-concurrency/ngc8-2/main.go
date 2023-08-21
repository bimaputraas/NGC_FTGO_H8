package main

import (
	"fmt"
	"math"
)

func main() {
	c := make(chan int)

	num := 100
	go SumSquare(num, c)
	result1 := <-c

	go SquareSum(num, c)
	result2 := <-c

	fmt.Printf("Release 1 (SumSquare) : %d\nRelease 2 (SquareSum) : %d\n,", result1, result2)

	close(c)

}

func SumSquare(num int, c chan int) {
	length := num
	sum := 0
	for i := 1; i < length+1; i++ {
		sum += i
	}

	sumsquare := int(math.Pow(float64(sum), 2))
	c <- sumsquare
}

func SquareSum(num int, c chan int) {
	length := num
	squaresum := 0
	for i := 1; i < length+1; i++ {
		squaresum += int(math.Pow(float64(i), 2))
	}

	c <- squaresum
}
