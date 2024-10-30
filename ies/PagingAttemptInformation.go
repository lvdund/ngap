package ies

import "github.com/lvdund/ngap/aper"

type PagingAttemptInformation struct {
	PagingAttemptCount             *PagingAttemptCount             `False,`
	IntendedNumberOfPagingAttempts *IntendedNumberOfPagingAttempts `False,`
	NextPagingAreaScope            *NextPagingAreaScope            `False,OPTIONAL`
	// IEExtensions PagingAttemptInformationExtIEs `False,OPTIONAL`
}

func (ie *PagingAttemptInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.NextPagingAreaScope != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.PagingAttemptCount != nil {
		if err = ie.PagingAttemptCount.Encode(w); err != nil {
			return
		}
	}
	if ie.IntendedNumberOfPagingAttempts != nil {
		if err = ie.IntendedNumberOfPagingAttempts.Encode(w); err != nil {
			return
		}
	}
	if ie.NextPagingAreaScope != nil {
		if err = ie.NextPagingAreaScope.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *PagingAttemptInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	ie.PagingAttemptCount = new(PagingAttemptCount)
	ie.IntendedNumberOfPagingAttempts = new(IntendedNumberOfPagingAttempts)
	ie.NextPagingAreaScope = new(NextPagingAreaScope)
	if err = ie.PagingAttemptCount.Decode(r); err != nil {
		return
	}
	if err = ie.IntendedNumberOfPagingAttempts.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.NextPagingAreaScope.Decode(r); err != nil {
			return
		}
	}
	return
}
