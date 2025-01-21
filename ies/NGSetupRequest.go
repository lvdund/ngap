package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type NGSetupRequest struct {
	GlobalRANNodeID        GlobalRANNodeID         `mandatory,reject`
	RANNodeName            []byte                  `lb:1,ub:150,optional,mandatory,ignore,valueExt`
	SupportedTAList        []SupportedTAItem       `lb:1,ub:maxnoofTACs,mandatory,reject`
	DefaultPagingDRX       PagingDRX               `mandatory,ignore`
	UERetentionInformation *UERetentionInformation `optional,mandatory,ignore`
}

func (msg *NGSetupRequest) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("NGSetupRequest"), err)
		return
	}
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_NGSetup, Criticality_PresentReject, ies)
}
func (msg *NGSetupRequest) toIes() (ies []NgapMessageIE, err error) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_GlobalRANNodeID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.GlobalRANNodeID,
	})
	if msg.RANNodeName != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANNodeName},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 1, Ub: 150},
				ext:   true,
				Value: msg.RANNodeName,
			}})
	}
	if len(msg.SupportedTAList) > 0 {
		tmp_SupportedTAList := Sequence[*SupportedTAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTACs},
			ext: false,
		}
		for _, i := range msg.SupportedTAList {
			tmp_SupportedTAList.Value = append(tmp_SupportedTAList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SupportedTAList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_SupportedTAList,
		})
	} else {
		err = utils.WrapError("SupportedTAList is nil", err)
		return
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_DefaultPagingDRX},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &msg.DefaultPagingDRX,
	})
	if msg.UERetentionInformation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UERetentionInformation},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UERetentionInformation,
		})
	}
	return
}
func (msg *NGSetupRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("NGSetupRequest"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := NGSetupRequestDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*NgapMessageIE),
	}
	if _, err = aper.ReadSequenceOf[NgapMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	if _, ok := decoder.list[ProtocolIEID_GlobalRANNodeID]; !ok {
		err = fmt.Errorf("Mandatory field GlobalRANNodeID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_GlobalRANNodeID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_SupportedTAList]; !ok {
		err = fmt.Errorf("Mandatory field SupportedTAList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_SupportedTAList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_DefaultPagingDRX]; !ok {
		err = fmt.Errorf("Mandatory field DefaultPagingDRX is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_DefaultPagingDRX},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type NGSetupRequestDecoder struct {
	msg      *NGSetupRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *NGSetupRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_GlobalRANNodeID:
		var tmp GlobalRANNodeID
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read GlobalRANNodeID", err)
			return
		}
		msg.GlobalRANNodeID = tmp
	case ProtocolIEID_RANNodeName:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 150},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RANNodeName", err)
			return
		}
		msg.RANNodeName = tmp.Value
	case ProtocolIEID_SupportedTAList:
		tmp := Sequence[*SupportedTAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTACs},
			ext: false,
		}
		fn := func() *SupportedTAItem { return new(SupportedTAItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read SupportedTAList", err)
			return
		}
		msg.SupportedTAList = []SupportedTAItem{}
		for _, i := range tmp.Value {
			msg.SupportedTAList = append(msg.SupportedTAList, *i)
		}
	case ProtocolIEID_DefaultPagingDRX:
		var tmp PagingDRX
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read DefaultPagingDRX", err)
			return
		}
		msg.DefaultPagingDRX = tmp
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
