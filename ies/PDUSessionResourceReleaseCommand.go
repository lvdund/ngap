package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionResourceReleaseCommand struct {
	AMFUENGAPID                           int64                                   `lb:0,ub:1099511627775,mandatory,reject`
	RANUENGAPID                           int64                                   `lb:0,ub:4294967295,mandatory,reject`
	RANPagingPriority                     *int64                                  `lb:1,ub:256,optional,ignore`
	NASPDU                                []byte                                  `lb:0,ub:0,optional,ignore`
	PDUSessionResourceToReleaseListRelCmd []PDUSessionResourceToReleaseItemRelCmd `lb:1,ub:maxnoofPDUSessions,mandatory,reject`
}

func (msg *PDUSessionResourceReleaseCommand) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("PDUSessionResourceReleaseCommand"), err)
		return
	}
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_PDUSessionResourceRelease, Criticality_PresentReject, ies)
}
func (msg *PDUSessionResourceReleaseCommand) toIes() (ies []NgapMessageIE, err error) {
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
	if msg.RANPagingPriority != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANPagingPriority},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 1, Ub: 256},
				ext:   false,
				Value: aper.Integer(*msg.RANPagingPriority),
			}})
	}
	if msg.NASPDU != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NASPDU},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   false,
				Value: msg.NASPDU,
			}})
	}
	if len(msg.PDUSessionResourceToReleaseListRelCmd) > 0 {
		tmp_PDUSessionResourceToReleaseListRelCmd := Sequence[*PDUSessionResourceToReleaseItemRelCmd]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range msg.PDUSessionResourceToReleaseListRelCmd {
			tmp_PDUSessionResourceToReleaseListRelCmd.Value = append(tmp_PDUSessionResourceToReleaseListRelCmd.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceToReleaseListRelCmd},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_PDUSessionResourceToReleaseListRelCmd,
		})
	} else {
		err = utils.WrapError("PDUSessionResourceToReleaseListRelCmd is nil", err)
		return
	}
	return
}
func (msg *PDUSessionResourceReleaseCommand) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("PDUSessionResourceReleaseCommand"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := PDUSessionResourceReleaseCommandDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_PDUSessionResourceToReleaseListRelCmd]; !ok {
		err = fmt.Errorf("Mandatory field PDUSessionResourceToReleaseListRelCmd is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceToReleaseListRelCmd},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type PDUSessionResourceReleaseCommandDecoder struct {
	msg      *PDUSessionResourceReleaseCommand
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *PDUSessionResourceReleaseCommandDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_RANPagingPriority:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 256},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RANPagingPriority", err)
			return
		}
		msg.RANPagingPriority = (*int64)(&tmp.Value)
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
	case ProtocolIEID_PDUSessionResourceToReleaseListRelCmd:
		tmp := Sequence[*PDUSessionResourceToReleaseItemRelCmd]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		fn := func() *PDUSessionResourceToReleaseItemRelCmd { return new(PDUSessionResourceToReleaseItemRelCmd) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read PDUSessionResourceToReleaseListRelCmd", err)
			return
		}
		msg.PDUSessionResourceToReleaseListRelCmd = []PDUSessionResourceToReleaseItemRelCmd{}
		for _, i := range tmp.Value {
			msg.PDUSessionResourceToReleaseListRelCmd = append(msg.PDUSessionResourceToReleaseListRelCmd, *i)
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
