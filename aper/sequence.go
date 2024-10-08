package aper

import (
	"fmt"
)

func WriteSequenceOf[T AperMarshaller](items []T, aw *AperWriter, c *Constraint, e bool) (err error) {
	defer func() {
		err = aperError("WriteSequenceOf", err)
	}()

	numElems := len(items)

	//determine lower bound and size range (contraintness)
	var lowerBound, sizeRange uint64 = 0, 0
	if c != nil {
		if c.Lb < 0 || uint64(c.Lb) >= POW_16 {
			err = ErrConstraint
			return
		}
		lowerBound = uint64(c.Lb)
		sizeRange = c.Range()
		if sizeRange > 0 && uint64(c.Ub) >= POW_16 { //upper bound too large, set as semi-constraint
			sizeRange = 0
		}
	}

	if uint64(numElems) < lowerBound { //too few items
		err = ErrUnderflow
		return
	}

	if e {
		if sizeRange == 0 { //conflict: no constraint vs extension
			err = ErrInextensible
			return
		}
		//write extension bit if needs
		if err = aw.WriteBool(int64(numElems) > c.Ub); err != nil {
			return
		}
	}
	//NOTE: if sizeRange == 1, no need to write sequence size
	if sizeRange > 1 {
		if err = aw.writeConstraintValue(sizeRange, uint64(numElems)-lowerBound); err != nil {
			return
		}
	} else if sizeRange == 0 { //unconstraint
		if err = aw.align(); err != nil {
			return
		}
		if err = aw.writeValue(uint64(numElems&0xff), 8); err != nil {
			return
		}
	}
	//finally, write all itemst
	for _, item := range items {
		if err = item.Encode(aw); err != nil {
			return
		}
	}

	// with case up_bound = low_bound
	err = aw.flush()

	return
}

func ReadSequenceOf[T any](decoder func(ar *AperReader) (*T, error), ar *AperReader, c *Constraint, e bool) (items []T, err error) {
	//NOTE: decoder is a function that read from the input stream (*AperReader) to decode
	//a specific aper data structure

	//1. determine lower bound and size range (contraintness)
	var lowerBound, sizeRange uint64 = 0, 0
	if c != nil {
		if c.Lb < 0 || uint64(c.Lb) >= POW_16 {
			err = ErrConstraint
			return
		}
		lowerBound = uint64(c.Lb)
		sizeRange = c.Range()
		if sizeRange > 0 && uint64(c.Ub) >= POW_16 { //upper bound too large, set as semi-constraint
			sizeRange = 0
		}
	}

	//2. read extension bit if needs
	var exBit bool
	if e {
		if sizeRange == 0 { //conflict: no constraint vs extension
			err = ErrInextensible
			return
		}

		if exBit, err = ar.ReadBool(); err != nil {
			return
		}
	}

	//3. read num elements
	var numElems uint64
	if sizeRange == 1 {
		numElems = lowerBound
	} else if sizeRange > 1 {
		if numElems, err = ar.readConstraintValue(sizeRange); err != nil {
			return
		}
		numElems += lowerBound
		if exBit && numElems <= uint64(c.Ub) { //check for consitency of extension bit
			err = fmt.Errorf("Inconsistent extension bit")
			return
		}
	} else { //no constraint
		ar.align()
		if numElems, err = ar.readValue(8); err != nil {
			return
		}
	}
	fmt.Printf("number of elem= %d\n", numElems)
	//4. fianly read every elements
	items = make([]T, numElems)
	var tmpItem *T
	for i := 0; i < int(numElems); i++ {
		//fmt.Println("SequenceOf", i)
		if tmpItem, err = decoder(ar); err != nil {
			fmt.Println("\terr")
			return
		}
		//fmt.Printf("----------: %v", *tmpItem)
		items[i] = *tmpItem
	}
	//fmt.Printf(" -> %p\n", tmpItem)
	return
}

func ReadSequenceOfEx[T AperUnmarshaller](fn func() T, ar *AperReader, c *Constraint, e bool) (items []T, err error) {
	decoder := func(ar *AperReader) (*T, error) {
		item := fn()
		if err := item.Decode(ar); err != nil {
			return nil, err
		}
		return &item, nil
	}
	items, err = ReadSequenceOf[T](decoder, ar, c, e)
	return
}

type ListContainer[T AperMarshaller] struct {
	list []T
	e    bool
	c    *Constraint
}

func NewListContainer[T AperMarshaller](list []T, c *Constraint, e bool) ListContainer[T] {
	return ListContainer[T]{
		list: list,
		e:    e,
		c:    c,
	}
}
func (l ListContainer[T]) Encode(aw *AperWriter) (err error) {
	err = WriteSequenceOf[T](l.list, aw, l.c, l.e)
	return
}
