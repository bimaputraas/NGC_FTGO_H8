package main

import (
	"fmt"
	"sync"
)

func produce(c chan int,wg *sync.WaitGroup) {
	defer wg.Done()
	for i:=1;i<11;i++{
		c <- i
	}
	close(c)
}

func consume(c chan int,wg *sync.WaitGroup) {
	defer wg.Done()
	for output := range c {
		fmt.Println(output)
	}
}

func main() {
	wg := sync.WaitGroup{}
	c := make(chan int, 5)
	
	wg.Add(2)
	go consume(c,&wg)
	go produce(c,&wg)
	
	wg.Wait()
}

// perbedaan dari task 3 : channel pada task 4 memiliki kapasitas penyimpanan data hingga 5 (integer) sebelum data tersebut di consume, sedangkan channel pada task 3 akan melakukan blocking setiap menerima 1 data namun belum di consume.