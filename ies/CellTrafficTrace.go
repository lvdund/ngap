package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type CellTrafficTrace struct {
	AMFUENGAPID                    int64
	RANUENGAPID                    int64
	NGRANTraceID                   []byte
	NGRANCGI                       NGRANCGI
	TraceCollectionEntityIPAddress []byte
}

func (msg *CellTrafficTrace) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_CellTrafficTrace, Criticality_PresentIgnore, msg.toIes())
}
func (msg *CellTrafficTrace) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext:   false,
			Value: aper.Integer(msg.AMFUENGAPID),
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
			Value: aper.Integer(msg.RANUENGAPID),
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_NGRANTraceID},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value: &OCTETSTRING{
			c:     aper.Constraint{Lb: 8, Ub: 8},
			ext:   false,
			Value: msg.NGRANTraceID,
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_NGRANCGI},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &msg.NGRANCGI,
	})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TraceCollectionEntityIPAddress},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value: &BITSTRING{
			Value: aper.BitString{
				Bytes:   msg.TraceCollectionEntityIPAddress,
				NumBits: uint64(len(msg.TraceCollectionEntityIPAddress))},
		}})
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
	var id int64
	var c uint64
	var buf []byte
	if id, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: int64(aper.POW_16) - 1}, false); err != nil {
		return
	}
	msgIe = new(NgapMessageIE)
	msgIe.Id.Value = aper.Integer(id)
	if c, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return
	}
	msgIe.Criticality.Value = aper.Enumerated(c)
	if buf, err = r.ReadOpenType(); err != nil {
		return
	}
	ieR := aper.NewReader(bytes.NewReader(buf))
	switch msgIe.Id.Value {
	case ProtocolIEID_AMFUENGAPID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFUENGAPID = int64(tmp.Value)
	case ProtocolIEID_RANUENGAPID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4294967295},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RANUENGAPID = int64(tmp.Value)
	case ProtocolIEID_NGRANTraceID:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 8, Ub: 8},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NGRANTraceID = tmp.Value
	case ProtocolIEID_NGRANCGI:
		var tmp NGRANCGI
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NGRANCGI = tmp
	case ProtocolIEID_TraceCollectionEntityIPAddress:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 160},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.TraceCollectionEntityIPAddress = tmp.Value.Bytes
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
