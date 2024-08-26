package ie

import "ngap/aper"

type PduSessionResourceModifyRequest struct {
	MessageType                         MessageType                           `bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                         AmfUeNgapId                           `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                         RanUeNgapId                           `bitstring:"sizeLB:0,sizeUB:150"`
	RanPagingPriority                   RanPagingPriority                     `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceModifyRequestList []PduSessionResourceModifyRequestItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionResourceModifyRequestItem struct {
	PduSessionId                            PduSessionId                `bitstring:"sizeLB:0,sizeUB:150"`
	NasPdu                                  NasPdu                      `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceModifyRequestTransfer aper.OctetString            `octetstring:"sizeLB:0,sizeUB:150"`
	SNssai                                  SNssai                      `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionExpectedUeActivityBehaviour   ExpectedUeActivityBehaviour `bitstring:"sizeLB:0,sizeUB:150"`
}
