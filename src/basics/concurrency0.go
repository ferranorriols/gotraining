package main

import "fmt"

func main() {
    foo()
    bar()
}
func foo() {
    fmt.Println("hello")
}
func bar() {
    fmt.Println("bye")
}

//hello
//bye