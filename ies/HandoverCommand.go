package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type HandoverCommand struct {
	AMFUENGAPID                          *AMFUENGAPID                          `,reject,mandatory`
	RANUENGAPID                          *RANUENGAPID                          `,reject,mandatory`
	HandoverType                         *HandoverType                         `,reject,mandatory`
	NASSecurityParametersFromNGRAN       *NASSecurityParametersFromNGRAN       `,reject,conditional`
	PDUSessionResourceHandoverList       *PDUSessionResourceHandoverList       `,ignore,optional`
	PDUSessionResourceToReleaseListHOCmd *PDUSessionResourceToReleaseListHOCmd `,ignore,optional`
	TargetToSourceTransparentContainer   *TargetToSourceTransparentContainer   `,reject,mandatory`
	CriticalityDiagnostics               *CriticalityDiagnostics               `,ignore,optional`
}

func (msg *HandoverCommand) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_HandoverPreparation, Criticality_PresentReject, msg.toIes())
}
func (msg *HandoverCommand) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.AMFUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.AMFUENGAPID})
	}
	if msg.RANUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.RANUENGAPID})
	}
	if msg.HandoverType != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_HandoverType},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.HandoverType})
	}
	if msg.NASSecurityParametersFromNGRAN != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NASSecurityParametersFromNGRAN},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.NASSecurityParametersFromNGRAN})
	}
	if msg.PDUSessionResourceHandoverList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceHandoverList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.PDUSessionResourceHandoverList})
	}
	if msg.PDUSessionResourceToReleaseListHOCmd != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceToReleaseListHOCmd},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.PDUSessionResourceToReleaseListHOCmd})
	}
	if msg.TargetToSourceTransparentContainer != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TargetToSourceTransparentContainer},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.TargetToSourceTransparentContainer})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics})
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
	case ProtocolIEID_AMFUENGAPID:
		var tmp AMFUENGAPID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFUENGAPID = &tmp
	case ProtocolIEID_RANUENGAPID:
		var tmp RANUENGAPID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RANUENGAPID = &tmp
	case ProtocolIEID_HandoverType:
		var tmp HandoverType
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.HandoverType = &tmp
	case ProtocolIEID_NASSecurityParametersFromNGRAN:
		var tmp NASSecurityParametersFromNGRAN
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NASSecurityParametersFromNGRAN = &tmp
	case ProtocolIEID_PDUSessionResourceHandoverList:
		var tmp PDUSessionResourceHandoverList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceHandoverList = &tmp
	case ProtocolIEID_PDUSessionResourceToReleaseListHOCmd:
		var tmp PDUSessionResourceToReleaseListHOCmd
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceToReleaseListHOCmd = &tmp
	case ProtocolIEID_TargetToSourceTransparentContainer:
		var tmp TargetToSourceTransparentContainer
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.TargetToSourceTransparentContainer = &tmp
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
