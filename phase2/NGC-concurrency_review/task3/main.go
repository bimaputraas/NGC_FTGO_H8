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
	c := make(chan int)
	
	wg.Add(2)
	go consume(c,&wg)
	go produce(c,&wg)
	
	wg.Wait()
}