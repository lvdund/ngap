package ie

import "gen/aper"

type PDUSessionResourceModifyItemModInd struct {
	PDUSessionID                               PDUSessionID                             `False,`
	PDUSessionResourceModifyIndicationTransfer aper.OctetString                         `False,`
	IEExtensions                               PDUSessionResourceModifyItemModIndExtIEs `False,OPTIONAL`
}
