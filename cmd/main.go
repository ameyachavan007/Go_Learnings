package main
import (
  "fmt"
  "os"
  "hello"
)

func main () {
  fmt.Printf(hello.Say(os.Args[1:]))
}