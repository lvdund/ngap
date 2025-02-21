package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type HandoverCommand struct {
	AMFUENGAPID                          int64                                  `lb:0,ub:1099511627775,mandatory,reject`
	RANUENGAPID                          int64                                  `lb:0,ub:4294967295,mandatory,reject`
	HandoverType                         HandoverType                           `mandatory,reject`
	NASSecurityParametersFromNGRAN       []byte                                 `lb:0,ub:0,conditional,reject`
	PDUSessionResourceHandoverList       []PDUSessionResourceHandoverItem       `lb:1,ub:maxnoofPDUSessions,optional,ignore`
	PDUSessionResourceToReleaseListHOCmd []PDUSessionResourceToReleaseItemHOCmd `lb:1,ub:maxnoofPDUSessions,optional,ignore`
	TargetToSourceTransparentContainer   []byte                                 `lb:0,ub:0,mandatory,reject`
	CriticalityDiagnostics               *CriticalityDiagnostics                `optional,ignore`
}

func (msg *HandoverCommand) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("HandoverCommand"), err)
		return
	}
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_HandoverPreparation, Criticality_PresentReject, ies)
}
func (msg *HandoverCommand) toIes() (ies []NgapMessageIE, err error) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext:   false,
			Value: aper.Integer(msg.AMFUENGAPID),
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
			Value: aper.Integer(msg.RANUENGAPID),
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_HandoverType},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.HandoverType,
	})
	if msg.NASSecurityParametersFromNGRAN != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NASSecurityParametersFromNGRAN},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   false,
				Value: msg.NASSecurityParametersFromNGRAN,
			}})
	}
	if len(msg.PDUSessionResourceHandoverList) > 0 {
		tmp_PDUSessionResourceHandoverList := Sequence[*PDUSessionResourceHandoverItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range msg.PDUSessionResourceHandoverList {
			tmp_PDUSessionResourceHandoverList.Value = append(tmp_PDUSessionResourceHandoverList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceHandoverList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_PDUSessionResourceHandoverList,
		})
	}
	if len(msg.PDUSessionResourceToReleaseListHOCmd) > 0 {
		tmp_PDUSessionResourceToReleaseListHOCmd := Sequence[*PDUSessionResourceToReleaseItemHOCmd]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range msg.PDUSessionResourceToReleaseListHOCmd {
			tmp_PDUSessionResourceToReleaseListHOCmd.Value = append(tmp_PDUSessionResourceToReleaseListHOCmd.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceToReleaseListHOCmd},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_PDUSessionResourceToReleaseListHOCmd,
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
func (msg *HandoverCommand) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("HandoverCommand"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := HandoverCommandDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*NgapMessageIE),
	}
	if _, err = aper.ReadSequenceOf[NgapMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	if _, ok := decoder.list[ProtocolIEID_AMFUENGAPID]; !ok {
		err = fmt.Errorf("Mandatory field AMFUENGAPID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
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
	if _, ok := decoder.list[ProtocolIEID_HandoverType]; !ok {
		err = fmt.Errorf("Mandatory field HandoverType is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_HandoverType},
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

type HandoverCommandDecoder struct {
	msg      *HandoverCommand
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *HandoverCommandDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_HandoverType:
		var tmp HandoverType
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read HandoverType", err)
			return
		}
		msg.HandoverType = tmp
	case ProtocolIEID_NASSecurityParametersFromNGRAN:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read NASSecurityParametersFromNGRAN", err)
			return
		}
		msg.NASSecurityParametersFromNGRAN = tmp.Value
	case ProtocolIEID_PDUSessionResourceHandoverList:
		tmp := Sequence[*PDUSessionResourceHandoverItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		fn := func() *PDUSessionResourceHandoverItem { return new(PDUSessionResourceHandoverItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read PDUSessionResourceHandoverList", err)
			return
		}
		msg.PDUSessionResourceHandoverList = []PDUSessionResourceHandoverItem{}
		for _, i := range tmp.Value {
			msg.PDUSessionResourceHandoverList = append(msg.PDUSessionResourceHandoverList, *i)
		}
	case ProtocolIEID_PDUSessionResourceToReleaseListHOCmd:
		tmp := Sequence[*PDUSessionResourceToReleaseItemHOCmd]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		fn := func() *PDUSessionResourceToReleaseItemHOCmd { return new(PDUSessionResourceToReleaseItemHOCmd) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read PDUSessionResourceToReleaseListHOCmd", err)
			return
		}
		msg.PDUSessionResourceToReleaseListHOCmd = []PDUSessionResourceToReleaseItemHOCmd{}
		for _, i := range tmp.Value {
			msg.PDUSessionResourceToReleaseListHOCmd = append(msg.PDUSessionResourceToReleaseListHOCmd, *i)
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
