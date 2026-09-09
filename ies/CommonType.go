package ies

import (
	"github.com/lvdund/ngap/aper"
)

type ENUMERATED struct {
	Value aper.Enumerated
	c     aper.Constraint
	ext   bool
}

func NewENUMERATED(v int64, c aper.Constraint, ext bool) ENUMERATED {
	return ENUMERATED{
		Value: aper.Enumerated(v),
		c:     c,
		ext:   ext,
	}
}
func (t *ENUMERATED) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(t.Value), t.c, t.ext)
	return
}
func (t *ENUMERATED) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(t.c, t.ext)
	t.Value = aper.Enumerated(v)
	return
}

type BITSTRING struct {
	Value aper.BitString
	c     aper.Constraint
	ext   bool
}

func NewBITSTRING(v aper.BitString, c aper.Constraint, ext bool) BITSTRING {
	return BITSTRING{
		Value: aper.BitString{
			Bytes:   v.Bytes,
			NumBits: v.NumBits,
		},
		c:   c,
		ext: ext,
	}
}
func (t *BITSTRING) Encode(w *aper.AperWriter) (err error) {
	if t.c.Lb == t.c.Ub {
		t.Value.NumBits = uint64(t.c.Lb)
	} else if len(t.Value.Bytes)*8 < int(t.c.Lb) {
		t.Value.NumBits = uint64(t.c.Lb)
	}
	err = w.WriteBitString(t.Value.Bytes, uint(t.Value.NumBits), &t.c, t.ext)
	return
}
func (t *BITSTRING) Decode(r *aper.AperReader) (err error) {
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&t.c, t.ext); err != nil {
		return
	}
	t.Value.Bytes = v
	t.Value.NumBits = uint64(n)
	return
}

type OCTETSTRING struct {
	Value aper.OctetString
	c     aper.Constraint
	ext   bool
}

func NewOCTETSTRING(v []byte, c aper.Constraint, ext bool) OCTETSTRING {
	return OCTETSTRING{
		Value: v,
		c:     c,
		ext:   ext,
	}
}
func (t *OCTETSTRING) Encode(w *aper.AperWriter) (err error) {
	if t.c.Lb == t.c.Ub && t.c.Lb == 0 {
		err = w.WriteOctetString(t.Value, nil, t.ext)
	} else {
		err = w.WriteOctetString(t.Value, &t.c, t.ext)
	}
	return
}
func (t *OCTETSTRING) Decode(r *aper.AperReader) (err error) {
	var v aper.OctetString
	if t.c.Lb == t.c.Ub && t.c.Lb == 0 {
		if v, err = r.ReadOctetString(nil, t.ext); err != nil {
			return
		}
	} else {
		if v, err = r.ReadOctetString(&t.c, t.ext); err != nil {
			return
		}
	}

	t.Value = v
	return
}

type INTEGER struct {
	Value aper.Integer
	c     aper.Constraint
	ext   bool
}

func NewINTEGER(v int64, c aper.Constraint, ext bool) INTEGER {
	return INTEGER{
		Value: aper.Integer(v),
		c:     c,
		ext:   ext,
	}
}
func (t *INTEGER) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(t.Value), &t.c, t.ext)
	return
}
func (t *INTEGER) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadInteger(&t.c, t.ext)
	t.Value = aper.Integer(v)
	return
}

type Sequence[T aper.IE] struct {
	Value []T
	c     aper.Constraint
	ext   bool
}

func NewSequence[T aper.IE](items []T, c aper.Constraint, ext bool) Sequence[T] {
	return Sequence[T]{
		Value: items,
		c:     c,
		ext:   ext,
	}
}

func (s *Sequence[T]) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[T](s.Value, w, &s.c, s.ext); err != nil {
		return
	}
	return
}
func (s *Sequence[T]) Decode(r *aper.AperReader, fn func() T) (err error) {
	var newItems []T
	newItems, err = aper.ReadSequenceOfEx(fn, r, &s.c, s.ext)
	if err != nil {
		return
	}
	s.Value = []T{}
	s.Value = append(s.Value, newItems...)
	return
}

// TAC and PLMNIdentity are both OCTET STRING (SIZE(3)). They exist as named
// types only because a few IEs carry them inside a SEQUENCE OF, which needs an
// element implementing aper.IE; everywhere else the same three octets are a
// plain []byte field encoded through OCTETSTRING, and these encode identically.
//
// A fixed-size octet string writes no length determinant, so the whole value is
// the three octets the caller supplies: an encoder that writes nothing here
// leaves the sequence claiming an element it never wrote, and a peer then reads
// the following three octets as this one.
type TAC struct {
	Value []byte
}

func (ie *TAC) Encode(w *aper.AperWriter) (err error) {
	t := NewOCTETSTRING(ie.Value, aper.Constraint{Lb: 3, Ub: 3}, false)
	return t.Encode(w)
}
func (ie *TAC) Decode(r *aper.AperReader) (err error) {
	t := OCTETSTRING{
		c:   aper.Constraint{Lb: 3, Ub: 3},
		ext: false,
	}
	if err = t.Decode(r); err != nil {
		return
	}
	ie.Value = t.Value
	return
}

type PLMNIdentity struct {
	Value []byte
}

func (ie *PLMNIdentity) Encode(w *aper.AperWriter) (err error) {
	t := NewOCTETSTRING(ie.Value, aper.Constraint{Lb: 3, Ub: 3}, false)
	return t.Encode(w)
}
func (ie *PLMNIdentity) Decode(r *aper.AperReader) (err error) {
	t := OCTETSTRING{
		c:   aper.Constraint{Lb: 3, Ub: 3},
		ext: false,
	}
	if err = t.Decode(r); err != nil {
		return
	}
	ie.Value = t.Value
	return
}

// EmergencyAreaID is OCTET STRING (SIZE(3)), the same three octets that every
// IE carrying one outside a SEQUENCE OF declares as a []byte with lb:3,ub:3.
type EmergencyAreaID struct {
	Value []byte
}

func (ie *EmergencyAreaID) Encode(w *aper.AperWriter) (err error) {
	t := NewOCTETSTRING(ie.Value, aper.Constraint{Lb: 3, Ub: 3}, false)
	return t.Encode(w)
}
func (ie *EmergencyAreaID) Decode(r *aper.AperReader) (err error) {
	t := OCTETSTRING{
		c:   aper.Constraint{Lb: 3, Ub: 3},
		ext: false,
	}
	if err = t.Decode(r); err != nil {
		return
	}
	ie.Value = t.Value
	return
}

// TransportLayerAddress is BIT STRING (SIZE(1..160,...)) - an extensible bit
// string, not an octet string, which is how XnExtTLAItem already carries the
// same value in the field beside a list of these.
//
// Value holds whole octets and the bit count follows from its length, which
// covers every address the protocol carries: 32 bits for IPv4, 128 for IPv6,
// 160 for both. A length that is not a whole number of octets cannot be
// expressed here and does not arise.
type TransportLayerAddress struct {
	Value []byte
}

func (ie *TransportLayerAddress) Encode(w *aper.AperWriter) (err error) {
	t := NewBITSTRING(aper.BitString{
		Bytes:   ie.Value,
		NumBits: uint64(len(ie.Value)) * 8,
	}, aper.Constraint{Lb: 1, Ub: 160}, true)
	return t.Encode(w)
}
func (ie *TransportLayerAddress) Decode(r *aper.AperReader) (err error) {
	t := BITSTRING{
		c:   aper.Constraint{Lb: 1, Ub: 160},
		ext: true,
	}
	if err = t.Decode(r); err != nil {
		return
	}
	ie.Value = t.Value.Bytes
	return
}

const (
	maxnoofAllowedAreas              int64 = 16
	maxnoofAllowedSNSSAIs            int64 = 8
	maxnoofBPLMNs                    int64 = 12
	maxnoofCellIDforWarning          int64 = 65535
	maxnoofCellinAoI                 int64 = 256
	maxnoofCellinEAI                 int64 = 65535
	maxnoofCellinTAI                 int64 = 65535
	maxnoofCellsingNB                int64 = 16384
	maxnoofCellsinngeNB              int64 = 256
	maxnoofCellsinUEHistoryInfo      int64 = 16
	maxnoofCellsUEMovingTrajectory   int64 = 16
	maxnoofDRBs                      int64 = 32
	maxnoofEmergencyAreaID           int64 = 65535
	maxnoofEAIforRestart             int64 = 256
	maxnoofEPLMNs                    int64 = 15
	maxnoofEPLMNsPlusOne             int64 = 16
	maxnoofERABs                     int64 = 256
	maxnoofErrors                    int64 = 256
	maxnoofForbTACs                  int64 = 4096
	maxnoofMultiConnectivity         int64 = 4
	maxnoofMultiConnectivityMinusOne int64 = 3
	maxnoofNGConnectionsToReset      int64 = 65536
	maxnoofPDUSessions               int64 = 256
	maxnoofPLMNs                     int64 = 12
	maxnoofQosFlows                  int64 = 64
	maxnoofRANNodeinAoI              int64 = 64
	maxnoofRecommendedCells          int64 = 16
	maxnoofRecommendedRANNodes       int64 = 16
	maxnoofAoI                       int64 = 64
	maxnoofServedGUAMIs              int64 = 256
	maxnoofSliceItems                int64 = 1024
	maxnoofTACs                      int64 = 256
	maxnoofTAIforInactive            int64 = 16
	maxnoofTAIforPaging              int64 = 16
	maxnoofTAIforRestart             int64 = 2048
	maxnoofTAIforWarning             int64 = 65535
	maxnoofTAIinAoI                  int64 = 16
	maxnoofTimePeriods               int64 = 2
	maxnoofTNLAssociations           int64 = 32
	maxnoofXnExtTLAs                 int64 = 16
	maxnoofXnGTPTLAs                 int64 = 16
	maxnoofXnTLAs                    int64 = 2
)
