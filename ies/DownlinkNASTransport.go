package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type DownlinkNASTransport struct {
	AMFUENGAPID               int64
	RANUENGAPID               int64
	OldAMF                    *[]byte
	RANPagingPriority         *int64
	NASPDU                    []byte
	MobilityRestrictionList   *MobilityRestrictionList
	IndexToRFSP               *int64
	UEAggregateMaximumBitRate *UEAggregateMaximumBitRate
	AllowedNSSAI              *[]AllowedNSSAIItem
}

func (msg *DownlinkNASTransport) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_DownlinkNASTransport, Criticality_PresentIgnore, msg.toIes())
}
func (msg *DownlinkNASTransport) toIes() (ies []NgapMessageIE) {
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
	if msg.OldAMF != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_OldAMF},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 1, Ub: 150},
				ext:   true,
				Value: *msg.OldAMF,
			}})
	}
	if msg.RANPagingPriority != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANPagingPriority},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 1, Ub: 256},
				ext:   false,
				Value: aper.Integer(*msg.RANPagingPriority),
			}})
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_NASPDU},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &OCTETSTRING{
			c:     aper.Constraint{Lb: 0, Ub: 0},
			ext:   false,
			Value: msg.NASPDU,
		}})
	if msg.MobilityRestrictionList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_MobilityRestrictionList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.MobilityRestrictionList,
		})
	}
	if msg.IndexToRFSP != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_IndexToRFSP},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 1, Ub: 256},
				ext:   true,
				Value: aper.Integer(*msg.IndexToRFSP),
			}})
	}
	if msg.UEAggregateMaximumBitRate != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UEAggregateMaximumBitRate},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UEAggregateMaximumBitRate,
		})
	}
	if msg.AllowedNSSAI != nil {
		tmp_AllowedNSSAI := Sequence[*AllowedNSSAIItem]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		for _, i := range *msg.AllowedNSSAI {
			tmp_AllowedNSSAI.Value = append(tmp_AllowedNSSAI.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AllowedNSSAI},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_AllowedNSSAI,
		})
	}
	return
}
func (msg *DownlinkNASTransport) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *DownlinkNASTransport) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_OldAMF:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 150},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.OldAMF = tmp.Value
	case ProtocolIEID_RANPagingPriority:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 256},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.RANPagingPriority = int64(tmp.Value)
	case ProtocolIEID_NASPDU:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NASPDU = tmp.Value
	case ProtocolIEID_MobilityRestrictionList:
		var tmp MobilityRestrictionList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.MobilityRestrictionList = &tmp
	case ProtocolIEID_IndexToRFSP:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 256},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.IndexToRFSP = int64(tmp.Value)
	case ProtocolIEID_UEAggregateMaximumBitRate:
		var tmp UEAggregateMaximumBitRate
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UEAggregateMaximumBitRate = &tmp
	case ProtocolIEID_AllowedNSSAI:
		tmp := Sequence[*AllowedNSSAIItem]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AllowedNSSAI = &[]AllowedNSSAIItem{}
		for _, i := range tmp.Value {
			*msg.AllowedNSSAI = append(*msg.AllowedNSSAI, *i)
		}
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
