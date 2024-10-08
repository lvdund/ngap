package ie

import "gen/aper"

type PDUSessionResourceToReleaseItemRelCmd struct {
	PDUSessionID                             PDUSessionID                                `False,`
	PDUSessionResourceReleaseCommandTransfer aper.OctetString                            `False,`
	IEExtensions                             PDUSessionResourceToReleaseItemRelCmdExtIEs `False,OPTIONAL`
}
