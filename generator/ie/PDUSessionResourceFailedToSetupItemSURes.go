package ie

import "gen/aper"

type PDUSessionResourceFailedToSetupItemSURes struct {
	PDUSessionID                                PDUSessionID                                   `False,`
	PDUSessionResourceSetupUnsuccessfulTransfer aper.OctetString                               `False,`
	IEExtensions                                PDUSessionResourceFailedToSetupItemSUResExtIEs `False,OPTIONAL`
}
