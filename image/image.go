package image

import (
    _ "golang.org/x/image/bmp"
    _ "golang.org/x/image/tiff"
    _ "golang.org/x/image/webp"
    "image"
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"
    "io"
)

//var ErrorNotSupported = errors.New("mime is not supported")
//
//type Image struct {
//    data interface{}
//    //i    image.Image
//    //g    *gif.GIF
//    mime string
//}

func Decode(in io.Reader) (image.Image, string, error) {
   return image.Decode(in)
}

//func New(img io.Reader, mime string) (*Image, error) {
//    switch mime {
//    case "image/jpeg", "image/jpg", "image/png", "image/bmp", "image/tiff", "image/webp":
//        i, _, err := image.Decode(img)
//        if err != nil {
//            return nil, err
//        }
//        return &Image{i, mime}, nil
//    case "image/gif":
//        i, err := gif.DecodeAll(img)
//        if err != nil {
//            return nil, err
//        }
//        return &Image{i, mime}, nil
//    }
//
//    return nil, ErrorNotSupported
//}
//
//func (i *Image) GetSize() Size {
//    x := 0
//    y := 0
//
//    switch i.data.(type) {
//    case image.Image:
//        x = i.data.(image.Image).Bounds().Size().X
//        y = i.data.(image.Image).Bounds().Size().Y
//        break
//    case gif.GIF:
//        x = i.data.(gif.GIF).Config.Width
//        y = i.data.(gif.GIF).Config.Height
//        break
//    }
//
//    return Size{x, y}
//}
//
//func (i *Image) Resize(size int) {
//    switch i.data.(type) {
//    case image.Image:
//        result := resizeImage(i.data.(image.Image), size)
//        i.data = result
//        break
//    case gif.GIF:
//        resizeGif(i.data.(*gif.GIF), size)
//        break
//    }
//}
//
//func resizeImage(in image.Image, size int) image.Image {
//    x := in.Bounds().Size().X
//    y := in.Bounds().Size().Y
//
//    var f []gg.Filter
//
//    if x > y {
//        if x > size {
//            f = append(f, gg.Resize(size, 0, gg.LanczosResampling))
//        }
//    } else if x < y {
//        if y > size {
//            f = append(f, gg.Resize(0, size, gg.LanczosResampling))
//        }
//    } else if x == y {
//        f = append(f, gg.Resize(size, size, gg.LanczosResampling))
//    }
//
//    g := gg.New(f...)
//    dst := image.NewNRGBA(g.Bounds(in.Bounds()))
//    g.Draw(dst, in)
//
//    return dst
//}
//
//func resizeGif(in *gif.GIF, size int) {
//    tmp := image.NewNRGBA(in.Image[0].Bounds())
//
//    for index := range in.Image {
//        x := in.Image[index].Bounds().Size().X
//        y := in.Image[index].Bounds().Size().Y
//
//        var f []gg.Filter
//
//        if x > y {
//            if x > size {
//                f = append(f, gg.Resize(size, 0, gg.LanczosResampling))
//            }
//        } else if x < y {
//            if y > size {
//                f = append(f, gg.Resize(0, size, gg.LanczosResampling))
//            }
//        } else if x == y {
//            f = append(f, gg.Resize(size, size, gg.LanczosResampling))
//        }
//
//        filter := gg.New(f...)
//
//        gg.New().DrawAt(tmp, in.Image[index], in.Image[index].Bounds().Min, gg.OverOperator)
//        dst := image.NewPaletted(filter.Bounds(tmp.Bounds()), in.Image[index].Palette)
//        filter.Draw(dst, tmp)
//        in.Image[index] = dst
//        in.Config.Width = dst.Bounds().Size().X
//        in.Config.Height = dst.Bounds().Size().Y
//    }
//}
//
//func (i *Image) Output() ([]byte, error) {
//    switch i.data.(type) {
//    case image.Image:
//        switch i.mime {
//        case "image/jpeg", "image/jpg":
//            w := bytes.Buffer{}
//            err := jpeg.Encode(&w, i.data.(image.Image), &jpeg.Options{Quality: 100})
//            if err != nil {
//                return nil, err
//            }
//
//            return w.Bytes(), nil
//        case "image/png":
//            w := bytes.Buffer{}
//            err := png.Encode(&w, i.data.(image.Image))
//            if err != nil {
//                return nil, err
//            }
//
//            return w.Bytes(), nil
//        }
//    case gif.GIF:
//        w := bytes.Buffer{}
//        err := gif.EncodeAll(&w, i.data.(*gif.GIF))
//        if err != nil {
//            return nil, err
//        }
//
//        return w.Bytes(), nil
//    }
//
//    return nil, nil
//}
