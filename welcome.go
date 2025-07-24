package main
import (
    "fmt"
    "strings"
)
func main() {
  message := "Welcome,{username}!"
  message = strings.ReplaceAll(message, "{username}","Bob")
  fmt.Println(message)
}
