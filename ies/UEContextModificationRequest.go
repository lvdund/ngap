package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type UEContextModificationRequest struct {
	AMFUENGAPID                                 *AMFUENGAPID                                 `,reject,mandatory`
	RANUENGAPID                                 *RANUENGAPID                                 `,reject,mandatory`
	RANPagingPriority                           *RANPagingPriority                           `,ignore,optional`
	SecurityKey                                 *SecurityKey                                 `,reject,optional`
	IndexToRFSP                                 *IndexToRFSP                                 `,ignore,optional`
	UEAggregateMaximumBitRate                   *UEAggregateMaximumBitRate                   `,ignore,optional`
	UESecurityCapabilities                      *UESecurityCapabilities                      `,reject,optional`
	CoreNetworkAssistanceInformationForInactive *CoreNetworkAssistanceInformationForInactive `,ignore,optional`
	EmergencyFallbackIndicator                  *EmergencyFallbackIndicator                  `,reject,optional`
	NewAMFUENGAPID                              *AMFUENGAPID                                 `,reject,optional`
	RRCInactiveTransitionReportRequest          *RRCInactiveTransitionReportRequest          `,ignore,optional`
	NewGUAMI                                    *GUAMI                                       `,reject,optional`
	CNAssistedRANTuning                         *CNAssistedRANTuning                         `,ignore,optional`
}

func (msg *UEContextModificationRequest) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_UEContextModification, Criticality_PresentReject, msg.toIes())
}
func (msg *UEContextModificationRequest) toIes() (ies []NgapMessageIE) {
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
	if msg.RANPagingPriority != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANPagingPriority},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RANPagingPriority})
	}
	if msg.SecurityKey != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SecurityKey},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.SecurityKey})
	}
	if msg.IndexToRFSP != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_IndexToRFSP},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.IndexToRFSP})
	}
	if msg.UEAggregateMaximumBitRate != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UEAggregateMaximumBitRate},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UEAggregateMaximumBitRate})
	}
	if msg.UESecurityCapabilities != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UESecurityCapabilities},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.UESecurityCapabilities})
	}
	if msg.CoreNetworkAssistanceInformationForInactive != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CoreNetworkAssistanceInformationForInactive},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CoreNetworkAssistanceInformationForInactive})
	}
	if msg.EmergencyFallbackIndicator != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_EmergencyFallbackIndicator},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.EmergencyFallbackIndicator})
	}
	if msg.NewAMFUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NewAMFUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.NewAMFUENGAPID})
	}
	if msg.RRCInactiveTransitionReportRequest != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RRCInactiveTransitionReportRequest},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RRCInactiveTransitionReportRequest})
	}
	if msg.NewGUAMI != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NewGUAMI},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.NewGUAMI})
	}
	if msg.CNAssistedRANTuning != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CNAssistedRANTuning},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CNAssistedRANTuning})
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
	case ProtocolIEID_RANPagingPriority:
		var tmp RANPagingPriority
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RANPagingPriority = &tmp
	case ProtocolIEID_SecurityKey:
		var tmp SecurityKey
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SecurityKey = &tmp
	case ProtocolIEID_IndexToRFSP:
		var tmp IndexToRFSP
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.IndexToRFSP = &tmp
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
		var tmp AMFUENGAPID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NewAMFUENGAPID = &tmp
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
