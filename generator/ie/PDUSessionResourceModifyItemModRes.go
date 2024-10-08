package ie

import "gen/aper"

type PDUSessionResourceModifyItemModRes struct {
	PDUSessionID                             PDUSessionID                             `False,`
	PDUSessionResourceModifyResponseTransfer aper.OctetString                         `False,`
	IEExtensions                             PDUSessionResourceModifyItemModResExtIEs `False,OPTIONAL`
}
