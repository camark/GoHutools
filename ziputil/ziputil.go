package ziputil

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Zip compresses the directory srcDir into the zip file dstZip (recursively).
func Zip(srcDir, dstZip string) error {
	f, err := os.Create(dstZip)
	if err != nil {
		return err
	}
	defer f.Close()

	absSrc, err := filepath.Abs(srcDir)
	if err != nil {
		return err
	}
	// clean trailing separator so base name matches
	clean := filepath.Clean(absSrc)

	zw := zip.NewWriter(f)
	defer zw.Close()

	err = filepath.Walk(clean, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == clean {
			return nil // skip root dir itself
		}
		rel, err := filepath.Rel(clean, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)

		if info.IsDir() {
			// directory entries are implicit — skip (created on unzip)
			return nil
		}
		// regular file
		w, err := zw.Create(rel)
		if err != nil {
			return err
		}
		src, err := os.Open(path)
		if err != nil {
			return err
		}
		defer src.Close()
		_, err = io.Copy(w, src)
		return err
	})
	return err
}

// ZipFile compresses a single file into a zip archive.
func ZipFile(srcFile, dstZip string) error {
	info, err := os.Stat(srcFile)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return fmt.Errorf("ziputil: %s is a directory, use Zip", srcFile)
	}
	f, err := os.Create(dstZip)
	if err != nil {
		return err
	}
	defer f.Close()

	zw := zip.NewWriter(f)
	defer zw.Close()

	w, err := zw.Create(filepath.Base(srcFile))
	if err != nil {
		return err
	}
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer src.Close()
	_, err = io.Copy(w, src)
	return err
}

// Unzip extracts the zip file srcZip into dstDir.
// Throws an error on entries that would escape dstDir (path traversal).
func Unzip(srcZip, dstDir string) error {
	absDst, err := filepath.Abs(dstDir)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(absDst, 0o755); err != nil {
		return err
	}

	r, err := zip.OpenReader(srcZip)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		clean := filepath.Clean(filepath.FromSlash(f.Name))
		// guard against path traversal
		if strings.HasPrefix(clean, ".."+string(os.PathSeparator)) || clean == ".." || filepath.IsAbs(clean) {
			return fmt.Errorf("ziputil: illegal path in archive: %s", f.Name)
		}
		target := filepath.Join(absDst, clean)

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(target, 0o755); err != nil {
				return err
			}
			continue
		}

		// ensure parent dir exists
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return err
		}
		rc, err := f.Open()
		if err != nil {
			return err
		}
		out, err := os.Create(target)
		if err != nil {
			rc.Close()
			return err
		}
		_, err = io.Copy(out, rc)
		out.Close()
		rc.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// ListFiles returns the list of file names inside the zip archive.
func ListFiles(srcZip string) ([]string, error) {
	r, err := zip.OpenReader(srcZip)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	files := make([]string, 0, len(r.File))
	for _, f := range r.File {
		if !f.FileInfo().IsDir() {
			files = append(files, f.Name)
		}
	}
	return files, nil
}

// Gzip compresses data using gzip.
func Gzip(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	if _, err := gw.Write(data); err != nil {
		return nil, err
	}
	if err := gw.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Gunzip decompresses gzip-compressed data.
func Gunzip(data []byte) ([]byte, error) {
	gr, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer gr.Close()
	return io.ReadAll(gr)
}

// GzipFile compresses srcFile into gzFile.
func GzipFile(srcFile, gzFile string) error {
	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(gzFile)
	if err != nil {
		return err
	}
	defer out.Close()
	gw := gzip.NewWriter(out)
	defer gw.Close()
	_, err = io.Copy(gw, in)
	return err
}

// GunzipFile decompresses gzFile into outFile (or stdout when outFile is "").
func GunzipFile(gzFile, outFile string) error {
	in, err := os.Open(gzFile)
	if err != nil {
		return err
	}
	defer in.Close()
	gr, err := gzip.NewReader(in)
	if err != nil {
		return err
	}
	defer gr.Close()

	var out *os.File
	if outFile == "" {
		out = os.Stdout
	} else {
		out, err = os.Create(outFile)
		if err != nil {
			return err
		}
		defer out.Close()
	}
	_, err = io.Copy(out, gr)
	return err
}

// TarGz compresses the directory srcDir into a tar.gz file.
func TarGz(srcDir, dstTarGz string) error {
	out, err := os.Create(dstTarGz)
	if err != nil {
		return err
	}
	defer out.Close()
	gw := gzip.NewWriter(out)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	absSrc, err := filepath.Abs(srcDir)
	if err != nil {
		return err
	}
	clean := filepath.Clean(absSrc)

	return filepath.Walk(clean, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == clean {
			return nil
		}
		rel, err := filepath.Rel(clean, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)

		hdr, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		hdr.Name = rel
		if info.IsDir() {
			hdr.Name += "/"
		}
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		src, err := os.Open(path)
		if err != nil {
			return err
		}
		defer src.Close()
		_, err = io.Copy(tw, src)
		return err
	})
}

// UntarGz extracts a tar.gz file into dstDir.
func UntarGz(srcTarGz, dstDir string) error {
	in, err := os.Open(srcTarGz)
	if err != nil {
		return err
	}
	defer in.Close()
	gr, err := gzip.NewReader(in)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)

	absDst, err := filepath.Abs(dstDir)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(absDst, 0o755); err != nil {
		return err
	}

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		clean := filepath.Clean(filepath.FromSlash(hdr.Name))
		if strings.HasPrefix(clean, ".."+string(os.PathSeparator)) || filepath.IsAbs(clean) {
			return fmt.Errorf("ziputil: illegal path in tar: %s", hdr.Name)
		}
		target := filepath.Join(absDst, clean)

		switch hdr.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0o755); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
				return err
			}
			out, err := os.Create(target)
			if err != nil {
				return err
			}
			if _, err := io.Copy(out, tr); err != nil {
				out.Close()
				return err
			}
			out.Close()
		}
	}
	return nil
}