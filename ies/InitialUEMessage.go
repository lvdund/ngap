package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type InitialUEMessage struct {
	RANUENGAPID                         int64                                `lb:0,ub:4294967295,mandatory,reject`
	NASPDU                              []byte                               `lb:0,ub:0,mandatory,reject`
	UserLocationInformation             UserLocationInformation              `mandatory,reject`
	RRCEstablishmentCause               RRCEstablishmentCause                `mandatory,ignore`
	FiveGSTMSI                          *FiveGSTMSI                          `optional,mandatory,reject`
	AMFSetID                            []byte                               `lb:10,ub:10,optional,mandatory,ignore`
	UEContextRequest                    *UEContextRequest                    `optional,mandatory,ignore`
	AllowedNSSAI                        []AllowedNSSAIItem                   `lb:1,ub:maxnoofAllowedSNSSAIs,optional,mandatory,reject`
	SourceToTargetAMFInformationReroute *SourceToTargetAMFInformationReroute `optional,mandatory,ignore`
	SelectedPLMNIdentity                []byte                               `lb:3,ub:3,optional,mandatory,ignore`
}

func (msg *InitialUEMessage) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		return
	}
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_InitialUEMessage, Criticality_PresentIgnore, ies)
}
func (msg *InitialUEMessage) toIes() (ies []NgapMessageIE, err error) {
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
		Id:          ProtocolIEID{Value: ProtocolIEID_NASPDU},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &OCTETSTRING{
			c:     aper.Constraint{Lb: 0, Ub: 0},
			ext:   false,
			Value: msg.NASPDU,
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_UserLocationInformation},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.UserLocationInformation,
	})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RRCEstablishmentCause},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &msg.RRCEstablishmentCause,
	})
	if msg.FiveGSTMSI != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_FiveGSTMSI},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.FiveGSTMSI,
		})
	}
	if msg.AMFSetID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFSetID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &BITSTRING{
				c:   aper.Constraint{Lb: 10, Ub: 10},
				ext: false,
				Value: aper.BitString{
					Bytes: msg.AMFSetID},
			}})
	}
	if msg.UEContextRequest != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UEContextRequest},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UEContextRequest,
		})
	}
	if len(msg.AllowedNSSAI) > 0 {
		tmp_AllowedNSSAI := Sequence[*AllowedNSSAIItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofAllowedSNSSAIs},
			ext: false,
		}
		for _, i := range msg.AllowedNSSAI {
			tmp_AllowedNSSAI.Value = append(tmp_AllowedNSSAI.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AllowedNSSAI},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_AllowedNSSAI,
		})
	}
	if msg.SourceToTargetAMFInformationReroute != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SourceToTargetAMFInformationReroute},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.SourceToTargetAMFInformationReroute,
		})
	}
	if msg.SelectedPLMNIdentity != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SelectedPLMNIdentity},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 3, Ub: 3},
				ext:   false,
				Value: msg.SelectedPLMNIdentity,
			}})
	}
	return
}
func (msg *InitialUEMessage) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("InitialUEMessage"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := InitialUEMessageDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_NASPDU]; !ok {
		err = fmt.Errorf("Mandatory field NASPDU is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_NASPDU},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_UserLocationInformation]; !ok {
		err = fmt.Errorf("Mandatory field UserLocationInformation is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_UserLocationInformation},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_RRCEstablishmentCause]; !ok {
		err = fmt.Errorf("Mandatory field RRCEstablishmentCause is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_RRCEstablishmentCause},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type InitialUEMessageDecoder struct {
	msg      *InitialUEMessage
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *InitialUEMessageDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_NASPDU:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read NASPDU", err)
			return
		}
		msg.NASPDU = tmp.Value
	case ProtocolIEID_UserLocationInformation:
		var tmp UserLocationInformation
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read UserLocationInformation", err)
			return
		}
		msg.UserLocationInformation = tmp
	case ProtocolIEID_RRCEstablishmentCause:
		var tmp RRCEstablishmentCause
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RRCEstablishmentCause", err)
			return
		}
		msg.RRCEstablishmentCause = tmp
	case ProtocolIEID_FiveGSTMSI:
		var tmp FiveGSTMSI
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read FiveGSTMSI", err)
			return
		}
		msg.FiveGSTMSI = &tmp
	case ProtocolIEID_AMFSetID:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 10, Ub: 10},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read AMFSetID", err)
			return
		}
		msg.AMFSetID = tmp.Value.Bytes
	case ProtocolIEID_UEContextRequest:
		var tmp UEContextRequest
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read UEContextRequest", err)
			return
		}
		msg.UEContextRequest = &tmp
	case ProtocolIEID_AllowedNSSAI:
		tmp := Sequence[*AllowedNSSAIItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofAllowedSNSSAIs},
			ext: false,
		}
		fn := func() *AllowedNSSAIItem { return new(AllowedNSSAIItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read AllowedNSSAI", err)
			return
		}
		msg.AllowedNSSAI = []AllowedNSSAIItem{}
		for _, i := range tmp.Value {
			msg.AllowedNSSAI = append(msg.AllowedNSSAI, *i)
		}
	case ProtocolIEID_SourceToTargetAMFInformationReroute:
		var tmp SourceToTargetAMFInformationReroute
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SourceToTargetAMFInformationReroute", err)
			return
		}
		msg.SourceToTargetAMFInformationReroute = &tmp
	case ProtocolIEID_SelectedPLMNIdentity:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 3, Ub: 3},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SelectedPLMNIdentity", err)
			return
		}
		msg.SelectedPLMNIdentity = tmp.Value
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
