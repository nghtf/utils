package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

// Calculate MD5 for a file
func GetMD5(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	md5 := md5.New()
	if _, err := io.Copy(md5, f); err != nil {
		return "", err
	}

	HashMD5 := hex.EncodeToString(md5.Sum(nil))
	f.Close()
	return HashMD5, nil
}

// Calculate SHA256 for a file
func GetSHA256(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	sha := sha256.New()
	if _, err := io.Copy(sha, f); err != nil {
		return "", err
	}

	HashSHA256 := hex.EncodeToString(sha.Sum(nil))
	f.Close()
	return HashSHA256, nil
}
