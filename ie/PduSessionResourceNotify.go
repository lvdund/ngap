package ie

import "ngap/aper"

type PduSessionResourceNotify struct {
MessageType	MessageType	//`bitstring:"sizeLB:0,sizeUB:150"`
AmfUeNgapId	AmfUeNgapId	//`bitstring:"sizeLB:0,sizeUB:150"`
RanUeNgapId	RanUeNgapId	//`bitstring:"sizeLB:0,sizeUB:150"`
PduSessionResourceNotifyList	[]PduSessionResourceNotifyItem	//`bitstring:"sizeLB:0,sizeUB:150"`
PduSessionResourceReleasedList	[]PduSessionResourceReleasedItem	//`bitstring:"sizeLB:0,sizeUB:150"`
UserLocationInformation	UserLocationInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionResourceNotifyItem struct {
PduSessionId	PduSessionId	//`bitstring:"sizeLB:0,sizeUB:150"`
PduSessionResourceNotifyTransfer	aper.OctetString	//`octetstring:"sizeLB:0,sizeUB:150"`
}
