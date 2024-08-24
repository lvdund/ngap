package ie

import "ngap/aper"

type UeRadioCapabilityForPaging struct {
	UeRadioCapabilityForPagingOfNr    aper.OctetString //`octetstring:"sizeLB:0,sizeUB:150"`
	UeRadioCapabilityForPagingOfEUtra aper.OctetString //`octetstring:"sizeLB:0,sizeUB:150"`
	UeRadioCapabilityForPagingOfNbIot aper.OctetString //`octetstring:"sizeLB:0,sizeUB:150"`
}
