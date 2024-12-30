package ies

import "github.com/lvdund/ngap/aper"

type PLMNSupportItem struct {
	PLMNIdentity     []byte
	SliceSupportList []SliceSupportItem
	// IEExtensions  *PLMNSupportItemExtIEs
}

func (ie *PLMNSupportItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PLMNIdentity := NewOCTETSTRING(ie.PLMNIdentity, aper.Constraint{Lb: 3, Ub: 3}, true)
	if err = tmp_PLMNIdentity.Encode(w); err != nil {
		return
	}
	if len(ie.SliceSupportList) > 0 {
		tmp := Sequence[*SliceSupportItem]{
			Value: []*SliceSupportItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofSliceItems},
			ext:   false,
		}
		for _, i := range ie.SliceSupportList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *PLMNSupportItem) Decode(r *aper.AperReader) (err error) {
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
	tmp_SliceSupportList := Sequence[*SliceSupportItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofSliceItems},
		ext: false,
	}
	if err = tmp_SliceSupportList.Decode(r); err != nil {
		return
	}
	ie.SliceSupportList = []SliceSupportItem{}
	for _, i := range tmp_SliceSupportList.Value {
		ie.SliceSupportList = append(ie.SliceSupportList, *i)
	}
	return
}
