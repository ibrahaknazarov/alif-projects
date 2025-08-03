package main

import (
	"fmt"
)
func main() {
	var commitions float64
	fmt.Println("Enter a sum of you want to transfer")
	fmt.Scanln(&commitions)



	var mounthTransfer = 5000

	if commitions < float64(mounthTransfer){
		fmt.Println("Transfer is complate without persent")
	} else if commitions > float64(mounthTransfer) {
		fmt.Printf("Transfer is complited with persent")
	}
}

