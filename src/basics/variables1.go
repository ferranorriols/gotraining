package main

import "fmt"

func main() {

    var defaultint int;
    varint := 10
    varstr := "hello world"
    varflo := 3.14
    varboo := true
    foo := int32(varint)

    fmt.Printf("%T %v \n", defaultint, defaultint)
    fmt.Printf("%T %v \n", varint, varint)
    fmt.Printf("%T %v \n", varstr, varstr)
    fmt.Printf("%T %v \n", varflo, varflo)
    fmt.Printf("%T %v \n", varboo, varboo)
    fmt.Printf("%T %v \n", foo, foo)
}

//output:
// int 0
// int 10
// string hello world
// float64 3.14
// bool true
// int32 10