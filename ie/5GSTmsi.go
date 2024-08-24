package ie

import "ngap/aper"

type Ie5GSTmsi struct {
	AmfSetId   AmfSetId         //`bitstring:"sizeLB:0,sizeUB:150"`
	AmfPointer AmfPointer       //`bitstring:"sizeLB:0,sizeUB:150"`
	Ie5GTmsi   aper.OctetString //`octetstring:"sizeLB:4,sizeUB:4"`
}
