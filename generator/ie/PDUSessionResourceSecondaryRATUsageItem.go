package ie

import "gen/aper"

type PDUSessionResourceSecondaryRATUsageItem struct {
	PDUSessionID                        PDUSessionID                                  `False,`
	SecondaryRATDataUsageReportTransfer aper.OctetString                              `False,`
	IEExtensions                        PDUSessionResourceSecondaryRATUsageItemExtIEs `False,OPTIONAL`
}
