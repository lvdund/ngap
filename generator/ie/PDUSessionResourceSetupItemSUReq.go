package ie

import "gen/aper"

type PDUSessionResourceSetupItemSUReq struct {
	PDUSessionID                           PDUSessionID                           `False,`
	PDUSessionNASPDU                       NASPDU                                 `False,OPTIONAL`
	SNSSAI                                 SNSSAI                                 `True,`
	PDUSessionResourceSetupRequestTransfer aper.OctetString                       `False,`
	IEExtensions                           PDUSessionResourceSetupItemSUReqExtIEs `False,OPTIONAL`
}
