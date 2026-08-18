package httpclient

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestDownloadBytes(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	}))
	defer srv.Close()

	resp, err := DownloadBytes(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	if string(resp) != "hello world" {
		t.Errorf("DownloadBytes = %q", resp)
	}
}

func TestDownloadToFile(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/dir/nested/data.txt" {
			w.Write([]byte("file data"))
			return
		}
		http.NotFound(w, r)
	}))
	defer srv.Close()

	dir := t.TempDir()
	target := filepath.Join(dir, "out", "nested", "data.txt")
	if err := DownloadToFile(srv.URL+"/dir/nested/data.txt", target); err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(target)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "file data" {
		t.Errorf("file content = %q", got)
	}
}

func TestUploadFile(t *testing.T) {
	var gotName, gotContent string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(1 << 20)
		f, header, err := r.FormFile("file")
		if err != nil {
			t.Errorf("FormFile: %v", err)
			http.Error(w, "no file", 400)
			return
		}
		defer f.Close()
		gotName = header.Filename
		b, _ := io.ReadAll(f)
		gotContent = string(b)
		w.WriteHeader(200)
	}))
	defer srv.Close()

	dir := t.TempDir()
	path := filepath.Join(dir, "upload.txt")
	if err := os.WriteFile(path, []byte("upload-content"), 0o644); err != nil {
		t.Fatal(err)
	}
	resp, err := UploadFile(srv.URL, "file", path)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Close()
	if !resp.IsOK() {
		t.Error("upload should succeed")
	}
	if gotName != "upload.txt" || gotContent != "upload-content" {
		t.Errorf("upload got name=%q content=%q", gotName, gotContent)
	}
}

func TestUploadFileBytes(t *testing.T) {
	got := ""
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// read raw multipart body length; verify part presence via content-type boundary
		ct := r.Header.Get("Content-Type")
		if !strings.HasPrefix(ct, "multipart/form-data; boundary=") {
			t.Errorf("content-type = %q", ct)
		}
		_ = r.ParseMultipartForm(1 << 20)
		f, _, err := r.FormFile("data")
		if err == nil {
			b, _ := io.ReadAll(f)
			got = string(b)
		}
		w.WriteHeader(200)
	}))
	defer srv.Close()

	resp, err := UploadFileBytes(srv.URL, "data", "a.bin", []byte("bytes-42"))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Close()
	if got != "bytes-42" {
		t.Errorf("uploaded bytes = %q", got)
	}
}

func TestClientCookieJarKeepsSession(t *testing.T) {
	var firstHadCookie bool
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := r.Cookie("session"); err == nil {
			firstHadCookie = true
		}
		http.SetCookie(w, &http.Cookie{Name: "session", Value: "abc123"})
		w.Write([]byte("ok"))
	}))
	defer srv.Close()

	c := New().EnableCookieJar()
	r1, err := c.Get(srv.URL).Do()
	if err != nil {
		t.Fatal(err)
	}
	r1.Close()
	if firstHadCookie {
		t.Error("first request should have no cookie")
	}
	r2, err := c.Get(srv.URL).Do()
	if err != nil {
		t.Fatal(err)
	}
	r2.Close()
	if !firstHadCookie {
		t.Error("second request should carry the session cookie from the jar")
	}
}

func TestClientSetUserAgent(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-UA", r.UserAgent())
		w.Write([]byte("ok"))
	}))
	defer srv.Close()

	resp, err := New().SetUserAgent("GoHutool-Test").Get(srv.URL).Do()
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Close()
	if got := resp.GetHeader("X-UA"); got != "GoHutool-Test" {
		t.Errorf("server saw UA = %q", got)
	}
}

func TestRequestTimeoutStillHasJar(t *testing.T) {
	// when a per-request timeout creates a fresh http.Client, the
	// cookie jar from the Client must survive.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "sid", Value: "v"})
		w.Write([]byte("ok"))
	}))
	defer srv.Close()

	c := New().EnableCookieJar()
	r1, err := c.Get(srv.URL).Do()
	if err != nil {
		t.Fatal(err)
	}
	r1.Close()
	r2, err := c.Get(srv.URL).Timeout(3 * time.Second).Do() // timeout path (0 means client default)
	if err != nil {
		t.Fatal(err)
	}
	r2.Close()
	// assert server saw no error and response ok; jar behavior verified by
	// absence of panic + successful second request. Use Response.Cookies nil check
	if !r2.IsOK() {
		t.Error("second request failed")
	}
	// now check jar actually forwarded via a third request that asserts cookie presence
	var saw bool
	srv2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("sid")
		saw = err == nil
		w.Write([]byte("ok"))
	}))
	defer srv2.Close()
	_, err = c.Get(srv2.URL).Timeout(3 * time.Second).Do()
	if err != nil {
		t.Fatal(err)
	}
	if !saw {
		t.Error("timeout-path request should carry jar cookie")
	}
}