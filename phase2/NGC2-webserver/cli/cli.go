package cli

import (
	"fmt"
	"time"
)

func StartServerCLI(port string) {
	fmt.Println("")
	for i:=3;i>=1;i--{
		fmt.Printf("Starting server in %d...\n",i)
		time.Sleep(time.Millisecond *750)
	}
	fmt.Println("")
	go func (){
		for {
			fmt.Printf("Reply from localhost port %s: bytes=32 time=0ms TTL=56\n",port)
			time.Sleep(time.Millisecond*1850)
		}
	}()
}
