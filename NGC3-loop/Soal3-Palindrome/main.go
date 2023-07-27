// main.go
package main

import "fmt"

func main() {
	// soal 3
	kata := "katak"
	var kata_palindrome string

	for i := (len(kata) - 1); i >= 0; i-- {
		kata_palindrome += string(kata[i])
	}

	if kata_palindrome == kata {
		fmt.Println("output 3 :", true)
	} else {
		fmt.Println("output 3 :", false)
	}

	// soal 4
	kata2 := "xoxxoo"

	count_x := 0
	count_o := 0

	for i := 0; i < len(kata2); i++ {
		if string(kata2[i]) == "x" {
			count_x++
		} else if string(kata2[i]) == "o" {
			count_o++
		}
	}

	if count_x == count_o {
		fmt.Println("output 4 :", true)
	} else {
		fmt.Println("output 4 :", false)
	}

	// soal 5

	panjang := 5

	fmt.Println("Output 5 : ")
	for i := 0; i < panjang; i++ {
		fmt.Println("*")
	}

	// soal 6

	panjang2 := 5

	fmt.Println("Output 6 : ")
	for i := 0; i < panjang2; i++ {
		for i := 0; i < panjang2; i++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	// soal 7

	panjang3 := 10

	fmt.Println("Output 7 : ")
	for i := 1; i <= panjang3; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	// soal 8
	panjang4 := 10

	fmt.Println("Output 8 : ")
	for i := panjang4; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	// soal 9 - sorting

	array := []int{6, 8, 1, 2, 80, 12, 42, 11, 26}
	for i := 0; i < len(array)-1; i++ {
		fmt.Println(i)
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				temp_num := array[i]
				array[i] = array[j]
				array[j] = temp_num
				fmt.Println(array)
			}
		}
	}
	fmt.Println("output sort : ", array)
}
