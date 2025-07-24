package main
import (
	"fmt"
)
func main()  {
	 phoneNumber := 4
	switch phoneNumber {
	case 1:
		fmt.Println("+992 48 888 1111")
	case 2:
		fmt.Println("+992 (48) 88-81-111")
	case 3:
		fmt.Println("+992-48-888-11-11")
	case 4:
		fmt.Println("+992 (48) 8881111")
	 default :
	 	fmt.Println("+992 000-000-0000","This number is invalid")
	}
}
