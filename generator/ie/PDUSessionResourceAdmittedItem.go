package ie

import "gen/aper"

type PDUSessionResourceAdmittedItem struct {
	PDUSessionID                       PDUSessionID                         `False,`
	HandoverRequestAcknowledgeTransfer aper.OctetString                     `False,`
	IEExtensions                       PDUSessionResourceAdmittedItemExtIEs `False,OPTIONAL`
}
