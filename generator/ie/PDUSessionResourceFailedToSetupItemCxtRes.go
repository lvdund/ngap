package ie

import "gen/aper"

type PDUSessionResourceFailedToSetupItemCxtRes struct {
	PDUSessionID                                PDUSessionID                                    `False,`
	PDUSessionResourceSetupUnsuccessfulTransfer aper.OctetString                                `False,`
	IEExtensions                                PDUSessionResourceFailedToSetupItemCxtResExtIEs `False,OPTIONAL`
}
