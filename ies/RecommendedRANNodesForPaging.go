package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type RecommendedRANNodesForPaging struct {
	RecommendedRANNodeList []RecommendedRANNodeItem `lb:1,ub:maxnoofRecommendedRANNodes,madatory`
	// IEExtensions *RecommendedRANNodesForPagingExtIEs `optional`
}

func (ie *RecommendedRANNodesForPaging) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if len(ie.RecommendedRANNodeList) > 0 {
		tmp := Sequence[*RecommendedRANNodeItem]{
			Value: []*RecommendedRANNodeItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofRecommendedRANNodes},
			ext:   false,
		}
		for _, i := range ie.RecommendedRANNodeList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = fmt.Errorf("Encode RecommendedRANNodeList: %w", err)
			return
		}
	} else {
		if err != nil {
			err = fmt.Errorf("RecommendedRANNodeList is nil: %w", err)
		}
		return
	}
	return
}
func (ie *RecommendedRANNodesForPaging) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_RecommendedRANNodeList := Sequence[*RecommendedRANNodeItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofRecommendedRANNodes},
		ext: false,
	}
	fn := func() *RecommendedRANNodeItem { return new(RecommendedRANNodeItem) }
	if err = tmp_RecommendedRANNodeList.Decode(r, fn); err != nil {
		err = fmt.Errorf("Read RecommendedRANNodeList: %w", err)
		return
	}
	ie.RecommendedRANNodeList = []RecommendedRANNodeItem{}
	for _, i := range tmp_RecommendedRANNodeList.Value {
		ie.RecommendedRANNodeList = append(ie.RecommendedRANNodeList, *i)
	}
	return
}
