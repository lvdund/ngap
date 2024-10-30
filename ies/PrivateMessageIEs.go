package ies

import (
	"io"

	"github.com/lvdund/ngap/aper"
)

type PrivateMessage struct {
}

func (msg *PrivateMessage) Encode(w io.Writer) (err error) {
	return
}

func (msg *PrivateMessage) toIes() (ies []NgapMessageIE) {
	return
}

func (msg *PrivateMessage) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	return
}

func (msg *PrivateMessage) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
	return
}
