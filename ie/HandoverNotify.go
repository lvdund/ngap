package ie

type HandoverNotify struct {
MessageType	*MessageType
AmfUeNgapId	*AmfUeNgapId
RanUeNgapId	*RanUeNgapId
UserLocationInformation	*UserLocationInformation
NotifySourceNgRanNode	*[]byte
}
