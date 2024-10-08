package ie

import "gen/aper"

type PDUSessionResourceFailedToModifyItemModRes struct {
	PDUSessionID                                 PDUSessionID                                     `False,`
	PDUSessionResourceModifyUnsuccessfulTransfer aper.OctetString                                 `False,`
	IEExtensions                                 PDUSessionResourceFailedToModifyItemModResExtIEs `False,OPTIONAL`
}
