package httpclient

import (
	"net/http"
)

// Response is HTTP response wrapper
type Response struct {
	resp *http.Response
	body []byte
}

// Status returns HTTP status code
func (r *Response) Status() int {
	return r.resp.StatusCode
}

// StatusCode returns HTTP status code
func (r *Response) StatusCode() int {
	return r.resp.StatusCode
}

// IsOK checks if status is 200
func (r *Response) IsOK() bool {
	return r.resp.StatusCode == http.StatusOK
}

// IsSuccess checks if status is 2xx
func (r *Response) IsSuccess() bool {
	return r.resp.StatusCode >= 200 && r.resp.StatusCode < 300
}

// IsError checks if status is 4xx or 5xx
func (r *Response) IsError() bool {
	return r.resp.StatusCode >= 400
}

// Header returns response headers
func (r *Response) Header() http.Header {
	return r.resp.Header
}

// GetHeader gets header value
func (r *Response) GetHeader(key string) string {
	return r.resp.Header.Get(key)
}

// ContentType returns Content-Type
func (r *Response) ContentType() string {
	return r.resp.Header.Get("Content-Type")
}

// ContentLength returns content length
func (r *Response) ContentLength() int64 {
	return r.resp.ContentLength
}

// Body returns response body bytes
func (r *Response) Body() ([]byte, error) {
	return r.body, nil
}

// BodyString returns response body as string
func (r *Response) BodyString() (string, error) {
	return string(r.body), nil
}

// Cookies returns response cookies
func (r *Response) Cookies() []*http.Cookie {
	return r.resp.Cookies()
}

// Raw returns raw http.Response
func (r *Response) Raw() *http.Response {
	return r.resp
}

// Close closes response body
func (r *Response) Close() error {
	return r.resp.Body.Close()
}
