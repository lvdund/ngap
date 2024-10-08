package ie

import "gen/aper"

type PDUSessionResourceFailedToModifyItemModCfm struct {
	PDUSessionID                                           PDUSessionID                                     `False,`
	PDUSessionResourceModifyIndicationUnsuccessfulTransfer aper.OctetString                                 `False,`
	IEExtensions                                           PDUSessionResourceFailedToModifyItemModCfmExtIEs `False,OPTIONAL`
}
