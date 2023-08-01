package main

import (
	"fmt"
)

func fizzbuzz(angka int, c chan string) {
	// n = 99
	result := ""

	if angka != 0 && angka%3 == 0 && angka%5 == 0 {
		result = fmt.Sprintf("%dFizzBuzz", angka)

	} else if angka%3 == 0 && angka != 0 {
		result = fmt.Sprintf("%dFizz", angka)

	} else if angka%5 == 0 && angka != 0 {
		result = fmt.Sprintf("%dBuzz", angka)

	} else {
		result = fmt.Sprintf("%d", angka)
	}
	c <- result

}

func countOddEven(angka int, c chan string) {
	result := ""
	if angka%2 == 0 {
		result = "even"
	} else {
		result = "odd"
	}
	c <- result

}

func main() {
	// release 1
	c := make(chan string)

	release2_sum := 0

	result := ""

	countodd := 0
	counteven := 0
	temp_count := ""
	for i := 0; i < 100; i++ {
		go fizzbuzz(i, c)
		result = <-c
		fmt.Println(result)

		// release 2
		release2_sum += i

		// release 3
		go countOddEven(i, c)
		temp_count = <-c
		if temp_count == "odd" {
			countodd += 1
		} else if temp_count == "even" {
			counteven += 1
		}
	}

	// release 2 end
	fmt.Printf("\nTotal penjumlahan dari 0-99 = %d", release2_sum)

	// release 3
	fmt.Printf("\nBilangan ganjil : %d\nBilangan genap: %d\n", countodd, counteven)

	close(c)

}
