package main


import "fmt"


func main() {

    // provide the buffer length as the second argument
    // to `make` to initialize a buffered channel
    ch := make(chan int, 2)

    // sends to a buffered channel block only when the buffer is full
    //
    // recieves block when the buffer is empty
    ch <- 1
    ch <- 2
    // ch <- 3 // dead lock

    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
