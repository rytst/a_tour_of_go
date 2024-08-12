package main


import "fmt"


// this `Index` function works for any type that supports comparison

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
    // this declaration means that `s` is a slice of any type `T`
    // that fulfills the built-in constraint `comparable`
    for i, v := range s {
        // v and x are type T, which has the comparable
        // constraint, so we can use `==` here
        if v == x {
            return i
        }
    }
    return -1
}


func main() {
    // Index works on a slice of ints
    si := []int{10, 20, 15, -10}
    fmt.Println(Index(si, 15))

    // Index also works on a slice of strings
    ss := []string{"foo", "bar", "baz"}
    fmt.Println(Index(ss, "hello"))
}
