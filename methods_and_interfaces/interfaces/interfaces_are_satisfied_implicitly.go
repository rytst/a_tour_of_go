package main


import "fmt"


type I interface {
    M()
}


type T struct {
    S string
}


// this method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
    fmt.Println(t.S)
}



func main() {
    var i T = T{"Hello"}
    i.M()
}
