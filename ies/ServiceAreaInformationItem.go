package ies

import "github.com/lvdund/ngap/aper"

type ServiceAreaInformationItem struct {
	PLMNIdentity   []byte
	// AllowedTACs    *[]TAC
	// NotAllowedTACs *[]TAC
	// IEExtensions  *ServiceAreaInformationItemExtIEs
}

func (ie *ServiceAreaInformationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	// if ie.AllowedTACs != nil {
	// 	aper.SetBit(optionals, 1)
	// }
	// if ie.NotAllowedTACs != nil {
	// 	aper.SetBit(optionals, 2)
	// }
	w.WriteBits(optionals, 3)
	tmp_PLMNIdentity := NewOCTETSTRING(ie.PLMNIdentity, aper.Constraint{Lb: 3, Ub: 3}, true)
	if err = tmp_PLMNIdentity.Encode(w); err != nil {
		return
	}
	// if ie.AllowedTACs != nil {
	// 	if len(*ie.AllowedTACs) > 0 {
	// 		tmp := Sequence[*TAC]{
	// 			Value: []*TAC{},
	// 			c:     aper.Constraint{Lb: 1, Ub: maxnoofAllowedAreas},
	// 			ext:   false,
	// 		}
	// 		for _, i := range *ie.AllowedTACs {
	// 			tmp.Value = append(tmp.Value, &i)
	// 		}
	// 		if err = tmp.Encode(w); err != nil {
	// 			return
	// 		}
	// 	}
	// }
	// if ie.NotAllowedTACs != nil {
	// 	if len(*ie.NotAllowedTACs) > 0 {
	// 		tmp := Sequence[*TAC]{
	// 			Value: []*TAC{},
	// 			c:     aper.Constraint{Lb: 1, Ub: maxnoofAllowedAreas},
	// 			ext:   false,
	// 		}
	// 		for _, i := range *ie.NotAllowedTACs {
	// 			tmp.Value = append(tmp.Value, &i)
	// 		}
	// 		if err = tmp.Encode(w); err != nil {
	// 			return
	// 		}
	// 	}
	// }
	return
}
func (ie *ServiceAreaInformationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	// var optionals []byte
	if _, err = r.ReadBits(3); err != nil {
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
	// if aper.IsBitSet(optionals, 1) {
	// 	tmp_AllowedTACs := Sequence[*TAC]{
	// 		c:   aper.Constraint{Lb: 1, Ub: maxnoofAllowedAreas},
	// 		ext: false,
	// 	}
	// 	if err = tmp_AllowedTACs.Decode(r); err != nil {
	// 		return
	// 	}
	// 	ie.AllowedTACs = &[]TAC{}
	// 	for _, i := range tmp_AllowedTACs.Value {
	// 		*ie.AllowedTACs = append(*ie.AllowedTACs, *i)
	// 	}
	// }
	// if aper.IsBitSet(optionals, 2) {
	// 	tmp_NotAllowedTACs := Sequence[*TAC]{
	// 		c:   aper.Constraint{Lb: 1, Ub: maxnoofAllowedAreas},
	// 		ext: false,
	// 	}
	// 	if err = tmp_NotAllowedTACs.Decode(r); err != nil {
	// 		return
	// 	}
	// 	ie.NotAllowedTACs = &[]TAC{}
	// 	for _, i := range tmp_NotAllowedTACs.Value {
	// 		*ie.NotAllowedTACs = append(*ie.NotAllowedTACs, *i)
	// 	}
	// }
	return
}
