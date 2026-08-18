package imgutil

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"strings"
)

// 本包对标 hutool 的 ImgUtil，提供图像读取、缩放、裁剪、旋转、
// 灰度和格式转换等常用操作，全部基于标准库 image 家族实现。

// Format identifies an image encoding.
type Format int

const (
	FormatAuto Format = iota
	FormatPNG
	FormatJPEG
	FormatGIF
)

// Read decodes an image file, auto-detecting PNG/JPEG/GIF by content.
func Read(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()
	img, _, err := image.Decode(f)
	return img, err
}

// Write encodes img to path, choosing the encoder from the file
// extension (.png/.jpg/.jpeg/.gif).
func Write(img image.Image, path string) error {
	format := formatFromExt(path)
	if format == FormatAuto {
		return errors.New("imgutil: unsupported image extension: " + path)
	}
	data, err := Encode(img, format)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// Encode encodes img in the given format to bytes.
func Encode(img image.Image, format Format) ([]byte, error) {
	var buf bytes.Buffer
	switch format {
	case FormatPNG:
		err := png.Encode(&buf, img)
		return buf.Bytes(), err
	case FormatJPEG:
		err := jpeg.Encode(&buf, img, nil)
		return buf.Bytes(), err
	case FormatGIF:
		err := gif.Encode(&buf, img, nil)
		return buf.Bytes(), err
	default:
		return nil, errors.New("imgutil: unsupported format")
	}
}

// Decode decodes an image from bytes, auto-detecting the format.
func Decode(data []byte) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	return img, err
}

// Base64 encodes img in the given format and returns standard base64 text.
func Base64(img image.Image, format Format) string {
	data, err := Encode(img, format)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(data)
}

// FromBase64 decodes an image from standard base64 text.
func FromBase64(s string) (image.Image, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return Decode(data)
}

// Scale downscales img to fit within width×height, preserving the aspect
// ratio (Hutool scale semantics: never enlarges). A non-positive width or
// height keeps the original dimension in that axis.
func Scale(img image.Image, width, height int) image.Image {
	b := img.Bounds()
	if width <= 0 {
		width = b.Dx()
	}
	if height <= 0 {
		height = b.Dy()
	}
	ratio := math.Min(float64(width)/float64(b.Dx()), float64(height)/float64(b.Dy()))
	if ratio >= 1 {
		return img
	}
	return scaleTo(img, int(float64(b.Dx())*ratio), int(float64(b.Dy())*ratio))
}

// ScaleByPercent scales img by a percentage (50 = half size).
func ScaleByPercent(img image.Image, percent float64) image.Image {
	b := img.Bounds()
	w := int(float64(b.Dx()) * percent / 100)
	h := int(float64(b.Dy()) * percent / 100)
	return scaleTo(img, w, h)
}

// ScaleExact stretches img to the exact target size (aspect may distort).
func ScaleExact(img image.Image, width, height int) image.Image {
	if width <= 0 || height <= 0 {
		return img
	}
	return scaleTo(img, width, height)
}

func scaleTo(img image.Image, w, h int) image.Image {
	if w <= 0 || h <= 0 {
		return img
	}
	return scaleNearest(img, w, h)
}

// scaleNearest resizes img to w×h using nearest-neighbor sampling,
// keeping this package free of external dependencies (the standard
// library image/draw lacks scaling interpolators — those live in
// golang.org/x/image/draw).
func scaleNearest(img image.Image, w, h int) image.Image {
	b := img.Bounds()
	srcW, srcH := b.Dx(), b.Dy()
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	for dy := 0; dy < h; dy++ {
		sy := dy * srcH / h
		for dx := 0; dx < w; dx++ {
			sx := dx * srcW / w
			dst.SetRGBA(dx, dy, toRGBA(img.At(b.Min.X+sx, b.Min.Y+sy)))
		}
	}
	return dst
}

// Crop returns the w×h rectangle starting at (x, y), clipped to the
// source bounds. Out-of-bounds areas are dropped.
func Crop(img image.Image, x, y, width, height int) image.Image {
	b := img.Bounds()
	r := image.Rect(x, y, x+width, y+height).Intersect(b)
	if r.Empty() {
		return image.NewRGBA(image.Rect(0, 0, 0, 0))
	}
	dst := image.NewRGBA(image.Rect(0, 0, r.Dx(), r.Dy()))
	draw.Draw(dst, dst.Bounds(), img, r.Min, draw.Src)
	return dst
}

// GreyScale converts img to luminance (grayscale) using the standard
// NTSC coefficients.
func GreyScale(img image.Image) image.Image {
	b := img.Bounds()
	dst := image.NewGray(image.Rect(0, 0, b.Dx(), b.Dy()))
	for y := 0; y < b.Dy(); y++ {
		for x := 0; x < b.Dx(); x++ {
			c := color.GrayModel.Convert(img.At(b.Min.X+x, b.Min.Y+y))
			dst.SetGray(x, y, c.(color.Gray))
		}
	}
	return dst
}

// Rotate rotates img by angle degrees; only multiples of 90 are
// supported. angle is normalized into [0, 360).
func Rotate(img image.Image, angle int) image.Image {
	switch ((angle % 360) + 360) % 360 {
	case 90:
		return Rotate90CW(img)
	case 180:
		return Rotate180(img)
	case 270:
		return Rotate90CCW(img)
	default:
		return img
	}
}

// Rotate90CW rotates img 90 degrees clockwise.
// dst(x, y) = src(y, H-1-x): the source bottom row becomes dst column 0.
func Rotate90CW(img image.Image) image.Image {
	b := img.Bounds()
	W, H := b.Dx(), b.Dy()
	dst := image.NewRGBA(image.Rect(0, 0, H, W))
	for x := 0; x < H; x++ {
		for y := 0; y < W; y++ {
			dst.SetRGBA(x, y, toRGBA(img.At(b.Min.X+y, b.Min.Y+H-1-x)))
		}
	}
	return dst
}

// Rotate90CCW rotates img 90 degrees counter-clockwise.
// dst(x, y) = src(W-1-y, x): the source right edge becomes dst row 0.
func Rotate90CCW(img image.Image) image.Image {
	b := img.Bounds()
	W, H := b.Dx(), b.Dy()
	dst := image.NewRGBA(image.Rect(0, 0, H, W))
	for x := 0; x < H; x++ {
		for y := 0; y < W; y++ {
			dst.SetRGBA(x, y, toRGBA(img.At(b.Min.X+W-1-y, b.Min.Y+x)))
		}
	}
	return dst
}

// Rotate180 rotates img 180 degrees.
func Rotate180(img image.Image) image.Image {
	b := img.Bounds()
	W, H := b.Dx(), b.Dy()
	dst := image.NewRGBA(image.Rect(0, 0, W, H))
	for x := 0; x < W; x++ {
		for y := 0; y < H; y++ {
			dst.SetRGBA(x, y, toRGBA(img.At(b.Min.X+W-1-x, b.Min.Y+H-1-y)))
		}
	}
	return dst
}

// Note: text watermarking is intentionally omitted — it needs the
// golang.org/x/image font packages, while this package stays on the
// pure standard library.

func formatFromExt(path string) Format {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".png":
		return FormatPNG
	case ".jpg", ".jpeg":
		return FormatJPEG
	case ".gif":
		return FormatGIF
	default:
		return FormatAuto
	}
}

func toRGBA(c color.Color) color.RGBA {
	r, g, b, a := c.RGBA()
	return color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
}