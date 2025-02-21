package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type CellTrafficTrace struct {
	AMFUENGAPID                    int64          `lb:0,ub:1099511627775,mandatory,reject`
	RANUENGAPID                    int64          `lb:0,ub:4294967295,mandatory,reject`
	NGRANTraceID                   []byte         `lb:8,ub:8,mandatory,ignore`
	NGRANCGI                       NGRANCGI       `mandatory,ignore`
	TraceCollectionEntityIPAddress aper.BitString `lb:1,ub:160,mandatory,ignore,valueExt`
}

func (msg *CellTrafficTrace) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("CellTrafficTrace"), err)
		return
	}
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_CellTrafficTrace, Criticality_PresentIgnore, ies)
}
func (msg *CellTrafficTrace) toIes() (ies []NgapMessageIE, err error) {
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
		Id:          ProtocolIEID{Value: ProtocolIEID_NGRANTraceID},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value: &OCTETSTRING{
			c:     aper.Constraint{Lb: 8, Ub: 8},
			ext:   false,
			Value: msg.NGRANTraceID,
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_NGRANCGI},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &msg.NGRANCGI,
	})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TraceCollectionEntityIPAddress},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value: &BITSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 160},
			ext: true,
			Value: aper.BitString{
				Bytes: msg.TraceCollectionEntityIPAddress.Bytes, NumBits: msg.TraceCollectionEntityIPAddress.NumBits},
		}})
	return
}
func (msg *CellTrafficTrace) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("CellTrafficTrace"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := CellTrafficTraceDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_NGRANTraceID]; !ok {
		err = fmt.Errorf("Mandatory field NGRANTraceID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_NGRANTraceID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_NGRANCGI]; !ok {
		err = fmt.Errorf("Mandatory field NGRANCGI is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_NGRANCGI},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_TraceCollectionEntityIPAddress]; !ok {
		err = fmt.Errorf("Mandatory field TraceCollectionEntityIPAddress is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_TraceCollectionEntityIPAddress},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type CellTrafficTraceDecoder struct {
	msg      *CellTrafficTrace
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *CellTrafficTraceDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_NGRANTraceID:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 8, Ub: 8},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read NGRANTraceID", err)
			return
		}
		msg.NGRANTraceID = tmp.Value
	case ProtocolIEID_NGRANCGI:
		var tmp NGRANCGI
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read NGRANCGI", err)
			return
		}
		msg.NGRANCGI = tmp
	case ProtocolIEID_TraceCollectionEntityIPAddress:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 160},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read TraceCollectionEntityIPAddress", err)
			return
		}
		msg.TraceCollectionEntityIPAddress = aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
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
