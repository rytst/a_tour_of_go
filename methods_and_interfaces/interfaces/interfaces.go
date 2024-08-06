package main


import (
    "fmt"
    "math"
)


type Abser interface {
    Abs() float64
}

func main() {
    var a Abser
    var b Abser

    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}

    a = f // a MyFloat implements Abser
    b = &v // b &Vertex implements Abser

    // in the following line, v is a Vertex (not *Vertex)
    // and does NOT implement Abser
    // a = v

    fmt.Println("a:", a.Abs())
    fmt.Println("b:", b.Abs())
}


type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}




type Vertex struct {
    X, Y float64
}


func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
