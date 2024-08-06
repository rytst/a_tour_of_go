package main


import "fmt"



// the interface type that specifies zero methods is known as the empty interface
// an empty interface may hold values of any type
// empty interfaces are used by code that handles values of unknown type
func main() {
    var i interface{}
    describe(i)

    i = 42
    describe(i)

    i = "hello"
    describe(i)
}



func describe(i interface{}) {
    fmt.Printf("(%v, %T)\n", i, i)
}
