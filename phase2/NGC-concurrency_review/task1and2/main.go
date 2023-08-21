package main

import (
	"fmt"
	"sync"
)

// task 1 - basic go routine
// 1.1. Write a function printNumbers that prints numbers from 1 to 10.
// 1.2. Write a function printLetters that prints letters from 'a' to 'j'.
// 1.3. Use goroutines to concurrently run both functions.

func printNumbers(wg *sync.WaitGroup){
	defer wg.Done()
	for i:=1;i<11;i++{
			fmt.Println(i)
	}
}

func printLetters(wg *sync.WaitGroup){ 
	// rune a to j
	defer wg.Done()
	a := 'a'
	j := 'j'
	for i:=a;i<=j;i++{
		fmt.Println(string(i))
	}
}

func main() {
	wg := sync.WaitGroup{}
	
	wg.Add(2)
	go printNumbers(&wg)
	go printLetters(&wg)

	wg.Wait()
}