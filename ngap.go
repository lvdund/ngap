package ngap

//import "ngap/aper"

const (
	NgapPresentNothing uint8 = iota
	NgapPduInitiatingMessage
	NgapPduSuccessfulOutcome
	NgapPduUnsuccessfulOutcome
)

/*
type IE struct {
	Value aper.BitString
}

func NewIE(val []byte) *IE {
	return &IE{
		Value: aper.BitString{
			Bytes:   val,
			NumBits: uint64(8 * len(val)),
		},
	}
}

func (ie *IE) Decode(r aper.AperReader) error {
	return nil
}

func (ie *IE) Encode(r aper.AperWriter) (err error) {
	if err = r.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &aper.Constraint{Lb: int64(ie.Value.NumBits), Ub: int64(ie.Value.NumBits)}, false); err != nil {
		return err
	}
	return nil
}
*/
