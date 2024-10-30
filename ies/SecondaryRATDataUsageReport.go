package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type SecondaryRATDataUsageReport struct {
	AMFUENGAPID                             *AMFUENGAPID                             `,ignore,mandatory`
	RANUENGAPID                             *RANUENGAPID                             `,ignore,mandatory`
	PDUSessionResourceSecondaryRATUsageList *PDUSessionResourceSecondaryRATUsageList `,ignore,mandatory`
	HandoverFlag                            *HandoverFlag                            `,ignore,optional`
	UserLocationInformation                 *UserLocationInformation                 `,ignore,optional`
}

func (msg *SecondaryRATDataUsageReport) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_SecondaryRATDataUsageReport, Criticality_PresentIgnore, msg.toIes())
}
func (msg *SecondaryRATDataUsageReport) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.AMFUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.AMFUENGAPID})
	}
	if msg.RANUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RANUENGAPID})
	}
	if msg.PDUSessionResourceSecondaryRATUsageList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceSecondaryRATUsageList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.PDUSessionResourceSecondaryRATUsageList})
	}
	if msg.HandoverFlag != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_HandoverFlag},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.HandoverFlag})
	}
	if msg.UserLocationInformation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UserLocationInformation},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UserLocationInformation})
	}
	return
}
func (msg *SecondaryRATDataUsageReport) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *SecondaryRATDataUsageReport) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_PDUSessionResourceSecondaryRATUsageList:
		var tmp PDUSessionResourceSecondaryRATUsageList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceSecondaryRATUsageList = &tmp
	case ProtocolIEID_HandoverFlag:
		var tmp HandoverFlag
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.HandoverFlag = &tmp
	case ProtocolIEID_UserLocationInformation:
		var tmp UserLocationInformation
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UserLocationInformation = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
