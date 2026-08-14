package httpclient

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	c := New()
	if c == nil {
		t.Fatal("New() returned nil")
	}
	if c.client == nil {
		t.Fatal("client is nil")
	}
	if c.headers == nil {
		t.Fatal("headers is nil")
	}
}

func TestNewWithTimeout(t *testing.T) {
	timeout := 10 * time.Second
	c := NewWithTimeout(timeout)
	if c == nil {
		t.Fatal("NewWithTimeout() returned nil")
	}
	if c.client.Timeout != timeout {
		t.Fatalf("expected timeout %v, got %v", timeout, c.client.Timeout)
	}
}

func TestSetTimeout(t *testing.T) {
	c := New()
	timeout := 5 * time.Second
	c.SetTimeout(timeout)
	if c.client.Timeout != timeout {
		t.Fatalf("expected timeout %v, got %v", timeout, c.client.Timeout)
	}
}

func TestSetHeader(t *testing.T) {
	c := New()
	c.SetHeader("X-Custom", "value")
	if c.headers["X-Custom"] != "value" {
		t.Fatalf("expected header 'value', got '%s'", c.headers["X-Custom"])
	}
}

func TestSetHeaders(t *testing.T) {
	c := New()
	c.SetHeaders(map[string]string{
		"X-Custom1": "value1",
		"X-Custom2": "value2",
	})
	if c.headers["X-Custom1"] != "value1" {
		t.Fatalf("expected header 'value1', got '%s'", c.headers["X-Custom1"])
	}
	if c.headers["X-Custom2"] != "value2" {
		t.Fatalf("expected header 'value2', got '%s'", c.headers["X-Custom2"])
	}
}

func TestSetCookie(t *testing.T) {
	c := New()
	cookie := &http.Cookie{Name: "test", Value: "value"}
	c.SetCookie(cookie)
	if len(c.cookies) != 1 {
		t.Fatalf("expected 1 cookie, got %d", len(c.cookies))
	}
	if c.cookies[0].Name != "test" {
		t.Fatalf("expected cookie name 'test', got '%s'", c.cookies[0].Name)
	}
}

func TestSetBasicAuth(t *testing.T) {
	c := New()
	c.SetBasicAuth("user", "pass")
	if !strings.HasPrefix(c.headers["Authorization"], "Basic ") {
		t.Fatal("expected Basic auth header")
	}
}

func TestSetBearerToken(t *testing.T) {
	c := New()
	c.SetBearerToken("token123")
	if c.headers["Authorization"] != "Bearer token123" {
		t.Fatalf("expected 'Bearer token123', got '%s'", c.headers["Authorization"])
	}
}

func TestGetRequest(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("expected GET method, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	resp, err := Get(server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !resp.IsOK() {
		t.Fatalf("expected status 200, got %d", resp.StatusCode())
	}
	body, _ := resp.BodyString()
	if body != "OK" {
		t.Fatalf("expected body 'OK', got '%s'", body)
	}
}

func TestPostJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST method, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json; charset=utf-8" {
			t.Fatalf("expected Content-Type 'application/json; charset=utf-8', got '%s'", r.Header.Get("Content-Type"))
		}
		body, _ := io.ReadAll(r.Body)
		var data map[string]string
		json.Unmarshal(body, &data)
		if data["key"] != "value" {
			t.Fatalf("expected key=value, got %s", data["key"])
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"status":"created"}`))
	}))
	defer server.Close()

	resp, err := PostJSON(server.URL, map[string]string{"key": "value"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode() != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", resp.StatusCode())
	}
}

func TestPostForm(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST method, got %s", r.Method)
		}
		if !strings.Contains(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
			t.Fatalf("expected form Content-Type, got '%s'", r.Header.Get("Content-Type"))
		}
		r.ParseForm()
		if r.FormValue("name") != "test" {
			t.Fatalf("expected name=test, got %s", r.FormValue("name"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	resp, err := PostForm(server.URL, map[string]string{"name": "test"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !resp.IsSuccess() {
		t.Fatalf("expected success status, got %d", resp.StatusCode())
	}
}

func TestPutRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("expected PUT method, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	resp, err := Put(server.URL, "data")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !resp.IsSuccess() {
		t.Fatalf("expected success status, got %d", resp.StatusCode())
	}
}

func TestDeleteRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("expected DELETE method, got %s", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	resp, err := Delete(server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode() != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", resp.StatusCode())
	}
}

func TestClientBuilder(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Custom") != "value" {
			t.Fatalf("expected X-Custom header 'value', got '%s'", r.Header.Get("X-Custom"))
		}
		if r.URL.Query().Get("page") != "1" {
			t.Fatalf("expected query param page=1, got %s", r.URL.Query().Get("page"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := New()
	c.SetHeader("X-Custom", "value")

	resp, err := c.Get(server.URL).Query("page", "1").Do()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !resp.IsSuccess() {
		t.Fatalf("expected success status, got %d", resp.StatusCode())
	}
}

func TestRequestHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Accept") != "application/json" {
			t.Fatalf("expected Accept header 'application/json', got '%s'", r.Header.Get("Accept"))
		}
		if r.Header.Get("User-Agent") != "GoHutool/1.0" {
			t.Fatalf("expected User-Agent 'GoHutool/1.0', got '%s'", r.Header.Get("User-Agent"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := New()
	resp, err := c.Get(server.URL).
		Accept("application/json").
		UserAgent("GoHutool/1.0").
		Do()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !resp.IsSuccess() {
		t.Fatalf("expected success status, got %d", resp.StatusCode())
	}
}

func TestRequestCookies(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			t.Fatalf("expected cookie 'session', got error: %v", err)
		}
		if cookie.Value != "abc123" {
			t.Fatalf("expected cookie value 'abc123', got '%s'", cookie.Value)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := New()
	resp, err := c.Get(server.URL).
		Cookie(&http.Cookie{Name: "session", Value: "abc123"}).
		Do()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !resp.IsSuccess() {
		t.Fatalf("expected success status, got %d", resp.StatusCode())
	}
}

func TestRequestBasicAuth(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok {
			t.Fatal("expected basic auth")
		}
		if user != "admin" || pass != "secret" {
			t.Fatalf("expected admin:secret, got %s:%s", user, pass)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := New()
	resp, err := c.Get(server.URL).
		BasicAuth("admin", "secret").
		Do()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !resp.IsSuccess() {
		t.Fatalf("expected success status, got %d", resp.StatusCode())
	}
}

func TestRequestBearerToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != "Bearer mytoken" {
			t.Fatalf("expected 'Bearer mytoken', got '%s'", auth)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := New()
	resp, err := c.Get(server.URL).
		BearerToken("mytoken").
		Do()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !resp.IsSuccess() {
		t.Fatalf("expected success status, got %d", resp.StatusCode())
	}
}

func TestResponseMethods(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("X-Custom", "custom-value")
		http.SetCookie(w, &http.Cookie{Name: "test", Value: "value"})
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}))
	defer server.Close()

	resp, err := Get(server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Close()

	if !resp.IsOK() {
		t.Fatal("expected IsOK() to be true")
	}
	if !resp.IsSuccess() {
		t.Fatal("expected IsSuccess() to be true")
	}
	if resp.IsError() {
		t.Fatal("expected IsError() to be false")
	}
	if resp.Status() != 200 {
		t.Fatalf("expected status 200, got %d", resp.Status())
	}
	if resp.ContentType() != "text/plain" {
		t.Fatalf("expected Content-Type 'text/plain', got '%s'", resp.ContentType())
	}
	if resp.GetHeader("X-Custom") != "custom-value" {
		t.Fatalf("expected X-Custom 'custom-value', got '%s'", resp.GetHeader("X-Custom"))
	}

	body, err := resp.BodyString()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if body != "Hello, World!" {
		t.Fatalf("expected body 'Hello, World!', got '%s'", body)
	}

	bodyBytes, err := resp.Body()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(bodyBytes) != "Hello, World!" {
		t.Fatalf("expected body bytes 'Hello, World!', got '%s'", string(bodyBytes))
	}

	cookies := resp.Cookies()
	if len(cookies) != 1 {
		t.Fatalf("expected 1 cookie, got %d", len(cookies))
	}
	if cookies[0].Name != "test" {
		t.Fatalf("expected cookie name 'test', got '%s'", cookies[0].Name)
	}

	raw := resp.Raw()
	if raw == nil {
		t.Fatal("expected Raw() to not be nil")
	}
}

func TestResponseError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}))
	defer server.Close()

	resp, err := Get(server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.IsOK() {
		t.Fatal("expected IsOK() to be false")
	}
	if resp.IsSuccess() {
		t.Fatal("expected IsSuccess() to be false")
	}
	if !resp.IsError() {
		t.Fatal("expected IsError() to be true")
	}
	if resp.StatusCode() != 500 {
		t.Fatalf("expected status 500, got %d", resp.StatusCode())
	}
}

func TestClientWithCookies(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("client-cookie")
		if err != nil {
			t.Fatalf("expected client cookie, got error: %v", err)
		}
		if cookie.Value != "client-value" {
			t.Fatalf("expected client cookie value 'client-value', got '%s'", cookie.Value)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := New()
	c.SetCookie(&http.Cookie{Name: "client-cookie", Value: "client-value"})

	resp, err := c.Get(server.URL).Do()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !resp.IsSuccess() {
		t.Fatalf("expected success status, got %d", resp.StatusCode())
	}
}

func TestRequestBodyReader(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		if string(body) != "reader body" {
			t.Fatalf("expected body 'reader body', got '%s'", string(body))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := New()
	resp, err := c.Post(server.URL).
		BodyReader(strings.NewReader("reader body")).
		Do()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !resp.IsSuccess() {
		t.Fatalf("expected success status, got %d", resp.StatusCode())
	}
}

func TestRequestBodyBytes(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		if string(body) != "bytes body" {
			t.Fatalf("expected body 'bytes body', got '%s'", string(body))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := New()
	resp, err := c.Post(server.URL).
		BodyBytes([]byte("bytes body")).
		Do()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !resp.IsSuccess() {
		t.Fatalf("expected success status, got %d", resp.StatusCode())
	}
}

func TestRequestFormValues(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if r.FormValue("key") != "value" {
			t.Fatalf("expected key=value, got %s", r.FormValue("key"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := New()
	values := map[string]string{"key": "value"}
	resp, err := c.Post(server.URL).
		Form(values).
		Do()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !resp.IsSuccess() {
		t.Fatalf("expected success status, got %d", resp.StatusCode())
	}
}

func TestRequestReferer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Referer") != "http://example.com" {
			t.Fatalf("expected Referer 'http://example.com', got '%s'", r.Header.Get("Referer"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := New()
	resp, err := c.Get(server.URL).
		Referer("http://example.com").
		Do()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !resp.IsSuccess() {
		t.Fatalf("expected success status, got %d", resp.StatusCode())
	}
}

func TestRequestTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := New()
	_, err := c.Get(server.URL).
		Timeout(100 * time.Millisecond).
		Do()
	if err == nil {
		t.Fatal("expected timeout error")
	}
}

func TestMultipleQueries(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("page") != "1" {
			t.Fatalf("expected page=1, got %s", r.URL.Query().Get("page"))
		}
		if r.URL.Query().Get("size") != "10" {
			t.Fatalf("expected size=10, got %s", r.URL.Query().Get("size"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := New()
	resp, err := c.Get(server.URL).
		Queries(map[string]string{"page": "1", "size": "10"}).
		Do()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !resp.IsSuccess() {
		t.Fatalf("expected success status, got %d", resp.StatusCode())
	}
}

func TestClientOverrideHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Custom") != "overridden" {
			t.Fatalf("expected X-Custom 'overridden', got '%s'", r.Header.Get("X-Custom"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := New()
	c.SetHeader("X-Custom", "default")

	resp, err := c.Get(server.URL).
		Header("X-Custom", "overridden").
		Do()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !resp.IsSuccess() {
		t.Fatalf("expected success status, got %d", resp.StatusCode())
	}
}

func TestContentLength(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "5")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello"))
	}))
	defer server.Close()

	resp, err := Get(server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.ContentLength() != 5 {
		t.Fatalf("expected ContentLength 5, got %d", resp.ContentLength())
	}
}
