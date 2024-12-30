package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type HandoverCommand struct {
	AMFUENGAPID                          int64
	RANUENGAPID                          int64
	HandoverType                         HandoverType
	NASSecurityParametersFromNGRAN       []byte
	PDUSessionResourceHandoverList       *[]PDUSessionResourceHandoverItem
	PDUSessionResourceToReleaseListHOCmd *[]PDUSessionResourceToReleaseItemHOCmd
	TargetToSourceTransparentContainer   []byte
	CriticalityDiagnostics               *CriticalityDiagnostics
}

func (msg *HandoverCommand) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_HandoverPreparation, Criticality_PresentReject, msg.toIes())
}
func (msg *HandoverCommand) toIes() (ies []NgapMessageIE) {
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
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_NASSecurityParametersFromNGRAN},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &OCTETSTRING{
			c:     aper.Constraint{Lb: 0, Ub: 0},
			ext:   false,
			Value: msg.NASSecurityParametersFromNGRAN,
		}})
	if msg.PDUSessionResourceHandoverList != nil {
		tmp_PDUSessionResourceHandoverList := Sequence[*PDUSessionResourceHandoverItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range *msg.PDUSessionResourceHandoverList {
			tmp_PDUSessionResourceHandoverList.Value = append(tmp_PDUSessionResourceHandoverList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceHandoverList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_PDUSessionResourceHandoverList,
		})
	}
	if msg.PDUSessionResourceToReleaseListHOCmd != nil {
		tmp_PDUSessionResourceToReleaseListHOCmd := Sequence[*PDUSessionResourceToReleaseItemHOCmd]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range *msg.PDUSessionResourceToReleaseListHOCmd {
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
func (msg *HandoverCommand) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *HandoverCommand) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_HandoverType:
		var tmp HandoverType
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.HandoverType = tmp
	case ProtocolIEID_NASSecurityParametersFromNGRAN:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NASSecurityParametersFromNGRAN = tmp.Value
	case ProtocolIEID_PDUSessionResourceHandoverList:
		tmp := Sequence[*PDUSessionResourceHandoverItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceHandoverList = &[]PDUSessionResourceHandoverItem{}
		for _, i := range tmp.Value {
			*msg.PDUSessionResourceHandoverList = append(*msg.PDUSessionResourceHandoverList, *i)
		}
	case ProtocolIEID_PDUSessionResourceToReleaseListHOCmd:
		tmp := Sequence[*PDUSessionResourceToReleaseItemHOCmd]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceToReleaseListHOCmd = &[]PDUSessionResourceToReleaseItemHOCmd{}
		for _, i := range tmp.Value {
			*msg.PDUSessionResourceToReleaseListHOCmd = append(*msg.PDUSessionResourceToReleaseListHOCmd, *i)
		}
	case ProtocolIEID_TargetToSourceTransparentContainer:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.TargetToSourceTransparentContainer = tmp.Value
	case ProtocolIEID_CriticalityDiagnostics:
		var tmp CriticalityDiagnostics
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.CriticalityDiagnostics = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
