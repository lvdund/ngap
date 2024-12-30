package ies

import "github.com/lvdund/ngap/aper"

type ForbiddenAreaInformationItem struct {
	PLMNIdentity  []byte
	// ForbiddenTACs []TAC
	// IEExtensions  *ForbiddenAreaInformationItemExtIEs
}

func (ie *ForbiddenAreaInformationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PLMNIdentity := NewOCTETSTRING(ie.PLMNIdentity, aper.Constraint{Lb: 3, Ub: 3}, true)
	if err = tmp_PLMNIdentity.Encode(w); err != nil {
		return
	}
	// if len(ie.ForbiddenTACs) > 0 {
	// 	tmp := Sequence[*TAC]{
	// 		Value: []*TAC{},
	// 		c:     aper.Constraint{Lb: 1, Ub: maxnoofForbTACs},
	// 		ext:   false,
	// 	}
	// 	for _, i := range ie.ForbiddenTACs {
	// 		tmp.Value = append(tmp.Value, &i)
	// 	}
	// 	if err = tmp.Encode(w); err != nil {
	// 		return
	// 	}
	// }
	return
}
func (ie *ForbiddenAreaInformationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_PLMNIdentity := OCTETSTRING{
		c:   aper.Constraint{Lb: 3, Ub: 3},
		ext: false,
	}
	if err = tmp_PLMNIdentity.Decode(r); err != nil {
		return
	}
	ie.PLMNIdentity = tmp_PLMNIdentity.Value
	// tmp_ForbiddenTACs := Sequence[*TAC]{
	// 	c:   aper.Constraint{Lb: 1, Ub: maxnoofForbTACs},
	// 	ext: false,
	// }
	// if err = tmp_ForbiddenTACs.Decode(r); err != nil {
	// 	return
	// }
	// ie.ForbiddenTACs = []TAC{}
	// for _, i := range tmp_ForbiddenTACs.Value {
	// 	ie.ForbiddenTACs = append(ie.ForbiddenTACs, *i)
	// }
	return
}
