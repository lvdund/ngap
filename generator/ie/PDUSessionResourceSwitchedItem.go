package ie

import "gen/aper"

type PDUSessionResourceSwitchedItem struct {
	PDUSessionID                         PDUSessionID                         `False,`
	PathSwitchRequestAcknowledgeTransfer aper.OctetString                     `False,`
	IEExtensions                         PDUSessionResourceSwitchedItemExtIEs `False,OPTIONAL`
}
