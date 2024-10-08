package ie

import "gen/aper"

type PDUSessionResourceReleasedItemPSFail struct {
	PDUSessionID                          PDUSessionID                               `False,`
	PathSwitchRequestUnsuccessfulTransfer aper.OctetString                           `False,`
	IEExtensions                          PDUSessionResourceReleasedItemPSFailExtIEs `False,OPTIONAL`
}
