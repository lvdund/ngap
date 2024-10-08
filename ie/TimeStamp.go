package ie

import "ngap/aper"

type TimeStamp struct {
	TimeStamp aper.OctetString `octetstring:"sizeLB:4,sizeUB:4"`
}
