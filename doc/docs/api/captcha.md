---
title: Captcha
---

# Captcha

验证码生成模块，提供多种风格的验证码图片生成，包括基础验证码、干扰线验证码、干扰圆验证码和扭曲验证码。支持 PNG 格式输出。

## 导入

```go
import "github.com/gongm/gohutool/captcha"
```

## 基础验证码

### New

```go
func New() *Captcha
```

创建默认设置的验证码生成器（200x80 像素，5 位字符）。

**示例:**

```go
c := captcha.New()
code, img, err := c.Generate()
```

### Captcha.SetWidth

```go
func (c *Captcha) SetWidth(width int) *Captcha
```

设置验证码图片宽度。

### Captcha.SetHeight

```go
func (c *Captcha) SetHeight(height int) *Captcha
```

设置验证码图片高度。

### Captcha.SetLength

```go
func (c *Captcha) SetLength(length int) *Captcha
```

设置验证码字符长度。

### Captcha.SetStrPool

```go
func (c *Captcha) SetStrPool(pool string) *Captcha
```

设置验证码字符池。默认包含数字和大小写字母。

**示例:**

```go
c := captcha.New().
    SetWidth(300).
    SetHeight(100).
    SetLength(6).
    SetStrPool("0123456789")
```

### Captcha.Generate

```go
func (c *Captcha) Generate() (string, image.Image, error)
```

生成验证码，返回验证码文本、图片对象和错误信息。

### Captcha.GeneratePNG

```go
func (c *Captcha) GeneratePNG() (string, []byte, error)
```

生成验证码，返回验证码文本、PNG 字节数据和错误信息。适合直接写入 HTTP 响应。

**示例:**

```go
c := captcha.New()
code, pngBytes, err := c.GeneratePNG()
if err != nil {
    log.Fatal(err)
}
// 将验证码文本存入 session
os.WriteFile("captcha.png", pngBytes, 0644)
```

### Captcha.Verify

```go
func (c *Captcha) Verify(input, expected string) bool
```

验证用户输入是否匹配（区分大小写）。

### Captcha.VerifyIgnoreCase

```go
func (c *Captcha) VerifyIgnoreCase(input, expected string) bool
```

验证用户输入是否匹配（不区分大小写）。

## 干扰线验证码

在基础验证码上添加随机干扰线，增加机器识别难度。

### NewLine

```go
func NewLine() *LineCaptcha
```

创建干扰线验证码生成器。

### LineCaptcha.SetLineCount

```go
func (c *LineCaptcha) SetLineCount(count int) *LineCaptcha
```

设置干扰线数量，默认为 5。

**示例:**

```go
c := captcha.NewLine().
    SetWidth(250).
    SetHeight(80).
    SetLength(4).
    SetLineCount(8)
code, img, err := c.Generate()
```

### LineCaptcha.Generate

```go
func (c *LineCaptcha) Generate() (string, image.Image, error)
```

生成带干扰线的验证码图片。

### LineCaptcha.GeneratePNG

```go
func (c *LineCaptcha) GeneratePNG() (string, []byte, error)
```

生成带干扰线的验证码 PNG 数据。

## 干扰圆验证码

在基础验证码上添加随机干扰圆，增加机器识别难度。

### NewCircle

```go
func NewCircle() *CircleCaptcha
```

创建干扰圆验证码生成器。

### CircleCaptcha.SetCircleCount

```go
func (c *CircleCaptcha) SetCircleCount(count int) *CircleCaptcha
```

设置干扰圆数量，默认为 5。

**示例:**

```go
c := captcha.NewCircle().
    SetWidth(200).
    SetHeight(80).
    SetCircleCount(10)
code, img, err := c.Generate()
```

### CircleCaptcha.Generate

```go
func (c *CircleCaptcha) Generate() (string, image.Image, error)
```

生成带干扰圆的验证码图片。

### CircleCaptcha.GeneratePNG

```go
func (c *CircleCaptcha) GeneratePNG() (string, []byte, error)
```

生成带干扰圆的验证码 PNG 数据。

## 扭曲验证码

通过剪切扭曲变换增加验证码的复杂度。

### NewShear

```go
func NewShear() *ShearCaptcha
```

创建扭曲验证码生成器。

### ShearCaptcha.SetShearCount

```go
func (c *ShearCaptcha) SetShearCount(count int) *ShearCaptcha
```

设置扭曲变换次数，默认为 3。

**示例:**

```go
c := captcha.NewShear().
    SetWidth(200).
    SetHeight(80).
    SetShearCount(5)
code, img, err := c.Generate()
```

### ShearCaptcha.Generate

```go
func (c *ShearCaptcha) Generate() (string, image.Image, error)
```

生成带扭曲效果的验证码图片。

### ShearCaptcha.GeneratePNG

```go
func (c *ShearCaptcha) GeneratePNG() (string, []byte, error)
```

生成带扭曲效果的验证码 PNG 数据。

## 完整示例

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/gongm/gohutool/captcha"
)

func main() {
    // HTTP 服务中使用验证码
    http.HandleFunc("/captcha", func(w http.ResponseWriter, r *http.Request) {
        c := captcha.NewLine().
            SetWidth(200).
            SetHeight(80).
            SetLength(4).
            SetLineCount(6)

        code, pngBytes, err := c.GeneratePNG()
        if err != nil {
            http.Error(w, "生成验证码失败", 500)
            return
        }

        // 将验证码存入 session（示例中仅打印）
        fmt.Println("验证码:", code)

        w.Header().Set("Content-Type", "image/png")
        w.Write(pngBytes)
    })

    http.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
        input := r.FormValue("code")
        expected := "从 session 获取"

        c := captcha.New()
        if c.VerifyIgnoreCase(input, expected) {
            fmt.Fprintln(w, "验证通过")
        } else {
            fmt.Fprintln(w, "验证失败")
        }
    })

    http.ListenAndServe(":8080", nil)
}
```
