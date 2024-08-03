package main


import "golang.org/x/tour/pic"



func Pic(dx, dy int) [][]uint8 {
    slice := make([][]uint8, dy)
    for i := range slice {
        slice[i] = make([]uint8, dx)
    }

    for y := 0; y < dy; y++ {
        for x := 0; x < dx; x++ {
            z := (x + y) / 2
            slice[y][x] = uint8(z)
        }
    }

    return slice

}

func main() {
    pic.Show(Pic)
}
