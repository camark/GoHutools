package httpclient

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

// Client is HTTP client wrapper
type Client struct {
	client  *http.Client
	headers map[string]string
	cookies []*http.Cookie
}

// New creates new client
func New() *Client {
	return &Client{
		client:  &http.Client{},
		headers: make(map[string]string),
	}
}

// NewWithTimeout creates new client with timeout
func NewWithTimeout(timeout time.Duration) *Client {
	return &Client{
		client:  &http.Client{Timeout: timeout},
		headers: make(map[string]string),
	}
}

// SetTimeout sets request timeout
func (c *Client) SetTimeout(d time.Duration) *Client {
	c.client.Timeout = d
	return c
}

// SetHeader sets default header
func (c *Client) SetHeader(key, value string) *Client {
	c.headers[key] = value
	return c
}

// SetHeaders sets multiple default headers
func (c *Client) SetHeaders(headers map[string]string) *Client {
	for k, v := range headers {
		c.headers[k] = v
	}
	return c
}

// SetCookie sets cookie
func (c *Client) SetCookie(cookie *http.Cookie) *Client {
	c.cookies = append(c.cookies, cookie)
	return c
}

// EnableCookieJar turns on automatic cookie storage and reuse
// (session support) for this client.
func (c *Client) EnableCookieJar() *Client {
	jar, err := cookiejar.New(nil)
	if err == nil {
		c.client.Jar = jar
	}
	return c
}

// SetUserAgent sets a default User-Agent header for all requests.
func (c *Client) SetUserAgent(ua string) *Client {
	c.headers["User-Agent"] = ua
	return c
}

// SetBasicAuth sets basic auth
func (c *Client) SetBasicAuth(username, password string) *Client {
	c.headers["Authorization"] = "Basic " + basicAuth(username, password)
	return c
}

// SetBearerToken sets bearer token
func (c *Client) SetBearerToken(token string) *Client {
	c.headers["Authorization"] = "Bearer " + token
	return c
}

// Get creates GET request
func (c *Client) Get(rawURL string) *Request {
	return &Request{
		client:  c,
		method:  http.MethodGet,
		url:     rawURL,
		headers: make(map[string]string),
		query:   url.Values{},
	}
}

// Post creates POST request
func (c *Client) Post(rawURL string) *Request {
	return &Request{
		client:  c,
		method:  http.MethodPost,
		url:     rawURL,
		headers: make(map[string]string),
		query:   url.Values{},
	}
}

// Put creates PUT request
func (c *Client) Put(rawURL string) *Request {
	return &Request{
		client:  c,
		method:  http.MethodPut,
		url:     rawURL,
		headers: make(map[string]string),
		query:   url.Values{},
	}
}

// Delete creates DELETE request
func (c *Client) Delete(rawURL string) *Request {
	return &Request{
		client:  c,
		method:  http.MethodDelete,
		url:     rawURL,
		headers: make(map[string]string),
		query:   url.Values{},
	}
}

// Patch creates PATCH request
func (c *Client) Patch(rawURL string) *Request {
	return &Request{
		client:  c,
		method:  http.MethodPatch,
		url:     rawURL,
		headers: make(map[string]string),
		query:   url.Values{},
	}
}

// Head creates HEAD request
func (c *Client) Head(rawURL string) *Request {
	return &Request{
		client:  c,
		method:  http.MethodHead,
		url:     rawURL,
		headers: make(map[string]string),
		query:   url.Values{},
	}
}

// Options creates OPTIONS request
func (c *Client) Options(rawURL string) *Request {
	return &Request{
		client:  c,
		method:  http.MethodOptions,
		url:     rawURL,
		headers: make(map[string]string),
		query:   url.Values{},
	}
}

// defaultClient is the default client for quick functions
var defaultClient = New()

// Get sends GET request
func Get(rawURL string) (*Response, error) {
	return defaultClient.Get(rawURL).Do()
}

// Post sends POST request
func Post(rawURL string, body interface{}) (*Response, error) {
	return defaultClient.Post(rawURL).Body(body).Do()
}

// PostForm sends POST request with form data
func PostForm(rawURL string, data map[string]string) (*Response, error) {
	return defaultClient.Post(rawURL).Form(data).Do()
}

// PostJSON sends POST request with JSON body
func PostJSON(rawURL string, data interface{}) (*Response, error) {
	return defaultClient.Post(rawURL).JSON(data).Do()
}

// Put sends PUT request
func Put(rawURL string, body interface{}) (*Response, error) {
	return defaultClient.Put(rawURL).Body(body).Do()
}

// Delete sends DELETE request
func Delete(rawURL string) (*Response, error) {
	return defaultClient.Delete(rawURL).Do()
}

// basicAuth encodes username and password for Basic auth
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64Encode([]byte(auth))
}

// base64Encode encodes bytes to base64 string
func base64Encode(data []byte) string {
	const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	const padChar = '='

	buf := make([]byte, 0, (len(data)/3+1)*4)

	for i := 0; i < len(data); i += 3 {
		var b0, b1, b2 byte
		b0 = data[i]
		if i+1 < len(data) {
			b1 = data[i+1]
		}
		if i+2 < len(data) {
			b2 = data[i+2]
		}

		buf = append(buf, encodeStd[(b0>>2)&0x3F])
		buf = append(buf, encodeStd[((b0<<4)|(b1>>4))&0x3F])

		if i+1 < len(data) {
			buf = append(buf, encodeStd[((b1<<2)|(b2>>6))&0x3F])
		} else {
			buf = append(buf, byte(padChar))
		}

		if i+2 < len(data) {
			buf = append(buf, encodeStd[b2&0x3F])
		} else {
			buf = append(buf, byte(padChar))
		}
	}

	return string(buf)
}

// buildURL builds URL with query parameters
func buildURL(rawURL string, query url.Values) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	if len(query) > 0 {
		existingQuery := u.Query()
		for k, vs := range query {
			for _, v := range vs {
				existingQuery.Add(k, v)
			}
		}
		u.RawQuery = existingQuery.Encode()
	}

	return u.String(), nil
}
