package main

import (
  "golang.org/x/tour/pic"
  "image"
  "image/color"
  "image/png"
  "os"
)

type Image struct{
    width int
    height int
}

func (img Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
    img_func := func(x, y int) uint8 {
        //return uint8(x*y)
        //return uint8((x+y) / 2)
        return uint8(x^y)
    }
    v := img_func(x, y)
    return color.RGBA{v, v, 255, 255}
}

func main() {
  m := Image{256, 64}
  pic.ShowImage(m)
  toimg, _ := os.Create("image.png")
  defer toimg.Close()

  png.Encode(toimg, m)
}
