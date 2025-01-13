package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PathSwitchRequest struct {
	RANUENGAPID                              int64                                      `lb:0,ub:4294967295,mandatory,reject`
	SourceAMFUENGAPID                        int64                                      `lb:0,ub:1099511627775,mandatory,reject`
	UserLocationInformation                  UserLocationInformation                    `mandatory,ignore`
	UESecurityCapabilities                   UESecurityCapabilities                     `mandatory,ignore`
	PDUSessionResourceToBeSwitchedDLList     []PDUSessionResourceToBeSwitchedDLItem     `lb:1,ub:maxnoofPDUSessions,mandatory,reject`
	PDUSessionResourceFailedToSetupListPSReq []PDUSessionResourceFailedToSetupItemPSReq `lb:1,ub:maxnoofPDUSessions,optional,ignore`
}

func (msg *PathSwitchRequest) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		return
	}
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_PathSwitchRequest, Criticality_PresentReject, ies)
}
func (msg *PathSwitchRequest) toIes() (ies []NgapMessageIE, err error) {
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
	if len(msg.PDUSessionResourceToBeSwitchedDLList) > 0 {
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
	} else {
		err = utils.WrapError("PDUSessionResourceToBeSwitchedDLList is nil", err)
		return
	}
	if len(msg.PDUSessionResourceFailedToSetupListPSReq) > 0 {
		tmp_PDUSessionResourceFailedToSetupListPSReq := Sequence[*PDUSessionResourceFailedToSetupItemPSReq]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range msg.PDUSessionResourceFailedToSetupListPSReq {
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
func (msg *PathSwitchRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := PathSwitchRequestDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*NgapMessageIE),
	}
	if _, err = aper.ReadSequenceOf[NgapMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	if _, ok := decoder.list[ProtocolIEID_RANUENGAPID]; !ok {
		err = fmt.Errorf("Mandatory field RANUENGAPID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_SourceAMFUENGAPID]; !ok {
		err = fmt.Errorf("Mandatory field SourceAMFUENGAPID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_SourceAMFUENGAPID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_UserLocationInformation]; !ok {
		err = fmt.Errorf("Mandatory field UserLocationInformation is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_UserLocationInformation},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_UESecurityCapabilities]; !ok {
		err = fmt.Errorf("Mandatory field UESecurityCapabilities is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_UESecurityCapabilities},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_PDUSessionResourceToBeSwitchedDLList]; !ok {
		err = fmt.Errorf("Mandatory field PDUSessionResourceToBeSwitchedDLList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceToBeSwitchedDLList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type PathSwitchRequestDecoder struct {
	msg      *PathSwitchRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *PathSwitchRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	ieId := msgIe.Id.Value
	if _, ok := decoder.list[ieId]; ok {
		err = fmt.Errorf("Duplicated protocol IEID[%d] found", ieId)
		return
	}
	decoder.list[ieId] = msgIe
	ieR := aper.NewReader(bytes.NewReader(buf))
	msg := decoder.msg
	switch msgIe.Id.Value {
	case ProtocolIEID_RANUENGAPID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4294967295},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RANUENGAPID", err)
			return
		}
		msg.RANUENGAPID = int64(tmp.Value)
	case ProtocolIEID_SourceAMFUENGAPID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SourceAMFUENGAPID", err)
			return
		}
		msg.SourceAMFUENGAPID = int64(tmp.Value)
	case ProtocolIEID_UserLocationInformation:
		var tmp UserLocationInformation
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read UserLocationInformation", err)
			return
		}
		msg.UserLocationInformation = tmp
	case ProtocolIEID_UESecurityCapabilities:
		var tmp UESecurityCapabilities
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read UESecurityCapabilities", err)
			return
		}
		msg.UESecurityCapabilities = tmp
	case ProtocolIEID_PDUSessionResourceToBeSwitchedDLList:
		tmp := Sequence[*PDUSessionResourceToBeSwitchedDLItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		fn := func() *PDUSessionResourceToBeSwitchedDLItem { return new(PDUSessionResourceToBeSwitchedDLItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read PDUSessionResourceToBeSwitchedDLList", err)
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
		fn := func() *PDUSessionResourceFailedToSetupItemPSReq { return new(PDUSessionResourceFailedToSetupItemPSReq) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read PDUSessionResourceFailedToSetupListPSReq", err)
			return
		}
		msg.PDUSessionResourceFailedToSetupListPSReq = []PDUSessionResourceFailedToSetupItemPSReq{}
		for _, i := range tmp.Value {
			msg.PDUSessionResourceFailedToSetupListPSReq = append(msg.PDUSessionResourceFailedToSetupListPSReq, *i)
		}
	default:
		switch msgIe.Criticality.Value {
		case Criticality_PresentReject:
			fmt.Errorf("Not comprehended IE ID 0x%04x (criticality: reject)", msgIe.Id.Value)
		case Criticality_PresentIgnore:
			fmt.Errorf("Not comprehended IE ID 0x%04x (criticality: ignore)", msgIe.Id.Value)
		case Criticality_PresentNotify:
			fmt.Errorf("Not comprehended IE ID 0x%04x (criticality: notify)", msgIe.Id.Value)
		}
		if msgIe.Criticality.Value != Criticality_PresentIgnore {
			decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
				IECriticality: msgIe.Criticality,
				IEID:          msgIe.Id,
				TypeOfError:   TypeOfError{Value: TypeOfErrorNotunderstood},
			})
		}
	}
	return
}
