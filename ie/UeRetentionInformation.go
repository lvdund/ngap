package ie

import "ngap/aper"

type UeRetentionInformation struct {
	UeRetentionInformation aper.BitString `bitstring:"sizeLB:0,sizeUB:150"`
}

func (ie *UeRetentionInformation) Decode(r *aper.AperReader) error {
	return nil
}

func (ie *UeRetentionInformation) Encode(r *aper.AperWriter) (err error) {
	if err = r.WriteBitString(ie.UeRetentionInformation.Bytes, uint(ie.UeRetentionInformation.NumBits), &aper.Constraint{Lb: 0, Ub: 150}, false); err != nil {
		return err
	}
	return nil
}
