package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Request is HTTP request builder
type Request struct {
	client  *Client
	method  string
	url     string
	headers map[string]string
	cookies []*http.Cookie
	body    io.Reader
	query   url.Values
	timeout time.Duration
}

// Header sets request header
func (r *Request) Header(key, value string) *Request {
	r.headers[key] = value
	return r
}

// Headers sets multiple headers
func (r *Request) Headers(headers map[string]string) *Request {
	for k, v := range headers {
		r.headers[k] = v
	}
	return r
}

// Query adds query parameter
func (r *Request) Query(key, value string) *Request {
	r.query.Add(key, value)
	return r
}

// Queries adds multiple query parameters
func (r *Request) Queries(params map[string]string) *Request {
	for k, v := range params {
		r.query.Add(k, v)
	}
	return r
}

// Body sets request body
func (r *Request) Body(body interface{}) *Request {
	if body == nil {
		return r
	}

	switch v := body.(type) {
	case string:
		r.body = strings.NewReader(v)
	case []byte:
		r.body = bytes.NewReader(v)
	case io.Reader:
		r.body = v
	default:
		// Try JSON marshal for other types
		data, err := json.Marshal(v)
		if err != nil {
			r.body = strings.NewReader(fmt.Sprintf("%v", v))
		} else {
			r.body = bytes.NewReader(data)
			if _, ok := r.headers["Content-Type"]; !ok {
				r.headers["Content-Type"] = "application/json; charset=utf-8"
			}
		}
	}
	return r
}

// BodyString sets string body
func (r *Request) BodyString(body string) *Request {
	r.body = strings.NewReader(body)
	return r
}

// BodyBytes sets bytes body
func (r *Request) BodyBytes(body []byte) *Request {
	r.body = bytes.NewReader(body)
	return r
}

// BodyReader sets io.Reader body
func (r *Request) BodyReader(body io.Reader) *Request {
	r.body = body
	return r
}

// Form sets form data
func (r *Request) Form(data map[string]string) *Request {
	values := url.Values{}
	for k, v := range data {
		values.Set(k, v)
	}
	r.body = strings.NewReader(values.Encode())
	if _, ok := r.headers["Content-Type"]; !ok {
		r.headers["Content-Type"] = "application/x-www-form-urlencoded"
	}
	return r
}

// FormValues sets url.Values form data
func (r *Request) FormValues(data url.Values) *Request {
	r.body = strings.NewReader(data.Encode())
	if _, ok := r.headers["Content-Type"]; !ok {
		r.headers["Content-Type"] = "application/x-www-form-urlencoded"
	}
	return r
}

// JSON sets JSON body
func (r *Request) JSON(data interface{}) *Request {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		// Store error in body for later handling
		r.body = strings.NewReader(fmt.Sprintf("json marshal error: %v", err))
		return r
	}
	r.body = bytes.NewReader(jsonBytes)
	if _, ok := r.headers["Content-Type"]; !ok {
		r.headers["Content-Type"] = "application/json; charset=utf-8"
	}
	return r
}

// Cookie adds cookie
func (r *Request) Cookie(cookie *http.Cookie) *Request {
	r.cookies = append(r.cookies, cookie)
	return r
}

// Timeout sets request timeout
func (r *Request) Timeout(d time.Duration) *Request {
	r.timeout = d
	return r
}

// ContentType sets Content-Type header
func (r *Request) ContentType(ct string) *Request {
	r.headers["Content-Type"] = ct
	return r
}

// Accept sets Accept header
func (r *Request) Accept(ct string) *Request {
	r.headers["Accept"] = ct
	return r
}

// UserAgent sets User-Agent header
func (r *Request) UserAgent(ua string) *Request {
	r.headers["User-Agent"] = ua
	return r
}

// Referer sets Referer header
func (r *Request) Referer(referer string) *Request {
	r.headers["Referer"] = referer
	return r
}

// BasicAuth sets basic auth
func (r *Request) BasicAuth(username, password string) *Request {
	r.headers["Authorization"] = "Basic " + basicAuth(username, password)
	return r
}

// BearerToken sets bearer token
func (r *Request) BearerToken(token string) *Request {
	r.headers["Authorization"] = "Bearer " + token
	return r
}

// Do executes request
func (r *Request) Do() (*Response, error) {
	// Build URL with query parameters
	reqURL, err := buildURL(r.url, r.query)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	// Create request
	req, err := http.NewRequest(r.method, reqURL, r.body)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	// Apply client default headers
	for k, v := range r.client.headers {
		if _, ok := r.headers[k]; !ok {
			req.Header.Set(k, v)
		}
	}

	// Apply request headers (override client defaults)
	for k, v := range r.headers {
		req.Header.Set(k, v)
	}

	// Apply client cookies
	for _, cookie := range r.client.cookies {
		req.AddCookie(cookie)
	}

	// Apply request cookies
	for _, cookie := range r.cookies {
		req.AddCookie(cookie)
	}

	// Handle timeout - use request timeout if set, otherwise use client timeout
	client := r.client.client
	if r.timeout > 0 {
		// preserve the cookie jar (and any redirect policy) when a
		// per-request timeout forces a fresh http.Client
		client = &http.Client{
			Timeout:       r.timeout,
			Jar:           r.client.client.Jar,
			CheckRedirect: r.client.client.CheckRedirect,
		}
	}

	// Execute request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return nil, fmt.Errorf("read response body failed: %w", err)
	}

	return &Response{
		resp: resp,
		body: bodyBytes,
	}, nil
}
