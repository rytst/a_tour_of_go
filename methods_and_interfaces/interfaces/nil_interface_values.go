package main


import "fmt"


type I interface {
    M()
}


// a nil interface value holds neither value nor concrete type
// calling a method on a nil interface is a run-time error
func main() {
    var i I
    describe(i)
    i.M()
}



func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}
