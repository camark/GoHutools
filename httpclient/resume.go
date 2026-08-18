package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// ResumeDownloadToFile downloads rawURL into filePath, resuming an
// existing partial file: when filePath already exists with size > 0 it
// sends "Range: bytes=<size>-" and appends the 206 remainder. If the
// server responds 200 (ignoring Range), the file is fully rewritten
// from scratch. Parent directories are created as needed. Returns the
// final file size.
//
// It streams with the net/http default client (like DownloadToFile)
// because the Response wrapper reads bodies into memory; resuming only
// makes sense for large files.
func ResumeDownloadToFile(rawURL, filePath string) (int64, error) {
	var start int64
	if st, err := os.Stat(filePath); err == nil {
		start = st.Size()
	}

	req, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		return 0, err
	}
	if start > 0 {
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-", start))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode >= http.StatusBadRequest {
		return 0, fmt.Errorf("download failed: %s", resp.Status)
	}

	if dir := filepath.Dir(filePath); dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return 0, err
		}
	}

	resumes := start > 0 && resp.StatusCode == http.StatusPartialContent
	var f *os.File
	if resumes {
		f, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0o644)
	} else {
		start = 0
		f, err = os.Create(filePath)
	}
	if err != nil {
		return 0, err
	}
	defer func() { _ = f.Close() }()

	n, err := io.Copy(f, resp.Body)
	return start + n, err
}