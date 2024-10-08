package ie

import "gen/aper"

type PDUSessionResourceReleasedItemRelRes struct {
	PDUSessionID                              PDUSessionID                               `False,`
	PDUSessionResourceReleaseResponseTransfer aper.OctetString                           `False,`
	IEExtensions                              PDUSessionResourceReleasedItemRelResExtIEs `False,OPTIONAL`
}
