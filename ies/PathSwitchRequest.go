package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PathSwitchRequest struct {
	RANUENGAPID                              *RANUENGAPID                              `,reject,mandatory`
	SourceAMFUENGAPID                        *AMFUENGAPID                              `,reject,mandatory`
	UserLocationInformation                  *UserLocationInformation                  `,ignore,mandatory`
	UESecurityCapabilities                   *UESecurityCapabilities                   `,ignore,mandatory`
	PDUSessionResourceToBeSwitchedDLList     *PDUSessionResourceToBeSwitchedDLList     `,reject,mandatory`
	PDUSessionResourceFailedToSetupListPSReq *PDUSessionResourceFailedToSetupListPSReq `,ignore,optional`
}

func (msg *PathSwitchRequest) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_PathSwitchRequest, Criticality_PresentReject, msg.toIes())
}
func (msg *PathSwitchRequest) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.RANUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.RANUENGAPID})
	}
	if msg.SourceAMFUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SourceAMFUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.SourceAMFUENGAPID})
	}
	if msg.UserLocationInformation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UserLocationInformation},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UserLocationInformation})
	}
	if msg.UESecurityCapabilities != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UESecurityCapabilities},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UESecurityCapabilities})
	}
	if msg.PDUSessionResourceToBeSwitchedDLList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceToBeSwitchedDLList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.PDUSessionResourceToBeSwitchedDLList})
	}
	if msg.PDUSessionResourceFailedToSetupListPSReq != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceFailedToSetupListPSReq},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.PDUSessionResourceFailedToSetupListPSReq})
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
	case ProtocolIEID_RANUENGAPID:
		var tmp RANUENGAPID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RANUENGAPID = &tmp
	case ProtocolIEID_SourceAMFUENGAPID:
		var tmp AMFUENGAPID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SourceAMFUENGAPID = &tmp
	case ProtocolIEID_UserLocationInformation:
		var tmp UserLocationInformation
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UserLocationInformation = &tmp
	case ProtocolIEID_UESecurityCapabilities:
		var tmp UESecurityCapabilities
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UESecurityCapabilities = &tmp
	case ProtocolIEID_PDUSessionResourceToBeSwitchedDLList:
		var tmp PDUSessionResourceToBeSwitchedDLList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceToBeSwitchedDLList = &tmp
	case ProtocolIEID_PDUSessionResourceFailedToSetupListPSReq:
		var tmp PDUSessionResourceFailedToSetupListPSReq
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceFailedToSetupListPSReq = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
