package ie

import "gen/aper"

type PDUSessionResourceFailedToSetupItemCxtFail struct {
	PDUSessionID                                PDUSessionID                                     `False,`
	PDUSessionResourceSetupUnsuccessfulTransfer aper.OctetString                                 `False,`
	IEExtensions                                PDUSessionResourceFailedToSetupItemCxtFailExtIEs `False,OPTIONAL`
}
