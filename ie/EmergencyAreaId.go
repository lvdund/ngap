package ie

import "ngap/aper"

type EmergencyAreaId struct {
	EmergencyAreaId aper.OctetString `octetstring:"sizeLB:3,sizeUB:3"`
}
