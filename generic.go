package ngap

import "ngap/aper"

type IE struct {
	Value aper.BitString
}

func NewIE(val []byte) *IE {
	return &IE{
		Value: aper.BitString{
			Bytes:     val,
			NumBits: uint64(8*len(val)),
		},
	}
}

func (ie *IE) Decode(r aper.AperReader) error {
	return nil
}

func (ie *IE) Encode(r aper.AperWriter) (err error) {
	if err = r.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &aper.Constrain{Lb: int64(ie.Value.NumBits), Ub: int64(ie.Value.NumBits)}, false); err != nil {
		return err
	}
	return nil
}

type IEs struct {
	ies []NgapIE
}

func NewIEs(val []NgapIE) *IEs {
	return &IEs{
		ies: val,
	}
}

func (ies *IEs) Decode(r aper.AperReader) error {
	return nil
}

func (ies *IEs) Encode(r aper.AperWriter) (err error) {
	// TODO: encode
	return nil
}
