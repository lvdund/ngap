package ies

import "github.com/lvdund/ngap/aper"

const (
	CauseProtocolTransfersyntaxerror                          aper.Enumerated = 0
	CauseProtocolAbstractsyntaxerrorreject                    aper.Enumerated = 1
	CauseProtocolAbstractsyntaxerrorignoreandnotify           aper.Enumerated = 2
	CauseProtocolMessagenotcompatiblewithreceiverstate        aper.Enumerated = 3
	CauseProtocolSemanticerror                                aper.Enumerated = 4
	CauseProtocolAbstractsyntaxerrorfalselyconstructedmessage aper.Enumerated = 5
	CauseProtocolUnspecified                                  aper.Enumerated = 6
)

type CauseProtocol struct {
	Value aper.Enumerated `True,0,6`
}

func (ie *CauseProtocol) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, true)
	return
}
func (ie *CauseProtocol) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, true)
	ie.Value = aper.Enumerated(v)
	return
}
