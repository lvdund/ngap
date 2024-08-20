package ie

type UplinkNasTransport struct {
MessageType	*MessageType
AmfUeNgapId	*AmfUeNgapId
RanUeNgapId	*RanUeNgapId
NasPdu	*NasPdu
UserLocationInformation	*UserLocationInformation
WAgfIdentityInformation	*[]byte
TngfIdentityInformation	*[]byte
TwifIdentityInformation	*[]byte
}
