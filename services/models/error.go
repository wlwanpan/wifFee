package models

import (
	"strings"
)

type ErrMissingAttr struct {
	attrs []string
}

func (err *ErrMissingAttr) Error() string {
	return "Missing attributes: " + strings.Join(err.attrs, ", ")
}
