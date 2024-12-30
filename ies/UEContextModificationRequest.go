package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type UEContextModificationRequest struct {
	AMFUENGAPID                                 int64
	RANUENGAPID                                 int64
	RANPagingPriority                           *int64
	SecurityKey                                 *[]byte
	IndexToRFSP                                 *int64
	UEAggregateMaximumBitRate                   *UEAggregateMaximumBitRate
	UESecurityCapabilities                      *UESecurityCapabilities
	CoreNetworkAssistanceInformationForInactive *CoreNetworkAssistanceInformationForInactive
	EmergencyFallbackIndicator                  *EmergencyFallbackIndicator
	NewAMFUENGAPID                              *int64
	RRCInactiveTransitionReportRequest          *RRCInactiveTransitionReportRequest
	NewGUAMI                                    *GUAMI
	CNAssistedRANTuning                         *CNAssistedRANTuning
}

func (msg *UEContextModificationRequest) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_UEContextModification, Criticality_PresentReject, msg.toIes())
}
func (msg *UEContextModificationRequest) toIes() (ies []NgapMessageIE) {
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
	if msg.SecurityKey != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SecurityKey},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value: &BITSTRING{
				Value: aper.BitString{
					Bytes:   *msg.SecurityKey,
					NumBits: uint64(len(*msg.SecurityKey))},
			}})
	}
	if msg.IndexToRFSP != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_IndexToRFSP},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 1, Ub: 256},
				ext:   true,
				Value: aper.Integer(*msg.IndexToRFSP),
			}})
	}
	if msg.UEAggregateMaximumBitRate != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UEAggregateMaximumBitRate},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UEAggregateMaximumBitRate,
		})
	}
	if msg.UESecurityCapabilities != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UESecurityCapabilities},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.UESecurityCapabilities,
		})
	}
	if msg.CoreNetworkAssistanceInformationForInactive != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CoreNetworkAssistanceInformationForInactive},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CoreNetworkAssistanceInformationForInactive,
		})
	}
	if msg.EmergencyFallbackIndicator != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_EmergencyFallbackIndicator},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.EmergencyFallbackIndicator,
		})
	}
	if msg.NewAMFUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NewAMFUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 1099511627775},
				ext:   false,
				Value: aper.Integer(*msg.NewAMFUENGAPID),
			}})
	}
	if msg.RRCInactiveTransitionReportRequest != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RRCInactiveTransitionReportRequest},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RRCInactiveTransitionReportRequest,
		})
	}
	if msg.NewGUAMI != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NewGUAMI},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.NewGUAMI,
		})
	}
	if msg.CNAssistedRANTuning != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CNAssistedRANTuning},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CNAssistedRANTuning,
		})
	}
	return
}
func (msg *UEContextModificationRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *UEContextModificationRequest) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_RANPagingPriority:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 256},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.RANPagingPriority = int64(tmp.Value)
	case ProtocolIEID_SecurityKey:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 256, Ub: 256},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.SecurityKey = tmp.Value.Bytes
	case ProtocolIEID_IndexToRFSP:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 256},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.IndexToRFSP = int64(tmp.Value)
	case ProtocolIEID_UEAggregateMaximumBitRate:
		var tmp UEAggregateMaximumBitRate
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UEAggregateMaximumBitRate = &tmp
	case ProtocolIEID_UESecurityCapabilities:
		var tmp UESecurityCapabilities
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UESecurityCapabilities = &tmp
	case ProtocolIEID_CoreNetworkAssistanceInformationForInactive:
		var tmp CoreNetworkAssistanceInformationForInactive
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.CoreNetworkAssistanceInformationForInactive = &tmp
	case ProtocolIEID_EmergencyFallbackIndicator:
		var tmp EmergencyFallbackIndicator
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.EmergencyFallbackIndicator = &tmp
	case ProtocolIEID_NewAMFUENGAPID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.NewAMFUENGAPID = int64(tmp.Value)
	case ProtocolIEID_RRCInactiveTransitionReportRequest:
		var tmp RRCInactiveTransitionReportRequest
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RRCInactiveTransitionReportRequest = &tmp
	case ProtocolIEID_NewGUAMI:
		var tmp GUAMI
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NewGUAMI = &tmp
	case ProtocolIEID_CNAssistedRANTuning:
		var tmp CNAssistedRANTuning
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.CNAssistedRANTuning = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
