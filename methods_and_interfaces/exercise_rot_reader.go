package main


import (
    "io"
    "os"
    "strings"
)


type rot13Reader struct {
    r io.Reader
}


func (rot rot13Reader) Read(b []byte) (int, error) {
    n, err := rot.r.Read(b)

	for i := 0; i < n; i++ {
      if b[i] < byte('a') + 13 {
          b[i] += 13
      } else if b[i] < byte('a') + 13 * 2 {
          b[i] -= 13
      } else if b[i] < byte('a') + 13 * 3 {
          b[i] += 13
      } else {
          b[i] -= 13
      }
	}


    return n, err
}


func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
