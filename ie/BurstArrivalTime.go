package ie

import "ngap/aper"

type BurstArrivalTime struct {
	BurstArrivalTime aper.OctetString `octetstring:"sizeLB:0,sizeUB:150"`
}
