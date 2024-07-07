package utils

import (
	"runtime"

	"github.com/google/uuid"
)

// Returns the name of the runtime caller
func CallerName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

// Checks if UUID is valid
func IsNotUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err != nil
}
