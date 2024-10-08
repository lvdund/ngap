package ie

import "gen/aper"

type PDUSessionResourceFailedToSetupItemPSReq struct {
	PDUSessionID                         PDUSessionID                                   `False,`
	PathSwitchRequestSetupFailedTransfer aper.OctetString                               `False,`
	IEExtensions                         PDUSessionResourceFailedToSetupItemPSReqExtIEs `False,OPTIONAL`
}
