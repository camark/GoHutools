package httpclient

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// This file ports high-value convenience functions from Hutool's
// HttpUtil that sit on top of the existing Client/Request API.

// DownloadBytes performs a GET and returns the raw response body.
func DownloadBytes(rawURL string) ([]byte, error) {
	resp, err := Get(rawURL)
	if err != nil {
		return nil, err
	}
	defer resp.Close()
	if !resp.IsOK() {
		return nil, fmt.Errorf("download failed: %s", resp.Raw().Status)
	}
	return resp.Body()
}

// DownloadToFile downloads rawURL and writes the body to filePath,
// streaming the content and creating parent directories as needed.
func DownloadToFile(rawURL, filePath string) error {
	resp, err := http.Get(rawURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("download failed: %s", resp.Status)
	}
	if dir := filepath.Dir(filePath); dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

// UploadFile uploads the file at filePath as a multipart form field.
func UploadFile(rawURL, field, filePath string) (*Response, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return UploadFileReader(rawURL, field, filepath.Base(filePath), f)
}

// UploadFileBytes uploads in-memory content as a multipart form field.
func UploadFileBytes(rawURL, field, filename string, content []byte) (*Response, error) {
	return UploadFileReader(rawURL, field, filename, bytes.NewReader(content))
}

// UploadFileReader uploads arbitrary content from an io.Reader as a
// multipart form field.
func UploadFileReader(rawURL, field, filename string, r io.Reader) (*Response, error) {
	var buf bytes.Buffer
	mw := multipart.NewWriter(&buf)
	fw, err := mw.CreateFormFile(field, filename)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(fw, r); err != nil {
		return nil, err
	}
	if err := mw.Close(); err != nil {
		return nil, err
	}
	return New().
		Post(rawURL).
		BodyReader(&buf).
		ContentType(mw.FormDataContentType()).
		Do()
}