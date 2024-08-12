package main


import (
    "fmt"
)


// a sender can close a channel to indicate that no more
// values will be sent
//
// recievers can test whether a channel has
// been closed by assigning a second parameter to the
// recieve expression
func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}


func main() {
    c := make(chan int, 10)

    go fibonacci(cap(c), c)


    // the loop `for i := range c` recieves values
    // from the channel repeatedly until it is closed
    for i := range c {
        fmt.Println(i)
    }
}

// closing is only necessary when the reciever must be told
// there are no more values coming, such as to terminate a range loop
