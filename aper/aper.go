package aper

type AperMarshaller interface {
	AperEncode(w AperWriter) error
}

type AperUnmarshaller interface {
	AperDecode(w AperWriter) error
}

type BitString struct {
	Bytes   []byte
	NumBits uint64
}

//type OctetString []byte

//type Integer int64

type Constrain struct {
	Lb int64
	Ub int64
}

func (c *Constrain) Range() uint64 {
	return uint64(c.Ub - c.Lb + 1)
}
