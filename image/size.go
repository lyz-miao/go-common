package image

import (
    "image"
    "io"
)

type Size struct {
    X, Y int
}

func GetSizeWithReader(in io.Reader) (*Size, error) {
    i, _, err := image.Decode(in)
    if err != nil {
        return nil, err
    }
    return getSize(i), nil
}

func GetSizeWithImage(in image.Image) *Size {
    return getSize(in)
}

func getSize(in image.Image) *Size {
    x := in.Bounds().Size().X
    y := in.Bounds().Size().Y

    return &Size{x, y}
}
