package main


import "fmt"


func main() {
    fmt.Println("counting")


    // deferred function calls are pushed onto a stack
    // when a function returns, its deferred calls are executed in last-in-fast-out order
    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }

    fmt.Println("done")
}
