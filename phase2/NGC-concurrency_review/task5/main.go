package main

import (
	"fmt"
)

func oddEven(chanOdd, chanEven chan int) {
	// defer wg.Done()
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			chanEven <- i
		}else{

			chanOdd <- i
		}
	}
	close(chanEven)
	close(chanOdd)
}

func main() {

	channelOdd := make(chan int)
	channelEven := make(chan int)

	
	go oddEven(channelOdd,channelEven)

	for{
		select{
		case resultOdd,ok := <- channelOdd:
			if !ok{
				return
			}
			fmt.Println("received an odd number",resultOdd)
		case resultEven,ok := <- channelEven:
			if !ok{
				return
			}
			fmt.Println("received an even number",resultEven)
		}
	}

}

// func main() {
// 	wg := sync.WaitGroup{}
// 	channelOdd := make(chan int)
// 	channelEven := make(chan int)
	
// 	wg.Add(1)
// 	go oddEven(channelOdd,channelEven,&wg)
	
// 	go func(){
// 		for resultEven := range channelEven{
// 			fmt.Println("received an even number",resultEven)
			
// 		}
// 	}()
	
// 	for resultOdd := range channelOdd{
// 		fmt.Println("received an odd number",resultOdd)
		
// 	}
// 	wg.Wait()
// }