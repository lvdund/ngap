package ie

import "gen/aper"

type PDUSessionResourceReleasedItemPSAck struct {
	PDUSessionID                          PDUSessionID                              `False,`
	PathSwitchRequestUnsuccessfulTransfer aper.OctetString                          `False,`
	IEExtensions                          PDUSessionResourceReleasedItemPSAckExtIEs `False,OPTIONAL`
}
