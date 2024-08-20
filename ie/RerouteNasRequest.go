package ie

type RerouteNasRequest struct {
MessageType	*MessageType
RanUeNgapId	*RanUeNgapId
AmfUeNgapId	*AmfUeNgapId
NgapMessage	*[]byte
AmfSetId	*AmfSetId
AllowedNssai	*AllowedNssai
SourceToTargetAmfInformationReroute	*SourceToTargetAmfInformationReroute
}
