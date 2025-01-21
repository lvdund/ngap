package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type HandoverRequestAcknowledge struct {
	AMFUENGAPID                              int64                                      `lb:0,ub:1099511627775,mandatory,ignore`
	RANUENGAPID                              int64                                      `lb:0,ub:4294967295,mandatory,ignore`
	PDUSessionResourceAdmittedList           []PDUSessionResourceAdmittedItem           `lb:1,ub:maxnoofPDUSessions,mandatory,ignore`
	PDUSessionResourceFailedToSetupListHOAck []PDUSessionResourceFailedToSetupItemHOAck `lb:1,ub:maxnoofPDUSessions,optional,mandatory,ignore`
	TargetToSourceTransparentContainer       []byte                                     `lb:0,ub:0,mandatory,reject`
	CriticalityDiagnostics                   *CriticalityDiagnostics                    `optional,mandatory,ignore`
}

func (msg *HandoverRequestAcknowledge) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("HandoverRequestAcknowledge"), err)
		return
	}
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_HandoverResourceAllocation, Criticality_PresentReject, ies)
}
func (msg *HandoverRequestAcknowledge) toIes() (ies []NgapMessageIE, err error) {
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
	if len(msg.PDUSessionResourceAdmittedList) > 0 {
		tmp_PDUSessionResourceAdmittedList := Sequence[*PDUSessionResourceAdmittedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range msg.PDUSessionResourceAdmittedList {
			tmp_PDUSessionResourceAdmittedList.Value = append(tmp_PDUSessionResourceAdmittedList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceAdmittedList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_PDUSessionResourceAdmittedList,
		})
	} else {
		err = utils.WrapError("PDUSessionResourceAdmittedList is nil", err)
		return
	}
	if len(msg.PDUSessionResourceFailedToSetupListHOAck) > 0 {
		tmp_PDUSessionResourceFailedToSetupListHOAck := Sequence[*PDUSessionResourceFailedToSetupItemHOAck]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range msg.PDUSessionResourceFailedToSetupListHOAck {
			tmp_PDUSessionResourceFailedToSetupListHOAck.Value = append(tmp_PDUSessionResourceFailedToSetupListHOAck.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceFailedToSetupListHOAck},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_PDUSessionResourceFailedToSetupListHOAck,
		})
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TargetToSourceTransparentContainer},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &OCTETSTRING{
			c:     aper.Constraint{Lb: 0, Ub: 0},
			ext:   false,
			Value: msg.TargetToSourceTransparentContainer,
		}})
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	return
}
func (msg *HandoverRequestAcknowledge) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("HandoverRequestAcknowledge"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := HandoverRequestAcknowledgeDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*NgapMessageIE),
	}
	if _, err = aper.ReadSequenceOf[NgapMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	if _, ok := decoder.list[ProtocolIEID_AMFUENGAPID]; !ok {
		err = fmt.Errorf("Mandatory field AMFUENGAPID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_RANUENGAPID]; !ok {
		err = fmt.Errorf("Mandatory field RANUENGAPID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_PDUSessionResourceAdmittedList]; !ok {
		err = fmt.Errorf("Mandatory field PDUSessionResourceAdmittedList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceAdmittedList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_TargetToSourceTransparentContainer]; !ok {
		err = fmt.Errorf("Mandatory field TargetToSourceTransparentContainer is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_TargetToSourceTransparentContainer},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type HandoverRequestAcknowledgeDecoder struct {
	msg      *HandoverRequestAcknowledge
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *HandoverRequestAcknowledgeDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_AMFUENGAPID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read AMFUENGAPID", err)
			return
		}
		msg.AMFUENGAPID = int64(tmp.Value)
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
	case ProtocolIEID_PDUSessionResourceAdmittedList:
		tmp := Sequence[*PDUSessionResourceAdmittedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		fn := func() *PDUSessionResourceAdmittedItem { return new(PDUSessionResourceAdmittedItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read PDUSessionResourceAdmittedList", err)
			return
		}
		msg.PDUSessionResourceAdmittedList = []PDUSessionResourceAdmittedItem{}
		for _, i := range tmp.Value {
			msg.PDUSessionResourceAdmittedList = append(msg.PDUSessionResourceAdmittedList, *i)
		}
	case ProtocolIEID_PDUSessionResourceFailedToSetupListHOAck:
		tmp := Sequence[*PDUSessionResourceFailedToSetupItemHOAck]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		fn := func() *PDUSessionResourceFailedToSetupItemHOAck { return new(PDUSessionResourceFailedToSetupItemHOAck) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read PDUSessionResourceFailedToSetupListHOAck", err)
			return
		}
		msg.PDUSessionResourceFailedToSetupListHOAck = []PDUSessionResourceFailedToSetupItemHOAck{}
		for _, i := range tmp.Value {
			msg.PDUSessionResourceFailedToSetupListHOAck = append(msg.PDUSessionResourceFailedToSetupListHOAck, *i)
		}
	case ProtocolIEID_TargetToSourceTransparentContainer:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read TargetToSourceTransparentContainer", err)
			return
		}
		msg.TargetToSourceTransparentContainer = tmp.Value
	case ProtocolIEID_CriticalityDiagnostics:
		var tmp CriticalityDiagnostics
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read CriticalityDiagnostics", err)
			return
		}
		msg.CriticalityDiagnostics = &tmp
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
