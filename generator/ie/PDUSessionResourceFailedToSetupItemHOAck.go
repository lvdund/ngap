package ie

import "gen/aper"

type PDUSessionResourceFailedToSetupItemHOAck struct {
	PDUSessionID                                   PDUSessionID                                   `False,`
	HandoverResourceAllocationUnsuccessfulTransfer aper.OctetString                               `False,`
	IEExtensions                                   PDUSessionResourceFailedToSetupItemHOAckExtIEs `False,OPTIONAL`
}
