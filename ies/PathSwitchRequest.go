package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PathSwitchRequest struct {
	RANUENGAPID                              int64
	SourceAMFUENGAPID                        int64
	UserLocationInformation                  UserLocationInformation
	UESecurityCapabilities                   UESecurityCapabilities
	PDUSessionResourceToBeSwitchedDLList     []PDUSessionResourceToBeSwitchedDLItem
	PDUSessionResourceFailedToSetupListPSReq *[]PDUSessionResourceFailedToSetupItemPSReq
}

func (msg *PathSwitchRequest) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_PathSwitchRequest, Criticality_PresentReject, msg.toIes())
}
func (msg *PathSwitchRequest) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
			Value: aper.Integer(msg.RANUENGAPID),
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_SourceAMFUENGAPID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext:   false,
			Value: aper.Integer(msg.SourceAMFUENGAPID),
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_UserLocationInformation},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &msg.UserLocationInformation,
	})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_UESecurityCapabilities},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &msg.UESecurityCapabilities,
	})
	tmp_PDUSessionResourceToBeSwitchedDLList := Sequence[*PDUSessionResourceToBeSwitchedDLItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
		ext: false,
	}
	for _, i := range msg.PDUSessionResourceToBeSwitchedDLList {
		tmp_PDUSessionResourceToBeSwitchedDLList.Value = append(tmp_PDUSessionResourceToBeSwitchedDLList.Value, &i)
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceToBeSwitchedDLList},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &tmp_PDUSessionResourceToBeSwitchedDLList,
	})
	if msg.PDUSessionResourceFailedToSetupListPSReq != nil {
		tmp_PDUSessionResourceFailedToSetupListPSReq := Sequence[*PDUSessionResourceFailedToSetupItemPSReq]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range *msg.PDUSessionResourceFailedToSetupListPSReq {
			tmp_PDUSessionResourceFailedToSetupListPSReq.Value = append(tmp_PDUSessionResourceFailedToSetupListPSReq.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceFailedToSetupListPSReq},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_PDUSessionResourceFailedToSetupListPSReq,
		})
	}
	return
}
func (msg *PathSwitchRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *PathSwitchRequest) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_RANUENGAPID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4294967295},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RANUENGAPID = int64(tmp.Value)
	case ProtocolIEID_SourceAMFUENGAPID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SourceAMFUENGAPID = int64(tmp.Value)
	case ProtocolIEID_UserLocationInformation:
		var tmp UserLocationInformation
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UserLocationInformation = tmp
	case ProtocolIEID_UESecurityCapabilities:
		var tmp UESecurityCapabilities
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UESecurityCapabilities = tmp
	case ProtocolIEID_PDUSessionResourceToBeSwitchedDLList:
		tmp := Sequence[*PDUSessionResourceToBeSwitchedDLItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceToBeSwitchedDLList = []PDUSessionResourceToBeSwitchedDLItem{}
		for _, i := range tmp.Value {
			msg.PDUSessionResourceToBeSwitchedDLList = append(msg.PDUSessionResourceToBeSwitchedDLList, *i)
		}
	case ProtocolIEID_PDUSessionResourceFailedToSetupListPSReq:
		tmp := Sequence[*PDUSessionResourceFailedToSetupItemPSReq]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceFailedToSetupListPSReq = &[]PDUSessionResourceFailedToSetupItemPSReq{}
		for _, i := range tmp.Value {
			*msg.PDUSessionResourceFailedToSetupListPSReq = append(*msg.PDUSessionResourceFailedToSetupListPSReq, *i)
		}
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
