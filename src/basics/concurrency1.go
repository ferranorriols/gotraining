package main

import (
    "fmt"
    "sync"
)

var waitgroup sync.WaitGroup

func main() {
    waitgroup.Add(2)
    go foo()
    go bar()
    waitgroup.Wait()
}
func foo() {
    fmt.Println("hello")
    waitgroup.Done()
}
func bar() {
    fmt.Println("bye")
    waitgroup.Done()
}

//bye
//hello