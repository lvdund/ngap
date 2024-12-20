package aper

import (
	"fmt"
)

var (
	ErrCritical      error = fmt.Errorf("Critical")
	ErrUnderflow     error = fmt.Errorf("Underflow")
	ErrOverflow      error = fmt.Errorf("Underflow")
	ErrTail          error = fmt.Errorf("Junk tail")
	ErrIncomplete    error = fmt.Errorf("Data truncated")
	ErrInextensible  error = fmt.Errorf("Field not extensible")
	ErrFixedLength   error = fmt.Errorf("Invalid fixed length")
	ErrConstraint    error = fmt.Errorf("Invalid constraint")
	ErrInvalidLength error = fmt.Errorf("Invalid length")
)
