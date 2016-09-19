package main

import (
  "io"
  "os"
  "strings"
)

type rot13Reader struct {
  r io.Reader
}

func rot13(b byte) byte {
    switch {
        case 'A' <= b && b <= 'M', 'a' <= b && b <= 'm':
            b += 13
        case 'N' <= b && b <= 'Z', 'n' <= b && b <= 'z':
            b -= 13
    }
    return b
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
    n, err = r.r.Read(p)
    for i, b := range p {
        p[i] = rot13(b)
    }
    return
}

func main() {
  s := strings.NewReader("Lbh penpxrq gur pbqr!")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}
