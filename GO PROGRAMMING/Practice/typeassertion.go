package main

import "fmt"

// define a simple interface
type myInterface interface {
    getValue() string
}

// define a struct that implements the interface
type myStruct struct {
    value string
}

func (s myStruct) getValue() string {
    return s.value
}

func main() {
    // create an object of type myStruct and assign it to the interface variable
    var iface myInterface = myStruct{value: "Hello, world!"}

    // use type assertion to extract the value from the interface
    if s, ok := iface.(myStruct); ok {
        fmt.Println("Value:", s.getValue())
    } else {
        fmt.Println("Error: unexpected type")
    }
}
