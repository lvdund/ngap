package ie

import "ngap/aper"

type PduSessionResourceReleaseCommand struct {
	MessageType                     MessageType                       //`bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                     AmfUeNgapId                       //`bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                     RanUeNgapId                       //`bitstring:"sizeLB:0,sizeUB:150"`
	RanPagingPriority               RanPagingPriority                 //`bitstring:"sizeLB:0,sizeUB:150"`
	NasPdu                          NasPdu                            //`bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceToReleaseList []PduSessionResourceToReleaseItem //`bitstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionResourceToReleaseItem struct {
	PduSessionId                             PduSessionId     //`bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceReleaseCommandTransfer aper.OctetString //`octetstring:"sizeLB:0,sizeUB:150"`
}
