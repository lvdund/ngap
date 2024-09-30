package ie

import (
	"encoding/hex"
	"fmt"
	"ngap/aper"
)

type PlmnIdentity struct {
	PlmnIdentity aper.OctetString `octetstring:"sizeLB:3,sizeUB:3"`
}

func (e *PlmnIdentity) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString([]byte(e.PlmnIdentity), &aper.Constraint{
		Lb: 3,
		Ub: 3,
	}, false)
	return
}

func (e *PlmnIdentity) Decode(r *aper.AperReader) (err error) {
	var octets []byte
	if octets, err = r.ReadOctetString(&aper.Constraint{
		Lb: 3,
		Ub: 3,
	}, false); err != nil {
		return
	}

	e.PlmnIdentity = octets
	mcc, mnc := MccMnc(octets)
	fmt.Printf("Decoded PlmnId: %s-%s\n", mcc, mnc)
	return
}

func MccMnc(id []byte) (mcc, mnc string) {
	//decode from string here
	mccDigit1 := id[0] & 0x0f
	mccDigit2 := (id[0] & 0xf0) >> 4
	mccDigit3 := (id[1] & 0x0f)

	mncDigit1 := (id[2] & 0x0f)
	mncDigit2 := (id[2] & 0xf0) >> 4
	mncDigit3 := (id[1] & 0xf0) >> 4

	plmnIdBytes := []byte{(mccDigit1 << 4) | mccDigit2, (mccDigit3 << 4) | mncDigit1, (mncDigit2 << 4) | mncDigit3}
	plmnId := hex.EncodeToString(plmnIdBytes)
	mcc = plmnId[0:3]
	if plmnId[5] == 'f' {
		mnc = plmnId[3:5]
	} else {
		mnc = plmnId[3:6]
	}

	return
}
