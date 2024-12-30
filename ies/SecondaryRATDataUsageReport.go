package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type SecondaryRATDataUsageReport struct {
	AMFUENGAPID                             int64
	RANUENGAPID                             int64
	PDUSessionResourceSecondaryRATUsageList []PDUSessionResourceSecondaryRATUsageItem
	HandoverFlag                            *HandoverFlag
	UserLocationInformation                 *UserLocationInformation
}

func (msg *SecondaryRATDataUsageReport) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_SecondaryRATDataUsageReport, Criticality_PresentIgnore, msg.toIes())
}
func (msg *SecondaryRATDataUsageReport) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext:   false,
			Value: aper.Integer(msg.AMFUENGAPID),
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
			Value: aper.Integer(msg.RANUENGAPID),
		}})
	tmp_PDUSessionResourceSecondaryRATUsageList := Sequence[*PDUSessionResourceSecondaryRATUsageItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
		ext: false,
	}
	for _, i := range msg.PDUSessionResourceSecondaryRATUsageList {
		tmp_PDUSessionResourceSecondaryRATUsageList.Value = append(tmp_PDUSessionResourceSecondaryRATUsageList.Value, &i)
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceSecondaryRATUsageList},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &tmp_PDUSessionResourceSecondaryRATUsageList,
	})
	if msg.HandoverFlag != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_HandoverFlag},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.HandoverFlag,
		})
	}
	if msg.UserLocationInformation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UserLocationInformation},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UserLocationInformation,
		})
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
	case ProtocolIEID_PDUSessionResourceSecondaryRATUsageList:
		tmp := Sequence[*PDUSessionResourceSecondaryRATUsageItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceSecondaryRATUsageList = []PDUSessionResourceSecondaryRATUsageItem{}
		for _, i := range tmp.Value {
			msg.PDUSessionResourceSecondaryRATUsageList = append(msg.PDUSessionResourceSecondaryRATUsageList, *i)
		}
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
