package aper

const (
	POW_16 uint64 = 65536
	POW_14 uint64 = 16384
	POW_8  uint64 = 256
	POW_7  uint64 = 128
	POW_6  uint64 = 64
)

type AperMarshaller interface {
	Encode(w AperWriter) error
}

type AperUnmarshaller interface {
	Decode(w AperReader) error
}

type BitString struct {
	Bytes   []byte
	NumBits uint64
}

type OctetString []byte

type Integer int64
type Enumerated int64

type Constraint struct {
	Lb int64
	Ub int64
}

// check if value is in range
func (c *Constraint) InRange(v int64) bool {
	if v < c.Lb {
		return false
	}
	if c.Lb <= c.Ub && v > c.Ub {
		return false
	}
	return true
}

// check if value is unconstrain
func (c *Constraint) IsUnconstrain(v int64) bool {
	if c.Lb > c.Ub && v >= c.Lb {
		return true
	}
	if v > c.Ub {
		return true
	}
	return false
}
func (c *Constraint) Range() uint64 {
	if c.Lb > c.Ub {
		return 0
	}
	return uint64(c.Ub - c.Lb + 1)
}
