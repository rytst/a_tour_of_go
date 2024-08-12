package main


import "fmt"


// list represents a singly-linked list that holds
// values of any type
type List[T any] struct {
    next *List[T]
    val T
}

func main() {
    var v0 *List[int] = nil
    for i := 0; i < 10; i++ {
        v0 = &List[int]{v0, i}
    }

    for v0.next != nil {
        fmt.Println(v0.val)
        v0 = v0.next
    }
    fmt.Println(v0.val)
}
