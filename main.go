package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, World!")

	for i := 1; i <= 3; i++ {
		fmt.Println("Hello for loop!")
	}

	curTime := time.Now().Local()
	fmt.Println("The current time is ", curTime.Format("02-01-2006"))
}
