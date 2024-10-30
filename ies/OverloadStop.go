package ies

import (
	"io"

	"github.com/lvdund/ngap/aper"
)

type OverloadStop struct {
}

func (msg *OverloadStop) Encode(w io.Writer) (err error) {
	return
}

func (msg *OverloadStop) toIes() (ies []NgapMessageIE) {
	return
}

func (msg *OverloadStop) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	return
}

func (msg *OverloadStop) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
	return
}
