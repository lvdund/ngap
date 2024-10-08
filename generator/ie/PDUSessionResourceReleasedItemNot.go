package ie

import "gen/aper"

type PDUSessionResourceReleasedItemNot struct {
	PDUSessionID                             PDUSessionID                            `False,`
	PDUSessionResourceNotifyReleasedTransfer aper.OctetString                        `False,`
	IEExtensions                             PDUSessionResourceReleasedItemNotExtIEs `False,OPTIONAL`
}
