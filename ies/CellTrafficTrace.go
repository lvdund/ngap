package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type CellTrafficTrace struct {
	AMFUENGAPID                    *AMFUENGAPID           `,reject,mandatory`
	RANUENGAPID                    *RANUENGAPID           `,reject,mandatory`
	NGRANTraceID                   *NGRANTraceID          `,ignore,mandatory`
	NGRANCGI                       *NGRANCGI              `,ignore,mandatory`
	TraceCollectionEntityIPAddress *TransportLayerAddress `,ignore,mandatory`
}

func (msg *CellTrafficTrace) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_CellTrafficTrace, Criticality_PresentIgnore, msg.toIes())
}
func (msg *CellTrafficTrace) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.AMFUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.AMFUENGAPID})
	}
	if msg.RANUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.RANUENGAPID})
	}
	if msg.NGRANTraceID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NGRANTraceID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.NGRANTraceID})
	}
	if msg.NGRANCGI != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NGRANCGI},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.NGRANCGI})
	}
	if msg.TraceCollectionEntityIPAddress != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TraceCollectionEntityIPAddress},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.TraceCollectionEntityIPAddress})
	}
	return
}
func (msg *CellTrafficTrace) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *CellTrafficTrace) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
	id, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: int64(aper.POW_16) - 1}, false)
	if err != nil {
		return
	}
	msgIe = new(NgapMessageIE)
	msgIe.Id.Value = aper.Integer(id)
	c, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return
	}
	msgIe.Criticality.Value = aper.Enumerated(c)
	var buf []byte
	if buf, err = r.ReadOpenType(); err != nil {
		return
	}
	ieR := aper.NewReader(bytes.NewReader(buf))
	switch msgIe.Id.Value {
	case ProtocolIEID_AMFUENGAPID:
		var tmp AMFUENGAPID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFUENGAPID = &tmp
	case ProtocolIEID_RANUENGAPID:
		var tmp RANUENGAPID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RANUENGAPID = &tmp
	case ProtocolIEID_NGRANTraceID:
		var tmp NGRANTraceID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NGRANTraceID = &tmp
	case ProtocolIEID_NGRANCGI:
		var tmp NGRANCGI
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NGRANCGI = &tmp
	case ProtocolIEID_TraceCollectionEntityIPAddress:
		var tmp TransportLayerAddress
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.TraceCollectionEntityIPAddress = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
