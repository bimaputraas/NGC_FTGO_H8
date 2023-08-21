package main

import (
	"belajar/worker"
	"fmt"
	"sync"
	"time"
)

func tes() {
	wi,err := worker.NewWorkerInstance(5)
	if err != nil{
		panic(err)
	}

	for i:=0;i<12;i++{
		wi.Do(func() {
			time.Sleep(3*time.Second)
			fmt.Println(time.Now())
		})
	}

	wi.Wait()
}

func main(){
	go func(){
		fmt.Println("Hello World")
	}()
	channelWorker := make(chan func())
	wg := sync.WaitGroup{}
	for i:=0;i<50;i++{
		go func(){
			for {
				j := <- channelWorker
				j()
				wg.Done()
			}
		}()
	}

	var angka int
	for i:=0;i<10;i++{
		wg.Add(1)
		x := i
		channelWorker <- func() {
			time.Sleep(time.Second)
			fmt.Println(x+1)
		}
		angka = 10
	
	}
	fmt.Println("Sebelum",angka)
	// wg.Add(1)
	// channelWorker <- func() {
		// 	time.Sleep(time.Second*3)
		// 	fmt.Println("tes2")
		// }
		
		wg.Wait()
		fmt.Println("Sesudah",angka)
		fmt.Println(num,"num")
	}

var num int
func init(){
	num = 10
}