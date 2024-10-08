package ie

import "gen/aper"

type PDUSessionResourceToReleaseItemHOCmd struct {
	PDUSessionID                            PDUSessionID                               `False,`
	HandoverPreparationUnsuccessfulTransfer aper.OctetString                           `False,`
	IEExtensions                            PDUSessionResourceToReleaseItemHOCmdExtIEs `False,OPTIONAL`
}
