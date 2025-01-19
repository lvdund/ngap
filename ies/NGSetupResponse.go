package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type NGSetupResponse struct {
	AMFName                []byte                  `lb:1,ub:150,mandatory,reject,valueExt`
	ServedGUAMIList        []ServedGUAMIItem       `lb:1,ub:maxnoofServedGUAMIs,mandatory,reject`
	RelativeAMFCapacity    int64                   `lb:0,ub:255,mandatory,ignore`
	PLMNSupportList        []PLMNSupportItem       `lb:1,ub:maxnoofPLMNs,mandatory,reject`
	CriticalityDiagnostics *CriticalityDiagnostics `optional,mandatory,ignore`
	UERetentionInformation *UERetentionInformation `optional,mandatory,ignore`
}

func (msg *NGSetupResponse) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		return
	}
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_NGSetup, Criticality_PresentReject, ies)
}
func (msg *NGSetupResponse) toIes() (ies []NgapMessageIE, err error) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_AMFName},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &OCTETSTRING{
			c:     aper.Constraint{Lb: 1, Ub: 150},
			ext:   true,
			Value: msg.AMFName,
		}})
	if len(msg.ServedGUAMIList) > 0 {
		tmp_ServedGUAMIList := Sequence[*ServedGUAMIItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofServedGUAMIs},
			ext: false,
		}
		for _, i := range msg.ServedGUAMIList {
			tmp_ServedGUAMIList.Value = append(tmp_ServedGUAMIList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ServedGUAMIList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_ServedGUAMIList,
		})
	} else {
		err = utils.WrapError("ServedGUAMIList is nil", err)
		return
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RelativeAMFCapacity},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 255},
			ext:   false,
			Value: aper.Integer(msg.RelativeAMFCapacity),
		}})
	if len(msg.PLMNSupportList) > 0 {
		tmp_PLMNSupportList := Sequence[*PLMNSupportItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPLMNs},
			ext: false,
		}
		for _, i := range msg.PLMNSupportList {
			tmp_PLMNSupportList.Value = append(tmp_PLMNSupportList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PLMNSupportList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_PLMNSupportList,
		})
	} else {
		err = utils.WrapError("PLMNSupportList is nil", err)
		return
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if msg.UERetentionInformation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UERetentionInformation},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UERetentionInformation,
		})
	}
	return
}
func (msg *NGSetupResponse) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("NGSetupResponse"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := NGSetupResponseDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*NgapMessageIE),
	}
	if _, err = aper.ReadSequenceOf[NgapMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	if _, ok := decoder.list[ProtocolIEID_AMFName]; !ok {
		err = fmt.Errorf("Mandatory field AMFName is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_AMFName},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_ServedGUAMIList]; !ok {
		err = fmt.Errorf("Mandatory field ServedGUAMIList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_ServedGUAMIList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_RelativeAMFCapacity]; !ok {
		err = fmt.Errorf("Mandatory field RelativeAMFCapacity is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_RelativeAMFCapacity},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_PLMNSupportList]; !ok {
		err = fmt.Errorf("Mandatory field PLMNSupportList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_PLMNSupportList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type NGSetupResponseDecoder struct {
	msg      *NGSetupResponse
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *NGSetupResponseDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_AMFName:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 150},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read AMFName", err)
			return
		}
		msg.AMFName = tmp.Value
	case ProtocolIEID_ServedGUAMIList:
		tmp := Sequence[*ServedGUAMIItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofServedGUAMIs},
			ext: false,
		}
		fn := func() *ServedGUAMIItem { return new(ServedGUAMIItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read ServedGUAMIList", err)
			return
		}
		msg.ServedGUAMIList = []ServedGUAMIItem{}
		for _, i := range tmp.Value {
			msg.ServedGUAMIList = append(msg.ServedGUAMIList, *i)
		}
	case ProtocolIEID_RelativeAMFCapacity:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 255},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RelativeAMFCapacity", err)
			return
		}
		msg.RelativeAMFCapacity = int64(tmp.Value)
	case ProtocolIEID_PLMNSupportList:
		tmp := Sequence[*PLMNSupportItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPLMNs},
			ext: false,
		}
		fn := func() *PLMNSupportItem { return new(PLMNSupportItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read PLMNSupportList", err)
			return
		}
		msg.PLMNSupportList = []PLMNSupportItem{}
		for _, i := range tmp.Value {
			msg.PLMNSupportList = append(msg.PLMNSupportList, *i)
		}
	case ProtocolIEID_CriticalityDiagnostics:
		var tmp CriticalityDiagnostics
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read CriticalityDiagnostics", err)
			return
		}
		msg.CriticalityDiagnostics = &tmp
	case ProtocolIEID_UERetentionInformation:
		var tmp UERetentionInformation
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read UERetentionInformation", err)
			return
		}
		msg.UERetentionInformation = &tmp
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
