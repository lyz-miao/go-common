package image

import (
    "fmt"
    "image"
    "image/color"
    "io"
)

func AverageImageColorWithReader(in io.Reader) (string, error) {
    i, _, err := image.Decode(in)
    if err != nil {
        return "", err
    }
    return averageImageColor(i), nil
}

func AverageImageColorWithImage(in image.Image) string {
    return averageImageColor(in)
}

func averageImageColor(in image.Image) string {
    var r, g, b, a uint64

    bounds := in.Bounds()

    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            pr, pg, pb, pa := in.At(x, y).RGBA()

            r += uint64(pr)
            g += uint64(pg)
            b += uint64(pb)
            a += uint64(pa)
        }
    }

    count := uint64(bounds.Dy() * bounds.Dx())
    mm := color.NRGBA{
        R: uint8(r / count / 0x101),
        G: uint8(b / count / 0x101),
        B: uint8(g / count / 0x101),
        A: 255,
    }
    rgba := color.RGBAModel.Convert(mm).(color.RGBA)

    return fmt.Sprintf("%.2x%.2x%.2x", rgba.R, rgba.B, rgba.G)
}
