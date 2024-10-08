package ie

import "gen/aper"

type PDUSessionResourceSetupItemCxtReq struct {
	PDUSessionID                           PDUSessionID                            `False,`
	NASPDU                                 NASPDU                                  `False,OPTIONAL`
	SNSSAI                                 SNSSAI                                  `True,`
	PDUSessionResourceSetupRequestTransfer aper.OctetString                        `False,`
	IEExtensions                           PDUSessionResourceSetupItemCxtReqExtIEs `False,OPTIONAL`
}
