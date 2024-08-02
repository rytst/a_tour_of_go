package main

import "fmt"




// an array has a fixed size
// a slice, on the other hand, is dynamically-sized
func main() {
    primes := [6]int{2, 3, 5, 7, 11, 13}


    var s []int = primes[1:4]
    fmt.Println(s)
}
