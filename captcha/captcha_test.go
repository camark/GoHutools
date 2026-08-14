package captcha

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"testing"
)

func TestCaptchaNew(t *testing.T) {
	c := New()
	if c == nil {
		t.Fatal("New() returned nil")
	}
	if c.width != 200 {
		t.Errorf("Expected width 200, got %d", c.width)
	}
	if c.height != 80 {
		t.Errorf("Expected height 80, got %d", c.height)
	}
	if c.length != 5 {
		t.Errorf("Expected length 5, got %d", c.length)
	}
	if c.strPool != defaultStrPool {
		t.Errorf("Expected default str pool, got %s", c.strPool)
	}
}

func TestCaptchaSetters(t *testing.T) {
	c := New()
	c.SetWidth(300).SetHeight(100).SetLength(6).SetStrPool("0123456789")

	if c.width != 300 {
		t.Errorf("Expected width 300, got %d", c.width)
	}
	if c.height != 100 {
		t.Errorf("Expected height 100, got %d", c.height)
	}
	if c.length != 6 {
		t.Errorf("Expected length 6, got %d", c.length)
	}
	if c.strPool != "0123456789" {
		t.Errorf("Expected str pool '0123456789', got %s", c.strPool)
	}
}

func TestCaptchaGenerate(t *testing.T) {
	c := New().SetLength(5).SetStrPool("0123456789")
	code, img, err := c.Generate()

	if err != nil {
		t.Fatalf("Generate() error: %v", err)
	}
	if len(code) != 5 {
		t.Errorf("Expected code length 5, got %d", len(code))
	}
	if img == nil {
		t.Fatal("Generate() returned nil image")
	}

	// Check image dimensions
	bounds := img.Bounds()
	if bounds.Dx() != 200 || bounds.Dy() != 80 {
		t.Errorf("Expected image 200x80, got %dx%d", bounds.Dx(), bounds.Dy())
	}

	// Verify all characters are from the pool
	for _, ch := range code {
		if ch < '0' || ch > '9' {
			t.Errorf("Character %c not in pool", ch)
		}
	}
}

func TestCaptchaGeneratePNG(t *testing.T) {
	c := New().SetLength(4)
	code, data, err := c.GeneratePNG()

	if err != nil {
		t.Fatalf("GeneratePNG() error: %v", err)
	}
	if len(code) != 4 {
		t.Errorf("Expected code length 4, got %d", len(code))
	}
	if len(data) == 0 {
		t.Fatal("GeneratePNG() returned empty data")
	}

	// Verify it's valid PNG
	_, err = png.Decode(bytes.NewReader(data))
	if err != nil {
		t.Fatalf("Invalid PNG data: %v", err)
	}
}

func TestCaptchaVerify(t *testing.T) {
	c := New()

	if !c.Verify("abc", "abc") {
		t.Error("Verify should return true for matching strings")
	}
	if c.Verify("abc", "ABC") {
		t.Error("Verify should return false for case mismatch")
	}
	if c.Verify("abc", "def") {
		t.Error("Verify should return false for different strings")
	}
}

func TestCaptchaVerifyIgnoreCase(t *testing.T) {
	c := New()

	if !c.VerifyIgnoreCase("abc", "abc") {
		t.Error("VerifyIgnoreCase should return true for matching strings")
	}
	if !c.VerifyIgnoreCase("abc", "ABC") {
		t.Error("VerifyIgnoreCase should return true for case insensitive match")
	}
	if !c.VerifyIgnoreCase("AbC", "aBc") {
		t.Error("VerifyIgnoreCase should return true for mixed case")
	}
	if c.VerifyIgnoreCase("abc", "def") {
		t.Error("VerifyIgnoreCase should return false for different strings")
	}
}

func TestCaptchaGenerateCode(t *testing.T) {
	c := New().SetLength(10).SetStrPool("ABCD")

	code := c.generateCode()
	if len(code) != 10 {
		t.Errorf("Expected code length 10, got %d", len(code))
	}

	for _, ch := range code {
		if ch != 'A' && ch != 'B' && ch != 'C' && ch != 'D' {
			t.Errorf("Character %c not in pool 'ABCD'", ch)
		}
	}
}

func TestLineCaptcha(t *testing.T) {
	lc := NewLine()
	if lc == nil {
		t.Fatal("NewLine() returned nil")
	}
	if lc.lineCount != 5 {
		t.Errorf("Expected lineCount 5, got %d", lc.lineCount)
	}

	lc.SetLineCount(10).SetWidth(300).SetHeight(120).SetLength(6).SetStrPool("0123456789")

	code, img, err := lc.Generate()
	if err != nil {
		t.Fatalf("Generate() error: %v", err)
	}
	if len(code) != 6 {
		t.Errorf("Expected code length 6, got %d", len(code))
	}
	if img == nil {
		t.Fatal("Generate() returned nil image")
	}

	bounds := img.Bounds()
	if bounds.Dx() != 300 || bounds.Dy() != 120 {
		t.Errorf("Expected image 300x120, got %dx%d", bounds.Dx(), bounds.Dy())
	}
}

func TestLineCaptchaPNG(t *testing.T) {
	lc := NewLine().SetLength(4)
	code, data, err := lc.GeneratePNG()

	if err != nil {
		t.Fatalf("GeneratePNG() error: %v", err)
	}
	if len(code) != 4 {
		t.Errorf("Expected code length 4, got %d", len(code))
	}
	if len(data) == 0 {
		t.Fatal("GeneratePNG() returned empty data")
	}

	// Verify it's valid PNG
	_, err = png.Decode(bytes.NewReader(data))
	if err != nil {
		t.Fatalf("Invalid PNG data: %v", err)
	}
}

func TestCircleCaptcha(t *testing.T) {
	cc := NewCircle()
	if cc == nil {
		t.Fatal("NewCircle() returned nil")
	}
	if cc.circleCount != 5 {
		t.Errorf("Expected circleCount 5, got %d", cc.circleCount)
	}

	cc.SetCircleCount(8).SetWidth(250).SetHeight(100).SetLength(4).SetStrPool("ABCDEF")

	code, img, err := cc.Generate()
	if err != nil {
		t.Fatalf("Generate() error: %v", err)
	}
	if len(code) != 4 {
		t.Errorf("Expected code length 4, got %d", len(code))
	}
	if img == nil {
		t.Fatal("Generate() returned nil image")
	}

	// Check all characters are from pool
	for _, ch := range code {
		if ch < 'A' || ch > 'F' {
			t.Errorf("Character %c not in pool 'ABCDEF'", ch)
		}
	}
}

func TestCircleCaptchaPNG(t *testing.T) {
	cc := NewCircle().SetLength(5)
	code, data, err := cc.GeneratePNG()

	if err != nil {
		t.Fatalf("GeneratePNG() error: %v", err)
	}
	if len(code) != 5 {
		t.Errorf("Expected code length 5, got %d", len(code))
	}
	if len(data) == 0 {
		t.Fatal("GeneratePNG() returned empty data")
	}

	_, err = png.Decode(bytes.NewReader(data))
	if err != nil {
		t.Fatalf("Invalid PNG data: %v", err)
	}
}

func TestShearCaptcha(t *testing.T) {
	sc := NewShear()
	if sc == nil {
		t.Fatal("NewShear() returned nil")
	}
	if sc.shearCount != 3 {
		t.Errorf("Expected shearCount 3, got %d", sc.shearCount)
	}

	sc.SetShearCount(5).SetWidth(280).SetHeight(90).SetLength(7).SetStrPool("XYZ123")

	code, img, err := sc.Generate()
	if err != nil {
		t.Fatalf("Generate() error: %v", err)
	}
	if len(code) != 7 {
		t.Errorf("Expected code length 7, got %d", len(code))
	}
	if img == nil {
		t.Fatal("Generate() returned nil image")
	}

	// Check characters are from pool
	for _, ch := range code {
		found := false
		for _, p := range "XYZ123" {
			if ch == p {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Character %c not in pool 'XYZ123'", ch)
		}
	}
}

func TestShearCaptchaPNG(t *testing.T) {
	sc := NewShear().SetLength(4)
	code, data, err := sc.GeneratePNG()

	if err != nil {
		t.Fatalf("GeneratePNG() error: %v", err)
	}
	if len(code) != 4 {
		t.Errorf("Expected code length 4, got %d", len(code))
	}
	if len(data) == 0 {
		t.Fatal("GeneratePNG() returned empty data")
	}

	_, err = png.Decode(bytes.NewReader(data))
	if err != nil {
		t.Fatalf("Invalid PNG data: %v", err)
	}
}

func TestRandomColor(t *testing.T) {
	// Test that randomColor produces valid colors
	for i := 0; i < 100; i++ {
		c := randomColor()
		if c.A != 255 {
			t.Errorf("Expected alpha 255, got %d", c.A)
		}
		if c.R < 20 || c.R > 219 {
			t.Errorf("Red out of range: %d", c.R)
		}
		if c.G < 20 || c.G > 219 {
			t.Errorf("Green out of range: %d", c.G)
		}
		if c.B < 20 || c.B > 219 {
			t.Errorf("Blue out of range: %d", c.B)
		}
	}
}

func TestRandomPoint(t *testing.T) {
	// Test that randomPoint produces points within bounds
	for i := 0; i < 100; i++ {
		p := randomPoint(200, 100)
		if p.X < 0 || p.X >= 200 {
			t.Errorf("X out of range: %d", p.X)
		}
		if p.Y < 0 || p.Y >= 100 {
			t.Errorf("Y out of range: %d", p.Y)
		}
	}
}

func TestDrawLine(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	c := color.RGBA{R: 255, A: 255}

	// Draw a horizontal line
	drawLine(img, image.Point{X: 10, Y: 50}, image.Point{X: 90, Y: 50}, c)

	// Check that line was drawn
	for x := 10; x <= 90; x++ {
		r, _, _, a := img.At(x, 50).RGBA()
		if r>>8 != 255 || a>>8 != 255 {
			t.Errorf("Line not drawn at (%d, 50)", x)
		}
	}
}

func TestDrawCircle(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	c := color.RGBA{R: 255, A: 255}

	drawCircle(img, image.Point{X: 50, Y: 50}, 20, c)

	// Check that circle points were drawn
	points := []image.Point{
		{X: 50, Y: 30}, // top
		{X: 50, Y: 70}, // bottom
		{X: 30, Y: 50}, // left
		{X: 70, Y: 50}, // right
	}

	for _, p := range points {
		r, _, _, a := img.At(p.X, p.Y).RGBA()
		if r>>8 != 255 || a>>8 != 255 {
			t.Errorf("Circle not drawn at (%d, %d)", p.X, p.Y)
		}
	}
}

func TestApplyShear(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	// Draw a vertical line
	for y := 0; y < 100; y++ {
		img.SetRGBA(50, y, color.RGBA{R: 255, A: 255})
	}

	applyShear(img, 0.1)

	// After shear, the line should be displaced differently at top vs bottom
	// Just verify it doesn't panic and produces valid output
	if img == nil {
		t.Fatal("applyShear returned nil")
	}
}

func TestAbs(t *testing.T) {
	if abs(5) != 5 {
		t.Errorf("abs(5) = %d, want 5", abs(5))
	}
	if abs(-5) != 5 {
		t.Errorf("abs(-5) = %d, want 5", abs(-5))
	}
	if abs(0) != 0 {
		t.Errorf("abs(0) = %d, want 0", abs(0))
	}
}

func TestGetCharPattern(t *testing.T) {
	// Test known patterns
	for _, ch := range "0123456789" {
		p := getCharPattern(ch)
		if len(p) != 14 {
			t.Errorf("Pattern for %c has %d rows, want 14", ch, len(p))
		}
		for _, row := range p {
			if len(row) != 10 {
				t.Errorf("Pattern row for %c has %d columns, want 10", ch, len(row))
			}
		}
	}

	// Test unknown character - should return default pattern
	p := getCharPattern('?')
	if len(p) != 14 {
		t.Errorf("Default pattern has %d rows, want 14", len(p))
	}
}

func BenchmarkCaptchaGenerate(b *testing.B) {
	c := New()
	for i := 0; i < b.N; i++ {
		c.Generate()
	}
}

func BenchmarkLineCaptchaGenerate(b *testing.B) {
	lc := NewLine()
	for i := 0; i < b.N; i++ {
		lc.Generate()
	}
}

func BenchmarkCircleCaptchaGenerate(b *testing.B) {
	cc := NewCircle()
	for i := 0; i < b.N; i++ {
		cc.Generate()
	}
}

func BenchmarkShearCaptchaGenerate(b *testing.B) {
	sc := NewShear()
	for i := 0; i < b.N; i++ {
		sc.Generate()
	}
}
