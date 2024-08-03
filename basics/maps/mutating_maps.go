package main


import "fmt"


func main() {
    m := make(map[string]int)

    m["Answer"] = 42
    fmt.Println("The value:", m["Answer"])

    m["Answer"] = 48
    fmt.Println("The value:", m["Answer"])

    delete(m, "Answer")
    fmt.Println("The value:", m["Answer"])


    // if key is in m, ok is true
    // if not, ok is false
    // if key is not in the map,
    // then elem is the zero value for map's element type
    v, ok := m["Answer"]
    fmt.Println("The value:", v, "Present?", ok)
}
