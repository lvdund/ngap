package aper

//NGAP types should be auto-generated using following templates

type IE interface {
	Encode(*AperWriter) error
	Decode(*AperReader) error
}

// AmfId
type AmfId struct {
	Region      []byte     //`bitstring:"sizeLB:8,sizeUB:8"`
	Set         []byte     //`bitstring:"sizeLB:10,sizeUB:10"`
	Pointer     []byte     //`bitstring:"sizeLB:6,sizeUB:6"`
	BitStringEx *BitString //`bitstring:"sizeLB:1,sizeUB:120"`
	ValueEx     *uint64    //`Integer:"valueLB:500,valueUB:1000`

}

func (id *AmfId) Encode(w *AperWriter) (err error) {
	//write extensible bit if needed

	//write option flags
	flags := []byte{0}
	if id.BitStringEx != nil {
		SetBit(flags, 0)
	}
	if id.ValueEx != nil {
		SetBit(flags, 1)
	}
	if err = w.WriteBits(flags, 2); err != nil {
		return
	}

	//write Region
	if err = w.WriteBitString(id.Region, 8, &Constraint{
		Lb: 8,
		Ub: 8,
	}, false); err != nil {
		return
	}
	//write Set
	if err = w.WriteBitString(id.Set, 10, &Constraint{
		Lb: 10,
		Ub: 10,
	}, false); err != nil {
		return
	}
	//write Pointer
	if err = w.WriteBitString(id.Pointer, 8, &Constraint{
		Lb: 6,
		Ub: 6,
	}, false); err != nil {
		return
	}
	if id.BitStringEx != nil {
		if err = w.WriteBitString(id.BitStringEx.Bytes, uint(id.BitStringEx.NumBits), &Constraint{
			Lb: 1,
			Ub: 120,
		}, false); err != nil {
			return
		}
	}
	if id.ValueEx != nil {
		if err = w.WriteInteger(int64(*id.ValueEx), &Constraint{
			Lb: 500,
			Ub: 1000,
		}, false); err != nil {
			return
		}
	}

	return nil
}

func (id *AmfId) Decode(r *AperReader) (err error) {
	if id == nil {
		id = &AmfId{}
	}
	//read exensible bit if needed
	//read option flag if needed
	var flags []byte
	if flags, err = r.ReadBits(2); err != nil {
		return
	}
	//read Region
	if id.Region, _, err = r.ReadBitString(&Constraint{
		Lb: 8,
		Ub: 8,
	}, false); err != nil {
		return
	}
	//read Set
	if id.Set, _, err = r.ReadBitString(&Constraint{
		Lb: 10,
		Ub: 10,
	}, false); err != nil {
		return
	}
	//read Pointer
	if id.Pointer, _, err = r.ReadBitString(&Constraint{
		Lb: 6,
		Ub: 6,
	}, false); err != nil {
		return
	}
	//read BitStringEx
	if IsBitSet(flags, 0) {

	}
	//read ValueEx
	if IsBitSet(flags, 1) {
	}
	return
}

// AmfName
type AmfName string // `aper:"sizeExt,sizeLB:1,sizeUB:150"`

func (n AmfName) Encode(w *AperWriter) (err error) {
	return w.WriteOctetString([]byte(n), &Constraint{
		Lb: 1,
		Ub: 150,
	}, true)
}

func (n AmfName) Decode(r *AperReader) (err error) {
	var octets []byte
	if octets, err = r.ReadOctetString(&Constraint{
		Lb: 1,
		Ub: 150,
	}, true); err != nil {
		return
	}
	n = AmfName(string(octets))
	return
}
