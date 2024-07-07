package utils

import (
	"fmt"
	"net/http"
	"os"
)

// Return size of the file
func GetFileSize(path string) (int64, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	stat, err := f.Stat()
	if err != nil {
		return 0, err
	}
	return stat.Size(), nil
}

// Convert number to human-readable format (kB, Mb, etc)
func SizeHumanized(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}

// Checks if file path exists
func Exist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist((err)) {
		return false
	}
	return true
}

// Detects MIME type with http.DetectContentType and first 512 bytes header. Returns "application/octet-stream" by default
func GetFileContentType(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	buf := make([]byte, 512)
	if _, err = f.Read(buf); err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buf)
	return contentType, nil
}
