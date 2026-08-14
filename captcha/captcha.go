package captcha

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"strings"
)

// defaultStrPool is the default character pool for captcha generation
const defaultStrPool = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Captcha is captcha generator
type Captcha struct {
	width   int
	height  int
	length  int
	strPool string
}

// New creates new captcha generator with default settings
func New() *Captcha {
	return &Captcha{
		width:   200,
		height:  80,
		length:  5,
		strPool: defaultStrPool,
	}
}

// SetWidth sets captcha width
func (c *Captcha) SetWidth(width int) *Captcha {
	c.width = width
	return c
}

// SetHeight sets captcha height
func (c *Captcha) SetHeight(height int) *Captcha {
	c.height = height
	return c
}

// SetLength sets captcha code length
func (c *Captcha) SetLength(length int) *Captcha {
	c.length = length
	return c
}

// SetStrPool sets the character pool for captcha generation
func (c *Captcha) SetStrPool(pool string) *Captcha {
	c.strPool = pool
	return c
}

// Generate generates captcha code and image
func (c *Captcha) Generate() (string, image.Image, error) {
	code := c.generateCode()
	img := image.NewRGBA(image.Rect(0, 0, c.width, c.height))

	// Fill background with white
	draw.Draw(img, img.Bounds(), image.White, image.Point{}, draw.Src)

	// Draw captcha characters
	charWidth := c.width / (c.length + 1)
	x := charWidth / 2
	for _, ch := range code {
		drawString(img, string(ch), x, c.height/2+rand.Intn(10)-5)
		x += charWidth
	}

	return code, img, nil
}

// GeneratePNG generates captcha as PNG bytes
func (c *Captcha) GeneratePNG() (string, []byte, error) {
	code, img, err := c.Generate()
	if err != nil {
		return "", nil, err
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return "", nil, err
	}

	return code, buf.Bytes(), nil
}

// Verify verifies captcha code (case sensitive)
func (c *Captcha) Verify(input, expected string) bool {
	return input == expected
}

// VerifyIgnoreCase verifies captcha code case insensitive
func (c *Captcha) VerifyIgnoreCase(input, expected string) bool {
	return strings.EqualFold(input, expected)
}

// generateCode generates random captcha code
func (c *Captcha) generateCode() string {
	pool := []rune(c.strPool)
	code := make([]rune, c.length)
	for i := 0; i < c.length; i++ {
		code[i] = pool[rand.Intn(len(pool))]
	}
	return string(code)
}

// LineCaptcha is line-based captcha
type LineCaptcha struct {
	*Captcha
	lineCount int
}

// NewLine creates new line-based captcha generator
func NewLine() *LineCaptcha {
	return &LineCaptcha{
		Captcha:   New(),
		lineCount: 5,
	}
}

// SetLineCount sets the number of interference lines
func (c *LineCaptcha) SetLineCount(count int) *LineCaptcha {
	c.lineCount = count
	return c
}

// SetWidth sets captcha width
func (c *LineCaptcha) SetWidth(width int) *LineCaptcha {
	c.width = width
	return c
}

// SetHeight sets captcha height
func (c *LineCaptcha) SetHeight(height int) *LineCaptcha {
	c.height = height
	return c
}

// SetLength sets captcha code length
func (c *LineCaptcha) SetLength(length int) *LineCaptcha {
	c.length = length
	return c
}

// SetStrPool sets the character pool
func (c *LineCaptcha) SetStrPool(pool string) *LineCaptcha {
	c.strPool = pool
	return c
}

// Generate generates line-based captcha with interference lines
func (c *LineCaptcha) Generate() (string, image.Image, error) {
	code, img, err := c.Captcha.Generate()
	if err != nil {
		return "", nil, err
	}

	rgba := img.(*image.RGBA)

	// Draw interference lines
	for i := 0; i < c.lineCount; i++ {
		drawLine(rgba,
			randomPoint(c.width, c.height),
			randomPoint(c.width, c.height),
			randomColor())
	}

	return code, rgba, nil
}

// GeneratePNG generates line captcha as PNG bytes
func (c *LineCaptcha) GeneratePNG() (string, []byte, error) {
	code, img, err := c.Captcha.Generate()
	if err != nil {
		return "", nil, err
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return "", nil, err
	}

	return code, buf.Bytes(), nil
}

// CircleCaptcha is circle-based captcha
type CircleCaptcha struct {
	*Captcha
	circleCount int
}

// NewCircle creates new circle-based captcha generator
func NewCircle() *CircleCaptcha {
	return &CircleCaptcha{
		Captcha:     New(),
		circleCount: 5,
	}
}

// SetCircleCount sets the number of interference circles
func (c *CircleCaptcha) SetCircleCount(count int) *CircleCaptcha {
	c.circleCount = count
	return c
}

// SetWidth sets captcha width
func (c *CircleCaptcha) SetWidth(width int) *CircleCaptcha {
	c.width = width
	return c
}

// SetHeight sets captcha height
func (c *CircleCaptcha) SetHeight(height int) *CircleCaptcha {
	c.height = height
	return c
}

// SetLength sets captcha code length
func (c *CircleCaptcha) SetLength(length int) *CircleCaptcha {
	c.length = length
	return c
}

// SetStrPool sets the character pool
func (c *CircleCaptcha) SetStrPool(pool string) *CircleCaptcha {
	c.strPool = pool
	return c
}

// Generate generates circle-based captcha with interference circles
func (c *CircleCaptcha) Generate() (string, image.Image, error) {
	code, img, err := c.Captcha.Generate()
	if err != nil {
		return "", nil, err
	}

	rgba := img.(*image.RGBA)

	// Draw interference circles
	for i := 0; i < c.circleCount; i++ {
		center := randomPoint(c.width, c.height)
		radius := rand.Intn(20) + 10
		drawCircle(rgba, center, radius, randomColor())
	}

	return code, rgba, nil
}

// GeneratePNG generates circle captcha as PNG bytes
func (c *CircleCaptcha) GeneratePNG() (string, []byte, error) {
	code, img, err := c.Captcha.Generate()
	if err != nil {
		return "", nil, err
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return "", nil, err
	}

	return code, buf.Bytes(), nil
}

// ShearCaptcha is shear-based captcha
type ShearCaptcha struct {
	*Captcha
	shearCount int
}

// NewShear creates new shear-based captcha generator
func NewShear() *ShearCaptcha {
	return &ShearCaptcha{
		Captcha:    New(),
		shearCount: 3,
	}
}

// SetShearCount sets the number of shear distortions
func (c *ShearCaptcha) SetShearCount(count int) *ShearCaptcha {
	c.shearCount = count
	return c
}

// SetWidth sets captcha width
func (c *ShearCaptcha) SetWidth(width int) *ShearCaptcha {
	c.width = width
	return c
}

// SetHeight sets captcha height
func (c *ShearCaptcha) SetHeight(height int) *ShearCaptcha {
	c.height = height
	return c
}

// SetLength sets captcha code length
func (c *ShearCaptcha) SetLength(length int) *ShearCaptcha {
	c.length = length
	return c
}

// SetStrPool sets the character pool
func (c *ShearCaptcha) SetStrPool(pool string) *ShearCaptcha {
	c.strPool = pool
	return c
}

// Generate generates shear-based captcha with distortion
func (c *ShearCaptcha) Generate() (string, image.Image, error) {
	code, img, err := c.Captcha.Generate()
	if err != nil {
		return "", nil, err
	}

	rgba := img.(*image.RGBA)

	// Apply shear distortion
	for i := 0; i < c.shearCount; i++ {
		applyShear(rgba, rand.Float64()*0.3-0.15)
	}

	return code, rgba, nil
}

// GeneratePNG generates shear captcha as PNG bytes
func (c *ShearCaptcha) GeneratePNG() (string, []byte, error) {
	code, img, err := c.Captcha.Generate()
	if err != nil {
		return "", nil, err
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return "", nil, err
	}

	return code, buf.Bytes(), nil
}

// Helper functions

// randomColor generates a random color
func randomColor() color.RGBA {
	return color.RGBA{
		R: uint8(rand.Intn(200) + 20),
		G: uint8(rand.Intn(200) + 20),
		B: uint8(rand.Intn(200) + 20),
		A: 255,
	}
}

// randomPoint generates a random point within bounds
func randomPoint(maxX, maxY int) image.Point {
	return image.Point{
		X: rand.Intn(maxX),
		Y: rand.Intn(maxY),
	}
}

// drawString draws a string on the image
func drawString(img *image.RGBA, s string, x, y int) {
	c := randomColor()
	// Simple character drawing - create a visible representation
	for i, ch := range s {
		drawChar(img, ch, x+i*12, y, c)
	}
}

// drawChar draws a simple character representation
func drawChar(img *image.RGBA, ch rune, x, y int, c color.RGBA) {
	// Create a simple 10x14 pixel character representation
	pattern := getCharPattern(ch)
	for dy, row := range pattern {
		for dx, val := range row {
			if val == 1 {
				px := x + dx
				py := y - 7 + dy
				if px >= 0 && px < img.Bounds().Dx() && py >= 0 && py < img.Bounds().Dy() {
					img.SetRGBA(px, py, c)
				}
			}
		}
	}
}

// getCharPattern returns a simple pixel pattern for a character
func getCharPattern(ch rune) [][]int {
	// Simple 10x14 patterns for alphanumeric characters
	patterns := map[rune][][]int{
		'0': {
			{0, 1, 1, 1, 1, 1, 1, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 1, 1, 1, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		'1': {
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 1, 1, 1, 0, 0, 0, 0, 0},
			{0, 1, 1, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 1, 1, 1, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		'2': {
			{0, 1, 1, 1, 1, 1, 1, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		'3': {
			{0, 1, 1, 1, 1, 1, 1, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 1, 1, 1, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		'4': {
			{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 1, 1, 0, 0, 1, 1, 0, 0, 0},
			{1, 1, 0, 0, 0, 1, 1, 0, 0, 0},
			{1, 1, 1, 1, 1, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		'5': {
			{1, 1, 1, 1, 1, 1, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 1, 1, 1, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		'6': {
			{0, 1, 1, 1, 1, 1, 1, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1, 1, 1, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 1, 1, 1, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		'7': {
			{1, 1, 1, 1, 1, 1, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		'8': {
			{0, 1, 1, 1, 1, 1, 1, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 1, 1, 1, 1, 1, 1, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 1, 1, 1, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		'9': {
			{0, 1, 1, 1, 1, 1, 1, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 1, 1, 1, 1, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 1, 1, 1, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	if p, ok := patterns[ch]; ok {
		return p
	}

	// Default pattern for unknown characters - simple rectangle
	return [][]int{
		{0, 1, 1, 1, 1, 1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
		{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
		{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
		{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
		{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
		{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
		{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
		{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
		{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
		{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
		{1, 1, 0, 0, 0, 0, 1, 1, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
}

// drawLine draws a line on the image
func drawLine(img *image.RGBA, start, end image.Point, c color.RGBA) {
	dx := abs(end.X - start.X)
	dy := abs(end.Y - start.Y)
	sx := 1
	sy := 1
	if start.X > end.X {
		sx = -1
	}
	if start.Y > end.Y {
		sy = -1
	}
	err := dx - dy
	x, y := start.X, start.Y

	for {
		if x >= 0 && x < img.Bounds().Dx() && y >= 0 && y < img.Bounds().Dy() {
			img.SetRGBA(x, y, c)
		}
		if x == end.X && y == end.Y {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x += sx
		}
		if e2 < dx {
			err += dx
			y += sy
		}
	}
}

// drawCircle draws a circle outline on the image
func drawCircle(img *image.RGBA, center image.Point, radius int, c color.RGBA) {
	x := radius
	y := 0
	err := 0

	for x >= y {
		drawCirclePoints(img, center, x, y, c)

		if err <= 0 {
			y++
			err += 2*y + 1
		}
		if err > 0 {
			x--
			err -= 2*x + 1
		}
	}
}

// drawCirclePoints draws the 8 octants of a circle
func drawCirclePoints(img *image.RGBA, center image.Point, x, y int, c color.RGBA) {
	points := []image.Point{
		{center.X + x, center.Y + y},
		{center.X - x, center.Y + y},
		{center.X + x, center.Y - y},
		{center.X - x, center.Y - y},
		{center.X + y, center.Y + x},
		{center.X - y, center.Y + x},
		{center.X + y, center.Y - x},
		{center.X - y, center.Y - x},
	}

	for _, p := range points {
		if p.X >= 0 && p.X < img.Bounds().Dx() && p.Y >= 0 && p.Y < img.Bounds().Dy() {
			img.SetRGBA(p.X, p.Y, c)
		}
	}
}

// applyShear applies shear transformation to the image
func applyShear(img *image.RGBA, factor float64) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Create temporary image
	tmp := image.NewRGBA(bounds)
	draw.Draw(tmp, bounds, img, bounds.Min, draw.Src)

	// Apply horizontal shear
	for y := 0; y < height; y++ {
		offset := int(float64(y-height/2) * factor * 10)
		for x := 0; x < width; x++ {
			srcX := x - offset
			if srcX >= 0 && srcX < width {
				tmp.SetRGBA(x, y, img.RGBAAt(srcX, y))
			} else {
				tmp.SetRGBA(x, y, color.RGBA{255, 255, 255, 255})
			}
		}
	}

	// Copy back
	draw.Draw(img, bounds, tmp, bounds.Min, draw.Src)
}

// abs returns absolute value of integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
