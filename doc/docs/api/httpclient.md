---
title: httpclient
---

# httpclient

HTTP 客户端工具包，提供链式调用风格的 HTTP 请求构建器，支持 GET、POST、PUT、DELETE 等方法，以及 JSON、Form、Cookie、认证等功能。

## 导入

```go
import "github.com/gongm/gohutool/httpclient"
```

## 快速函数

### Get

```go
func Get(rawURL string) (*Response, error)
```

发送 GET 请求。

**示例:**

```go
resp, err := httpclient.Get("https://api.example.com/users")
if err != nil {
    log.Fatal(err)
}
body, _ := resp.BodyString()
fmt.Println(body)
```

### Post

```go
func Post(rawURL string, body interface{}) (*Response, error)
```

发送 POST 请求。

**示例:**

```go
resp, err := httpclient.Post("https://api.example.com/users", strings.NewReader("data"))
```

### PostForm

```go
func PostForm(rawURL string, data map[string]string) (*Response, error)
```

发送带有表单数据的 POST 请求。

**示例:**

```go
resp, err := httpclient.PostForm("https://api.example.com/login", map[string]string{
    "username": "admin",
    "password": "123456",
})
```

### PostJSON

```go
func PostJSON(rawURL string, data interface{}) (*Response, error)
```

发送带有 JSON 请求体的 POST 请求。

**示例:**

```go
resp, err := httpclient.PostJSON("https://api.example.com/users", map[string]string{
    "name": "John",
    "email": "john@example.com",
})
```

### Put

```go
func Put(rawURL string, body interface{}) (*Response, error)
```

发送 PUT 请求。

**示例:**

```go
resp, err := httpclient.Put("https://api.example.com/users/1", strings.NewReader("data"))
```

### Delete

```go
func Delete(rawURL string) (*Response, error)
```

发送 DELETE 请求。

**示例:**

```go
resp, err := httpclient.Delete("https://api.example.com/users/1")
```

## Client

### New

```go
func New() *Client
```

创建新的 HTTP 客户端。

**示例:**

```go
client := httpclient.New()
```

### NewWithTimeout

```go
func NewWithTimeout(timeout time.Duration) *Client
```

创建带有超时时间的 HTTP 客户端。

**示例:**

```go
client := httpclient.NewWithTimeout(30 * time.Second)
```

### SetTimeout

```go
func (c *Client) SetTimeout(d time.Duration) *Client
```

设置请求超时时间。

**示例:**

```go
client := httpclient.New().SetTimeout(10 * time.Second)
```

### SetHeader

```go
func (c *Client) SetHeader(key, value string) *Client
```

设置默认请求头。

**示例:**

```go
client := httpclient.New().SetHeader("X-Custom", "value")
```

### SetHeaders

```go
func (c *Client) SetHeaders(headers map[string]string) *Client
```

批量设置默认请求头。

**示例:**

```go
client := httpclient.New().SetHeaders(map[string]string{
    "X-App-Id":     "myapp",
    "X-App-Secret": "secret",
})
```

### SetCookie

```go
func (c *Client) SetCookie(cookie *http.Cookie) *Client
```

设置 Cookie。

**示例:**

```go
cookie := &http.Cookie{Name: "session", Value: "abc123"}
client := httpclient.New().SetCookie(cookie)
```

### SetBasicAuth

```go
func (c *Client) SetBasicAuth(username, password string) *Client
```

设置 Basic 认证头。

**示例:**

```go
client := httpclient.New().SetBasicAuth("admin", "password")
```

### SetBearerToken

```go
func (c *Client) SetBearerToken(token string) *Client
```

设置 Bearer Token 认证头。

**示例:**

```go
client := httpclient.New().SetBearerToken("eyJhbGciOiJIUzI1NiIs...")
```

### Get

```go
func (c *Client) Get(rawURL string) *Request
```

创建 GET 请求。

**示例:**

```go
resp, err := httpclient.New().Get("https://api.example.com/users").Do()
```

### Post

```go
func (c *Client) Post(rawURL string) *Request
```

创建 POST 请求。

**示例:**

```go
resp, err := httpclient.New().Post("https://api.example.com/users").
    JSON(map[string]string{"name": "John"}).
    Do()
```

### Put

```go
func (c *Client) Put(rawURL string) *Request
```

创建 PUT 请求。

### Delete

```go
func (c *Client) Delete(rawURL string) *Request
```

创建 DELETE 请求。

### Patch

```go
func (c *Client) Patch(rawURL string) *Request
```

创建 PATCH 请求。

### Head

```go
func (c *Client) Head(rawURL string) *Request
```

创建 HEAD 请求。

### Options

```go
func (c *Client) Options(rawURL string) *Request
```

创建 OPTIONS 请求。

## Request

### Header

```go
func (r *Request) Header(key, value string) *Request
```

设置单个请求头。

**示例:**

```go
resp, err := httpclient.New().Get("https://api.example.com/users").
    Header("Accept", "application/json").
    Do()
```

### Headers

```go
func (r *Request) Headers(headers map[string]string) *Request
```

批量设置请求头。

### Query

```go
func (r *Request) Query(key, value string) *Request
```

添加 URL 查询参数。

**示例:**

```go
resp, err := httpclient.New().Get("https://api.example.com/users").
    Query("page", "1").
    Query("size", "10").
    Do()
```

### Queries

```go
func (r *Request) Queries(params map[string]string) *Request
```

批量添加 URL 查询参数。

**示例:**

```go
resp, err := httpclient.New().Get("https://api.example.com/users").
    Queries(map[string]string{"page": "1", "size": "10"}).
    Do()
```

### Body

```go
func (r *Request) Body(body interface{}) *Request
```

设置请求体，支持 string、[]byte、io.Reader 或可 JSON 序列化的对象。

**示例:**

```go
resp, err := httpclient.New().Post("https://api.example.com/users").
    Body(strings.NewReader("data")).
    Do()
```

### BodyString

```go
func (r *Request) BodyString(body string) *Request
```

设置字符串请求体。

### BodyBytes

```go
func (r *Request) BodyBytes(body []byte) *Request
```

设置字节请求体。

### BodyReader

```go
func (r *Request) BodyReader(body io.Reader) *Request
```

设置 io.Reader 请求体。

### Form

```go
func (r *Request) Form(data map[string]string) *Request
```

设置表单数据（自动设置 Content-Type 为 application/x-www-form-urlencoded）。

**示例:**

```go
resp, err := httpclient.New().Post("https://api.example.com/login").
    Form(map[string]string{"username": "admin", "password": "123456"}).
    Do()
```

### FormValues

```go
func (r *Request) FormValues(data url.Values) *Request
```

设置 url.Values 表单数据。

### JSON

```go
func (r *Request) JSON(data interface{}) *Request
```

设置 JSON 请求体（自动设置 Content-Type 为 application/json; charset=utf-8）。

**示例:**

```go
resp, err := httpclient.New().Post("https://api.example.com/users").
    JSON(map[string]interface{}{
        "name":  "John",
        "age":   30,
        "email": "john@example.com",
    }).
    Do()
```

### Cookie

```go
func (r *Request) Cookie(cookie *http.Cookie) *Request
```

添加请求 Cookie。

### Timeout

```go
func (r *Request) Timeout(d time.Duration) *Request
```

设置单次请求超时。

**示例:**

```go
resp, err := httpclient.New().Get("https://api.example.com/slow").
    Timeout(5 * time.Second).
    Do()
```

### ContentType

```go
func (r *Request) ContentType(ct string) *Request
```

设置 Content-Type 头。

### Accept

```go
func (r *Request) Accept(ct string) *Request
```

设置 Accept 头。

### UserAgent

```go
func (r *Request) UserAgent(ua string) *Request
```

设置 User-Agent 头。

### Referer

```go
func (r *Request) Referer(referer string) *Request
```

设置 Referer 头。

### BasicAuth

```go
func (r *Request) BasicAuth(username, password string) *Request
```

设置请求级 Basic 认证。

### BearerToken

```go
func (r *Request) BearerToken(token string) *Request
```

设置请求级 Bearer Token。

### Do

```go
func (r *Request) Do() (*Response, error)
```

执行请求并返回响应。

**示例:**

```go
resp, err := httpclient.New().
    Get("https://api.example.com/users").
    Header("Accept", "application/json").
    Query("page", "1").
    Do()
if err != nil {
    log.Fatal(err)
}
defer resp.Close()
```

## Response

### Status / StatusCode

```go
func (r *Response) Status() int
func (r *Response) StatusCode() int
```

获取 HTTP 状态码。

**示例:**

```go
fmt.Println(resp.StatusCode()) // 200
```

### IsOK

```go
func (r *Response) IsOK() bool
```

判断状态码是否为 200。

### IsSuccess

```go
func (r *Response) IsSuccess() bool
```

判断状态码是否为 2xx。

### IsError

```go
func (r *Response) IsError() bool
```

判断状态码是否为 4xx 或 5xx。

### Header

```go
func (r *Response) Header() http.Header
```

获取所有响应头。

### GetHeader

```go
func (r *Response) GetHeader(key string) string
```

获取指定响应头的值。

### ContentType

```go
func (r *Response) ContentType() string
```

获取 Content-Type 头。

### ContentLength

```go
func (r *Response) ContentLength() int64
```

获取内容长度。

### Body

```go
func (r *Response) Body() ([]byte, error)
```

获取响应体字节数据。

**示例:**

```go
data, err := resp.Body()
```

### BodyString

```go
func (r *Response) BodyString() (string, error)
```

获取响应体字符串。

**示例:**

```go
body, err := resp.BodyString()
fmt.Println(body)
```

### Cookies

```go
func (r *Response) Cookies() []*http.Cookie
```

获取响应中的所有 Cookie。

### Raw

```go
func (r *Response) Raw() *http.Response
```

获取原始的 http.Response 对象。

### Close

```go
func (r *Response) Close() error
```

关闭响应体。

**示例:**

```go
resp, err := httpclient.Get("https://api.example.com/users")
if err != nil {
    log.Fatal(err)
}
defer resp.Close()
```

## 完整示例

```go
package main

import (
    "fmt"
    "log"

    "github.com/gongm/gohutool/httpclient"
)

func main() {
    // 创建客户端并设置认证
    client := httpclient.New().
        SetBearerToken("your-token").
        SetHeader("Accept", "application/json")

    // GET 请求带查询参数
    resp, err := client.Get("https://api.example.com/users").
        Query("page", "1").
        Query("size", "10").
        Do()
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Close()

    if resp.IsSuccess() {
        body, _ := resp.BodyString()
        fmt.Println(body)
    }

    // POST JSON 请求
    resp, err = client.Post("https://api.example.com/users").
        JSON(map[string]interface{}{
            "name":  "John",
            "email": "john@example.com",
        }).
        Do()
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Close()

    fmt.Println(resp.StatusCode())
}
```
