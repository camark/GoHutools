package imgutil

import (
	"bytes"
	"image"
	"image/color"
	"path/filepath"
	"testing"
)

// makeSolid returns an w×h opaque solid-color RGBA canvas.
func makeSolid(w, h int, c color.RGBA) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.SetRGBA(x, y, c)
		}
	}
	return img
}

// makeHgrad returns an w×h RGBA canvas whose red channel ramps
// from 0 (left) to 0xFF (right). Used to verify pixel placement.
func makeHgrad(w, h int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r := uint8(float64(x) * 255 / float64(w-1))
			img.SetRGBA(x, y, color.RGBA{R: r, A: 255})
		}
	}
	return img
}

func TestScaleKeepsAspectRatio(t *testing.T) {
	img := makeSolid(400, 200, color.RGBA{255, 0, 0, 255})
	got := Scale(img, 100, 100)
	b := got.Bounds()
	if b.Dx() != 100 || b.Dy() != 50 {
		t.Errorf("Scale(400x200 -> <=100x100) = %dx%d, want 100x50", b.Dx(), b.Dy())
	}
	// enlarging is a no-op (Hutool only downscales)
	if got := Scale(img, 1000, 1000); got.Bounds().Dx() != 400 {
		t.Errorf("Scale up should keep original, got %d", got.Bounds().Dx())
	}
	// zero dimension keeps aspect ratio via the other dimension
	// (width 100 on a 400x200 image => 100x50)
	if got := Scale(img, 100, 0); got.Bounds().Dx() != 100 || got.Bounds().Dy() != 50 {
		t.Errorf("Scale(100, 0) = %dx%d, want 100x50", got.Bounds().Dx(), got.Bounds().Dy())
	}
}

func TestScaleExact(t *testing.T) {
	img := makeSolid(400, 200, color.RGBA{0, 255, 0, 255})
	got := ScaleExact(img, 100, 100)
	b := got.Bounds()
	if b.Dx() != 100 || b.Dy() != 100 {
		t.Errorf("ScaleExact = %dx%d, want 100x100", b.Dx(), b.Dy())
	}
	// color survives the stretch
	c := got.At(50, 50).(color.RGBA)
	if c.R != 0 || c.G != 255 || c.B != 0 {
		t.Errorf("ScaleExact color = %v", c)
	}
}

func TestScaleByPercent(t *testing.T) {
	img := makeSolid(200, 100, color.RGBA{0, 0, 255, 255})
	got := ScaleByPercent(img, 50)
	b := got.Bounds()
	if b.Dx() != 100 || b.Dy() != 50 {
		t.Errorf("ScaleByPercent(50) = %dx%d, want 100x50", b.Dx(), b.Dy())
	}
}

func TestCropMapsPixels(t *testing.T) {
	img := makeHgrad(100, 100)
	got := Crop(img, 10, 10, 20, 20)
	b := got.Bounds()
	if b.Dx() != 20 || b.Dy() != 20 {
		t.Fatalf("Crop size = %dx%d", b.Dx(), b.Dy())
	}
	// pixel (0,0) of crop must equal source (10,10)
	src := img.At(10, 10).(color.RGBA)
	dst := got.At(0, 0).(color.RGBA)
	if src.R != dst.R || src.A != dst.A {
		t.Errorf("Crop(0,0) = %v, want source(10,10) = %v", dst, src)
	}
	// far corner also matches
	if got.At(19, 19).(color.RGBA).R != img.At(29, 29).(color.RGBA).R {
		t.Error("Crop corner mismatch")
	}
}

func TestGreyScaleInvariant(t *testing.T) {
	img := makeHgrad(64, 64)
	got := GreyScale(img)
	for y := 0; y < 64; y += 7 {
		for x := 0; x < 64; x += 7 {
			c := got.At(x, y).(color.Gray)
			// luminance methods keep R == G == B
			if c.Y != c.Y {
			}
		}
	}
	// simpler: every pixel of a grey output is a pure grey (R==G==B)
	for _, p := range []image.Point{{0, 0}, {31, 17}, {63, 63}} {
		c := color.RGBAModel.Convert(got.At(p.X, p.Y)).(color.RGBA)
		if c.R != c.G || c.G != c.B {
			t.Errorf("GreyScale pixel %v not grey: %v", p, c)
		}
	}
	// red image becomes uniform grey with R==G==B (NTSC luma ~76)
	red := makeSolid(10, 10, color.RGBA{255, 0, 0, 255})
	g := color.RGBAModel.Convert(GreyScale(red).At(5, 5)).(color.RGBA)
	if g.R != g.G || g.G != g.B || g.R < 70 || g.R > 82 {
		t.Errorf("GreyScale(red) luma = %v", g)
	}
}

func TestRotate(t *testing.T) {
	img := makeHgrad(200, 100) // gradient blue? no: red ramp left->right

	// 90° CW: output is 100x200; the result's bottom-left (0,199)
	// is the source's bottom-right corner (199,99)
	cw := Rotate(img, 90)
	b := cw.Bounds()
	if b.Dx() != 100 || b.Dy() != 200 {
		t.Fatalf("Rotate90CW = %dx%d, want 100x200", b.Dx(), b.Dy())
	}
	bl := cw.At(0, 199).(color.RGBA)
	srcBR := img.At(199, 99).(color.RGBA)
	if bl.R != srcBR.R {
		t.Errorf("Rotate90CW bottom-left R = %d, want source bottom-right %d", bl.R, srcBR.R)
	}
	// top-left of result comes from source bottom-left (x=0), R=0
	if cw.At(0, 0).(color.RGBA).R != img.At(0, 99).(color.RGBA).R {
		t.Error("Rotate90CW top-left mismatch")
	}

	// 180°: corners swap
	r180 := Rotate(img, 180)
	if r180.At(0, 0).(color.RGBA).R != img.At(199, 99).(color.RGBA).R {
		t.Error("Rotate180 corner mismatch")
	}

	// 270° CCW vs 90° CW are complementary (dimensions equal)
	ccw := Rotate(img, 270)
	if ccw.Bounds().Dx() != 100 {
		t.Errorf("Rotate270 = %dx%d", ccw.Bounds().Dx(), ccw.Bounds().Dy())
	}

	// 360 is identity
	if Rotate(img, 360) != img {
		t.Error("Rotate(360) should return input")
	}
	// negative angle normalized
	if Rotate(img, -90).Bounds().Dx() != 100 {
		t.Errorf("Rotate(-90) = %dx%d", Rotate(img, -90).Bounds().Dx(), Rotate(img, -90).Bounds().Dy())
	}
}

func TestEncodeDecodeRoundTrip(t *testing.T) {
	img := makeSolid(50, 30, color.RGBA{10, 200, 30, 255})

	data, err := Encode(img, FormatPNG)
	if err != nil {
		t.Fatal(err)
	}
	// PNG magic
	if !bytes.HasPrefix(data, []byte{0x89, 'P', 'N', 'G'}) {
		t.Error("PNG magic missing")
	}
	decoded, err := Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	if decoded.Bounds().Dx() != 50 || decoded.Bounds().Dy() != 30 {
		t.Errorf("roundtrip size = %dx%d", decoded.Bounds().Dx(), decoded.Bounds().Dy())
	}

	jpegData, err := Encode(img, FormatJPEG)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.HasPrefix(jpegData, []byte{0xFF, 0xD8}) {
		t.Error("JPEG magic missing")
	}
	if _, err := Decode(jpegData); err != nil {
		t.Errorf("decode jpeg = %v", err)
	}
}

func TestBase64(t *testing.T) {
	img := makeSolid(20, 20, color.RGBA{255, 255, 255, 255})
	b64 := Base64(img, FormatPNG)
	if len(b64) == 0 {
		t.Fatal("empty base64")
	}
	back, err := FromBase64(b64)
	if err != nil {
		t.Fatal(err)
	}
	if back.Bounds().Dx() != 20 || back.Bounds().Dy() != 20 {
		t.Errorf("FromBase64 size = %v", back.Bounds())
	}
}

func TestWriteReadFile(t *testing.T) {
	img := makeSolid(16, 16, color.RGBA{1, 2, 3, 255})
	dir := t.TempDir()

	for _, name := range []string{"out.png", "out.jpg", "out.gif"} {
		path := filepath.Join(dir, name)
		if err := Write(img, path); err != nil {
			t.Fatalf("Write(%s): %v", name, err)
		}
		got, err := Read(path)
		if err != nil {
			t.Fatalf("Read(%s): %v", name, err)
		}
		if got.Bounds().Dx() != 16 || got.Bounds().Dy() != 16 {
			t.Errorf("Read(%s) size = %v", name, got.Bounds())
		}
	}

	// unsupported extension
	if err := Write(img, filepath.Join(dir, "x.bmp")); err == nil {
		t.Error("Write .bmp should error (unsupported)")
	}
}