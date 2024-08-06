package main


import (
    "fmt"
)



type ErrNegativeSqrt float64


// error message
func (e *ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(*e))
}


// absolute function
func Abs(x float64) float64 {
    if x < 0 {
        return float64(-x)
    }
    return float64(x)
}


func Sqrt(x float64) (float64, error) {
    if x < 0 {
        e := ErrNegativeSqrt(x)
        return 0, &e
    }
    z := float64(x / 2)
    diff := Abs((z*z - x) / (2*x))
    for diff > 1e-15 {
        diff = (z*z - x) / (2*x)
        z -= diff
        diff = Abs(diff)
    }
    return z, nil
}



func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))


    y, ok := Sqrt(3)
    if ok != nil {
        fmt.Println(ok)
    } else {
        fmt.Println(y)
    }


    y, ok = Sqrt(-3)
    if ok != nil {
        fmt.Println(ok)
    } else{
        fmt.Println(y)
    }

}
