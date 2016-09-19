package main

import (
  "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  return fmt.Sprintf("Sqrt: negative number %g", e)
}

const delta = 1e-15

func Sqrt(x float64) (float64, error) {
  if x < 0 {
    return 0, ErrNegativeSqrt(x)
  }
    var z, d float64 = x, 1
    for d > delta {
        z0 := z
        z = z-(z*z-x)/(2*z)
        d = z-z0
        if d < 0 {
            d = -d
        }
    }
    return z, nil
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}
