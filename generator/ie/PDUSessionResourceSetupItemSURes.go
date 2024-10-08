package ie

import "gen/aper"

type PDUSessionResourceSetupItemSURes struct {
	PDUSessionID                            PDUSessionID                           `False,`
	PDUSessionResourceSetupResponseTransfer aper.OctetString                       `False,`
	IEExtensions                            PDUSessionResourceSetupItemSUResExtIEs `False,OPTIONAL`
}
