package main


import "fmt"


type Vertex struct {
    Lat, Long float64
}


var m map[string]Vertex

// a map maps keys to values
func main() {
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
        40.68433, -74.39967,
    }
    m["Hello"] = Vertex{
        0, 1,
    }
    fmt.Println(m["Bell Labs"])
    fmt.Println(m["Hello"])

}
