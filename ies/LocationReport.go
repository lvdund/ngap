package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type LocationReport struct {
	AMFUENGAPID                    int64                            `lb:0,ub:1099511627775,mandatory,reject`
	RANUENGAPID                    int64                            `lb:0,ub:4294967295,mandatory,reject`
	UserLocationInformation        UserLocationInformation          `mandatory,ignore`
	UEPresenceInAreaOfInterestList []UEPresenceInAreaOfInterestItem `lb:1,ub:maxnoofAoI,optional,mandatory,ignore`
	LocationReportingRequestType   LocationReportingRequestType     `mandatory,ignore`
}

func (msg *LocationReport) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("LocationReport"), err)
		return
	}
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_LocationReport, Criticality_PresentIgnore, ies)
}
func (msg *LocationReport) toIes() (ies []NgapMessageIE, err error) {
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
		Id:          ProtocolIEID{Value: ProtocolIEID_UserLocationInformation},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &msg.UserLocationInformation,
	})
	if len(msg.UEPresenceInAreaOfInterestList) > 0 {
		tmp_UEPresenceInAreaOfInterestList := Sequence[*UEPresenceInAreaOfInterestItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofAoI},
			ext: false,
		}
		for _, i := range msg.UEPresenceInAreaOfInterestList {
			tmp_UEPresenceInAreaOfInterestList.Value = append(tmp_UEPresenceInAreaOfInterestList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UEPresenceInAreaOfInterestList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_UEPresenceInAreaOfInterestList,
		})
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_LocationReportingRequestType},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &msg.LocationReportingRequestType,
	})
	return
}
func (msg *LocationReport) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("LocationReport"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := LocationReportDecoder{
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
	if _, ok := decoder.list[ProtocolIEID_UserLocationInformation]; !ok {
		err = fmt.Errorf("Mandatory field UserLocationInformation is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_UserLocationInformation},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_LocationReportingRequestType]; !ok {
		err = fmt.Errorf("Mandatory field LocationReportingRequestType is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentIgnore},
			IEID:          ProtocolIEID{Value: ProtocolIEID_LocationReportingRequestType},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type LocationReportDecoder struct {
	msg      *LocationReport
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *LocationReportDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
		msg.UserLocationInformation = tmp
	case ProtocolIEID_UEPresenceInAreaOfInterestList:
		tmp := Sequence[*UEPresenceInAreaOfInterestItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofAoI},
			ext: false,
		}
		fn := func() *UEPresenceInAreaOfInterestItem { return new(UEPresenceInAreaOfInterestItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read UEPresenceInAreaOfInterestList", err)
			return
		}
		msg.UEPresenceInAreaOfInterestList = []UEPresenceInAreaOfInterestItem{}
		for _, i := range tmp.Value {
			msg.UEPresenceInAreaOfInterestList = append(msg.UEPresenceInAreaOfInterestList, *i)
		}
	case ProtocolIEID_LocationReportingRequestType:
		var tmp LocationReportingRequestType
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read LocationReportingRequestType", err)
			return
		}
		msg.LocationReportingRequestType = tmp
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
