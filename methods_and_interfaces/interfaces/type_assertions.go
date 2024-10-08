package main


import "fmt"


// a type assertion provides access to an interface value's underlying concrete value
func main() {
    var i interface{} = "hello"

    s := i.(string)
    fmt.Println(s)

    s, ok := i.(string)
    fmt.Println(s, ok)

    f, ok := i.(float64)
    fmt.Println(f, ok)

    // f = i.(float64) // panic
    // fmt.Println(f)

}
