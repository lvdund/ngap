package ie

import "gen/aper"

type PDUSessionResourceSetupItemCxtRes struct {
	PDUSessionID                            PDUSessionID                            `False,`
	PDUSessionResourceSetupResponseTransfer aper.OctetString                        `False,`
	IEExtensions                            PDUSessionResourceSetupItemCxtResExtIEs `False,OPTIONAL`
}
