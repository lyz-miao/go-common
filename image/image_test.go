package image

import (
    "os"
    "testing"
)

func TestGETJPEGSize(t *testing.T) {
    f, err := os.Open("./t.jpg")
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    result, err := GetSizeWithReader(f)
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    t.Logf("x: %v y: %v\n", result.X, result.Y)

    t.Cleanup(func() {
        _ = f.Close()
    })
}

func TestGETPNGSize(t *testing.T) {
    f, err := os.Open("./t.png")
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    result, err := GetSizeWithReader(f)
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    t.Logf("x: %v y: %v\n", result.X, result.Y)

    t.Cleanup(func() {
        _ = f.Close()
    })
}

func TestGETBMPSize(t *testing.T) {
    f, err := os.Open("./t.bmp")
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    result, err := GetSizeWithReader(f)
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    t.Logf("x: %v y: %v\n", result.X, result.Y)

    t.Cleanup(func() {
        _ = f.Close()
    })
}

func TestGETGIFSize(t *testing.T) {
    f, err := os.Open("./t.gif")
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    result, err := GetSizeWithReader(f)
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    t.Logf("x: %v y: %v\n", result.X, result.Y)

    t.Cleanup(func() {
        _ = f.Close()
    })
}

func TestGETWEBPSize(t *testing.T) {
    f, err := os.Open("./t.webp")
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    result, err := GetSizeWithReader(f)
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    t.Logf("x: %v y: %v\n", result.X, result.Y)

    t.Cleanup(func() {
        _ = f.Close()
    })
}

func TestGETTIFFSize(t *testing.T) {
    f, err := os.Open("./t.tiff")
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    result, err := GetSizeWithReader(f)
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    t.Logf("x: %v y: %v\n", result.X, result.Y)

    t.Cleanup(func() {
        _ = f.Close()
    })
}


func TestGetJPEGAverageColor(t *testing.T) {
    f, err := os.Open("./t.jpg")
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    result, err := AverageImageColorWithReader(f)
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    t.Logf("result: #%v\n", result)

    t.Cleanup(func() {
        _ = f.Close()
    })
}

func TestGetGIFAverageColor(t *testing.T) {
    f, err := os.Open("./t.gif")
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    result, err := AverageImageColorWithReader(f)
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    t.Logf("result: #%v\n", result)

    t.Cleanup(func() {
        _ = f.Close()
    })
}

func TestGetPNGAverageColor(t *testing.T) {
    f, err := os.Open("./t.png")
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    result, err := AverageImageColorWithReader(f)
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    t.Logf("result: #%v\n", result)

    t.Cleanup(func() {
        _ = f.Close()
    })
}

func TestGetBMPAverageColor(t *testing.T) {
    f, err := os.Open("./t.bmp")
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    result, err := AverageImageColorWithReader(f)
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    t.Logf("result: #%v\n", result)

    t.Cleanup(func() {
        _ = f.Close()
    })
}

func TestGetWEBPAverageColor(t *testing.T) {
    f, err := os.Open("./t.webp")
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    result, err := AverageImageColorWithReader(f)
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    t.Logf("result: #%v\n", result)

    t.Cleanup(func() {
        _ = f.Close()
    })
}

func TestGetTIFFAverageColor(t *testing.T) {
    f, err := os.Open("./t.tiff")
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    result, err := AverageImageColorWithReader(f)
    if err != nil {
        t.Errorf("os.Open failed: %v", err)
    }

    t.Logf("result: #%v\n", result)

    t.Cleanup(func() {
        _ = f.Close()
    })
}
