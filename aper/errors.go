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

type AperError struct {
	msg  string
	prev error
}

func aperError(msg string, prev error) error {
	if prev == nil {
		return nil
	}
	return &AperError{
		msg:  msg,
		prev: prev,
	}
}

func (aperErr *AperError) Error() (errStr string) {
	errStr = aperErr.msg

	prev := aperErr.prev
	for prev != nil {
		if cur, ok := prev.(*AperError); ok {
			errStr = fmt.Sprintf("%s <<= %s", errStr, cur.msg)
			prev = cur.prev
		} else {
			errStr = fmt.Sprintf("%s <<= %s", errStr, prev.Error())
			prev = nil
		}
	}
	return
}
