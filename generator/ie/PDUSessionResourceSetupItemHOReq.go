package ie

import "gen/aper"

type PDUSessionResourceSetupItemHOReq struct {
	PDUSessionID            PDUSessionID                           `False,`
	SNSSAI                  SNSSAI                                 `True,`
	HandoverRequestTransfer aper.OctetString                       `False,`
	IEExtensions            PDUSessionResourceSetupItemHOReqExtIEs `False,OPTIONAL`
}
