package main

import "fmt"
import "rsc.io/quote" // use external package

func HelloWorld() {
    fmt.Println("Hello, World")
}

func HelloWorldWithQuote() {
    fmt.Println(quote.Go())
}

func main() {
    HelloWorld() // Hello, World
    HelloWorldWithQuote() // Don't communicate by sharing memory, share memory by communicating.
}
// go mod init hello // enable dependency tracking for your code by creating a go.mod file
// go run helloworld.go 
// go build helloworld.go & ./helloworld
