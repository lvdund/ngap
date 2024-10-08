package ie

import "gen/aper"

type PDUSessionResourceHandoverItem struct {
	PDUSessionID            PDUSessionID                         `False,`
	HandoverCommandTransfer aper.OctetString                     `False,`
	IEExtensions            PDUSessionResourceHandoverItemExtIEs `False,OPTIONAL`
}
