package main

import (
    "fmt"
)


func Abs(x float64) float64 {
    if x < 0 {
        return -x
    } else {
        return x
    }
}


func Sqrt(x float64) float64 {
    // z := float64(1)
    z := x / 2

    // for i := 0; i < 15; i++ {
    //     z -= (z*z - x) / (2 * z)
    // }
    
    diff := Abs((z*z - x) / (2 * z))
    for diff > 1e-15 {
        z -= (z*z - x) / (2 * z)
        diff = Abs((z*z - x) / (2 * z))
    }

    return z
}



func main() {
    fmt.Println(Sqrt(2))
}
