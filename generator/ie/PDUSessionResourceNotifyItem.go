package ie

import "gen/aper"

type PDUSessionResourceNotifyItem struct {
	PDUSessionID                     PDUSessionID                       `False,`
	PDUSessionResourceNotifyTransfer aper.OctetString                   `False,`
	IEExtensions                     PDUSessionResourceNotifyItemExtIEs `False,OPTIONAL`
}
