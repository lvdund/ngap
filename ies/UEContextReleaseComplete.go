package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UEContextReleaseComplete struct {
	AMFUENGAPID                                int64                                       `lb:0,ub:1099511627775,mandatory,ignore`
	RANUENGAPID                                int64                                       `lb:0,ub:4294967295,mandatory,ignore`
	UserLocationInformation                    *UserLocationInformation                    `optional,mandatory,ignore`
	InfoOnRecommendedCellsAndRANNodesForPaging *InfoOnRecommendedCellsAndRANNodesForPaging `optional,mandatory,ignore`
	PDUSessionResourceListCxtRelCpl            []PDUSessionResourceItemCxtRelCpl           `lb:1,ub:maxnoofPDUSessions,optional,mandatory,reject`
	CriticalityDiagnostics                     *CriticalityDiagnostics                     `optional,mandatory,ignore`
}

func (msg *UEContextReleaseComplete) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("UEContextReleaseComplete"), err)
		return
	}
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_UEContextRelease, Criticality_PresentReject, ies)
}
func (msg *UEContextReleaseComplete) toIes() (ies []NgapMessageIE, err error) {
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
	if msg.UserLocationInformation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UserLocationInformation},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UserLocationInformation,
		})
	}
	if msg.InfoOnRecommendedCellsAndRANNodesForPaging != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_InfoOnRecommendedCellsAndRANNodesForPaging},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.InfoOnRecommendedCellsAndRANNodesForPaging,
		})
	}
	if len(msg.PDUSessionResourceListCxtRelCpl) > 0 {
		tmp_PDUSessionResourceListCxtRelCpl := Sequence[*PDUSessionResourceItemCxtRelCpl]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range msg.PDUSessionResourceListCxtRelCpl {
			tmp_PDUSessionResourceListCxtRelCpl.Value = append(tmp_PDUSessionResourceListCxtRelCpl.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceListCxtRelCpl},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_PDUSessionResourceListCxtRelCpl,
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
func (msg *UEContextReleaseComplete) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("UEContextReleaseComplete"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := UEContextReleaseCompleteDecoder{
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

type UEContextReleaseCompleteDecoder struct {
	msg      *UEContextReleaseComplete
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *UEContextReleaseCompleteDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_UserLocationInformation:
		var tmp UserLocationInformation
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read UserLocationInformation", err)
			return
		}
		msg.UserLocationInformation = &tmp
	case ProtocolIEID_InfoOnRecommendedCellsAndRANNodesForPaging:
		var tmp InfoOnRecommendedCellsAndRANNodesForPaging
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read InfoOnRecommendedCellsAndRANNodesForPaging", err)
			return
		}
		msg.InfoOnRecommendedCellsAndRANNodesForPaging = &tmp
	case ProtocolIEID_PDUSessionResourceListCxtRelCpl:
		tmp := Sequence[*PDUSessionResourceItemCxtRelCpl]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		fn := func() *PDUSessionResourceItemCxtRelCpl { return new(PDUSessionResourceItemCxtRelCpl) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read PDUSessionResourceListCxtRelCpl", err)
			return
		}
		msg.PDUSessionResourceListCxtRelCpl = []PDUSessionResourceItemCxtRelCpl{}
		for _, i := range tmp.Value {
			msg.PDUSessionResourceListCxtRelCpl = append(msg.PDUSessionResourceListCxtRelCpl, *i)
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
