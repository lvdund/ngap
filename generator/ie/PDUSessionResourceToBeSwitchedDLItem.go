package ie

import "gen/aper"

type PDUSessionResourceToBeSwitchedDLItem struct {
	PDUSessionID              PDUSessionID                               `False,`
	PathSwitchRequestTransfer aper.OctetString                           `False,`
	IEExtensions              PDUSessionResourceToBeSwitchedDLItemExtIEs `False,OPTIONAL`
}
