package httpclient

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// content is the canonical payload for the resume tests (20 bytes).
const resumeData = "0123456789abcdefghij"

// resumeServer serves resumeData and honors Range requests with a 206.
// It records every Range header it receives.
func resumeServer(t *testing.T) (*httptest.Server, *[]string) {
	var ranges []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rng := r.Header.Get("Range")
		ranges = append(ranges, rng)
		if rng != "" {
			var start int
			if _, err := fmt.Sscanf(rng, "bytes=%d-", &start); err != nil || start < 0 || start > len(resumeData) {
				http.Error(w, "bad range", http.StatusRequestedRangeNotSatisfiable)
				return
			}
			w.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, len(resumeData)-1, len(resumeData)))
			w.WriteHeader(http.StatusPartialContent)
			_, _ = w.Write([]byte(resumeData[start:]))
			return
		}
		w.Write([]byte(resumeData))
	}))
	return srv, &ranges
}

func TestResumeDownloadToFileFullDownload(t *testing.T) {
	srv, ranges := resumeServer(t)
	defer srv.Close()

	path := filepath.Join(t.TempDir(), "resume.txt")
	total, err := ResumeDownloadToFile(srv.URL+"/f", path)
	if err != nil {
		t.Fatal(err)
	}
	if total != int64(len(resumeData)) {
		t.Errorf("total = %d, want %d", total, len(resumeData))
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != resumeData {
		t.Errorf("file = %q", got)
	}
	// first request must not carry a Range header
	if len(*ranges) != 1 || (*ranges)[0] != "" {
		t.Errorf("first request Range = %v", *ranges)
	}
}

func TestResumeDownloadToFileResumes(t *testing.T) {
	srv, ranges := resumeServer(t)
	defer srv.Close()

	dir := t.TempDir()
	path := filepath.Join(dir, "resume.txt")
	// simulate a partially downloaded file of 15 bytes
	if err := os.WriteFile(path, []byte(resumeData[:15]), 0o644); err != nil {
		t.Fatal(err)
	}

	total, err := ResumeDownloadToFile(srv.URL+"/f", path)
	if err != nil {
		t.Fatal(err)
	}
	if total != int64(len(resumeData)) {
		t.Errorf("total = %d, want %d", total, len(resumeData))
	}
	got, _ := os.ReadFile(path)
	if string(got) != resumeData {
		t.Errorf("final file = %q", got)
	}
	// request must carry Range: bytes=15-
	if len(*ranges) != 1 || (*ranges)[0] != "bytes=15-" {
		t.Errorf("resume request Range = %v", *ranges)
	}
}

func TestResumeDownloadToFileServerIgnoresRange(t *testing.T) {
	// a server that answers 200 (full body) even when Range is sent
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = r.Header.Get("Range") // present but ignored
		w.Write([]byte(resumeData))
	}))
	defer srv.Close()

	path := filepath.Join(t.TempDir(), "partial.bin")
	if err := os.WriteFile(path, []byte("GARBAGE-old-partial"), 0o644); err != nil {
		t.Fatal(err)
	}
	total, err := ResumeDownloadToFile(srv.URL, path)
	if err != nil {
		t.Fatal(err)
	}
	if total != int64(len(resumeData)) {
		t.Errorf("total = %d, want %d", total, len(resumeData))
	}
	// 200 means full rewrite: old bytes must be gone
	got, _ := os.ReadFile(path)
	if string(got) != resumeData {
		t.Errorf("file after 200 rewrite = %q", got)
	}
}

func TestResumeDownloadToFileCoversNetworkErrors(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "gone", http.StatusNotFound)
	}))
	defer srv.Close()

	path := filepath.Join(t.TempDir(), "x")
	if _, err := ResumeDownloadToFile(srv.URL, path); err == nil {
		t.Error("4xx response should error")
	} else if !strings.Contains(err.Error(), "404") {
		t.Errorf("error should mention status, got %v", err)
	}
}