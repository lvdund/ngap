package ie

import "gen/aper"

type PDUSessionResourceModifyItemModReq struct {
	PDUSessionID                            PDUSessionID                             `False,`
	NASPDU                                  NASPDU                                   `False,OPTIONAL`
	PDUSessionResourceModifyRequestTransfer aper.OctetString                         `False,`
	IEExtensions                            PDUSessionResourceModifyItemModReqExtIEs `False,OPTIONAL`
}
