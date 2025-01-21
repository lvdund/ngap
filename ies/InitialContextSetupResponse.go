package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type InitialContextSetupResponse struct {
	AMFUENGAPID                               int64                                       `lb:0,ub:1099511627775,mandatory,ignore`
	RANUENGAPID                               int64                                       `lb:0,ub:4294967295,mandatory,ignore`
	PDUSessionResourceSetupListCxtRes         []PDUSessionResourceSetupItemCxtRes         `lb:1,ub:maxnoofPDUSessions,optional,mandatory,ignore`
	PDUSessionResourceFailedToSetupListCxtRes []PDUSessionResourceFailedToSetupItemCxtRes `lb:1,ub:maxnoofPDUSessions,optional,mandatory,ignore`
	CriticalityDiagnostics                    *CriticalityDiagnostics                     `optional,mandatory,ignore`
}

func (msg *InitialContextSetupResponse) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("InitialContextSetupResponse"), err)
		return
	}
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_InitialContextSetup, Criticality_PresentReject, ies)
}
func (msg *InitialContextSetupResponse) toIes() (ies []NgapMessageIE, err error) {
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
	if len(msg.PDUSessionResourceSetupListCxtRes) > 0 {
		tmp_PDUSessionResourceSetupListCxtRes := Sequence[*PDUSessionResourceSetupItemCxtRes]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range msg.PDUSessionResourceSetupListCxtRes {
			tmp_PDUSessionResourceSetupListCxtRes.Value = append(tmp_PDUSessionResourceSetupListCxtRes.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceSetupListCxtRes},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_PDUSessionResourceSetupListCxtRes,
		})
	}
	if len(msg.PDUSessionResourceFailedToSetupListCxtRes) > 0 {
		tmp_PDUSessionResourceFailedToSetupListCxtRes := Sequence[*PDUSessionResourceFailedToSetupItemCxtRes]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range msg.PDUSessionResourceFailedToSetupListCxtRes {
			tmp_PDUSessionResourceFailedToSetupListCxtRes.Value = append(tmp_PDUSessionResourceFailedToSetupListCxtRes.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceFailedToSetupListCxtRes},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_PDUSessionResourceFailedToSetupListCxtRes,
		})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	return
}
func (msg *InitialContextSetupResponse) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("InitialContextSetupResponse"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := InitialContextSetupResponseDecoder{
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
	return
}

type InitialContextSetupResponseDecoder struct {
	msg      *InitialContextSetupResponse
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *InitialContextSetupResponseDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_PDUSessionResourceSetupListCxtRes:
		tmp := Sequence[*PDUSessionResourceSetupItemCxtRes]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		fn := func() *PDUSessionResourceSetupItemCxtRes { return new(PDUSessionResourceSetupItemCxtRes) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read PDUSessionResourceSetupListCxtRes", err)
			return
		}
		msg.PDUSessionResourceSetupListCxtRes = []PDUSessionResourceSetupItemCxtRes{}
		for _, i := range tmp.Value {
			msg.PDUSessionResourceSetupListCxtRes = append(msg.PDUSessionResourceSetupListCxtRes, *i)
		}
	case ProtocolIEID_PDUSessionResourceFailedToSetupListCxtRes:
		tmp := Sequence[*PDUSessionResourceFailedToSetupItemCxtRes]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		fn := func() *PDUSessionResourceFailedToSetupItemCxtRes {
			return new(PDUSessionResourceFailedToSetupItemCxtRes)
		}
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read PDUSessionResourceFailedToSetupListCxtRes", err)
			return
		}
		msg.PDUSessionResourceFailedToSetupListCxtRes = []PDUSessionResourceFailedToSetupItemCxtRes{}
		for _, i := range tmp.Value {
			msg.PDUSessionResourceFailedToSetupListCxtRes = append(msg.PDUSessionResourceFailedToSetupListCxtRes, *i)
		}
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
