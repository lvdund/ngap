package aper

import ()

const (
	POW_16 uint64 = 65536
	POW_14 uint64 = 16384
	POW_8  uint64 = 256
	POW_7  uint64 = 128
	POW_6  uint64 = 64
)

type AperMarshaller interface {
	Encode(*AperWriter) error
}

type AperUnmarshaller interface {
	Decode(*AperReader) error
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

/*
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
*/

func (c *Constraint) Range() uint64 {
	if c.Lb > c.Ub {
		return 0
	}
	return uint64(c.Ub - c.Lb + 1)
}

func (aw *AperWriter) writeExtBit(bitsLength uint64, e bool,c *Constraint) (int64, uint64, error) {
	exBit := false
	var lRange uint64 = 0    //length range
	var lowerBound int64 = 0 //length lower bound, default=0

	if c != nil {
		if lowerBound = c.Lb; lowerBound < 0 { //make sure lower bound is not negative
			return 0,0,ErrConstraint
		}
		if int64(bitsLength) <=c.Ub {
			lRange = (c.Range())
		} else if !e {
			//err = ErrInextensible
			return 0,0,ErrInextensible
		}else{
			exBit = true
		}
	}
	
	if e {
		if err := aw.WriteBool(exBit); err != nil {
			return 0,0,nil
		}
	}
    return lowerBound, lRange, nil
}

func (ar *AperReader) readExBit(c *Constraint, e bool) (lRange uint64,lowerBound int64,err error) {
	var exBit bool = false
	if e { //read extension bit
		if exBit, err = ar.ReadBool(); err != nil {
			return 0, 0, err
		}
	}

	if c != nil {
		if lowerBound = c.Lb; lowerBound < 0 { //make sure lower bound is not negative
			return 0, 0, ErrConstraint 
		}
		if !exBit {
			lRange = c.Range()
		}
		if uint64(c.Ub) > POW_16 {
			lRange = c.Range()
		}
	}
	return lRange, lowerBound, nil
}