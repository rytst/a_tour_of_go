package main


import "golang.org/x/tour/pic"



func Pic(dx, dy int) [][]uint8 {
    array := make([][]uint8, dy)
    for y := 0; y < dy; y++ {
        for x := 0; x < dx; x++ {
            z := uint8(x^y)
            array[y] = append(array[y], z)
        }
    }

    return array

}

func main() {
    pic.Show(Pic)
}
