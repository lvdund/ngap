package ie

import "ngap/aper"

type SourceToTargetAmfInformationReroute struct {
	ConfiguredNssai     aper.OctetString //`octetstring:"sizeLB:128,sizeUB:128"`
	RejectedNssaiInPlmn aper.OctetString //`octetstring:"sizeLB:32,sizeUB:32"`
	RejectedNssaiInTa   aper.OctetString //`octetstring:"sizeLB:32,sizeUB:32"`
}
