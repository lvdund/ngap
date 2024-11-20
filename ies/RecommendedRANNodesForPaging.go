package ies

import "github.com/lvdund/ngap/aper"

type RecommendedRANNodesForPaging struct {
	RecommendedRANNodeList *RecommendedRANNodeList `False,`
	// IEExtensions RecommendedRANNodesForPagingExtIEs `False,OPTIONAL`
}

func (ie *RecommendedRANNodesForPaging) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.RecommendedRANNodeList != nil {
		if err = ie.RecommendedRANNodeList.Encode(w); err != nil {
			return
		}
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
	ie.RecommendedRANNodeList = new(RecommendedRANNodeList)
	if err = ie.RecommendedRANNodeList.Decode(r); err != nil {
		return
	}
	return
}
