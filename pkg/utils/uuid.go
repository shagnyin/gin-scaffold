package utils

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

func UUID() string {
	v4 := uuid.NewV4()
	return strings.ReplaceAll(v4.String(), "-", "")
}
