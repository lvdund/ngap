package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type LocationReport struct {
	AMFUENGAPID                    int64
	RANUENGAPID                    int64
	UserLocationInformation        UserLocationInformation
	UEPresenceInAreaOfInterestList *[]UEPresenceInAreaOfInterestItem
	LocationReportingRequestType   LocationReportingRequestType
}

func (msg *LocationReport) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_LocationReport, Criticality_PresentIgnore, msg.toIes())
}
func (msg *LocationReport) toIes() (ies []NgapMessageIE) {
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
	if msg.UEPresenceInAreaOfInterestList != nil {
		tmp_UEPresenceInAreaOfInterestList := Sequence[*UEPresenceInAreaOfInterestItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofAoI},
			ext: false,
		}
		for _, i := range *msg.UEPresenceInAreaOfInterestList {
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
func (msg *LocationReport) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *LocationReport) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_UserLocationInformation:
		var tmp UserLocationInformation
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UserLocationInformation = tmp
	case ProtocolIEID_UEPresenceInAreaOfInterestList:
		tmp := Sequence[*UEPresenceInAreaOfInterestItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofAoI},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UEPresenceInAreaOfInterestList = &[]UEPresenceInAreaOfInterestItem{}
		for _, i := range tmp.Value {
			*msg.UEPresenceInAreaOfInterestList = append(*msg.UEPresenceInAreaOfInterestList, *i)
		}
	case ProtocolIEID_LocationReportingRequestType:
		var tmp LocationReportingRequestType
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.LocationReportingRequestType = tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
