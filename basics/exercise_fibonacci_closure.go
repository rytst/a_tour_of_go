package main


import "fmt"


// fibonacci is a function that returns
// a function that returns an int
func fibonacci() func() int {
    f_pre := 0
    f := 1
    return func() int {
        res := f_pre
        f_next := f_pre + f
        f_pre = f
        f = f_next
        return res
    }
}


func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
