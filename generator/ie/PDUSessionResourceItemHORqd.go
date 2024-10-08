package ie

import "gen/aper"

type PDUSessionResourceItemHORqd struct {
	PDUSessionID             PDUSessionID                      `False,`
	HandoverRequiredTransfer aper.OctetString                  `False,`
	IEExtensions             PDUSessionResourceItemHORqdExtIEs `False,OPTIONAL`
}
