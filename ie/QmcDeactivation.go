package ie

import "ngap/aper"

type QmcDeactivation struct {
QoeReferenceList	QoeReferenceList	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type QoeReferenceList struct {
QoeReference	aper.OctetString	//`octetstring:"sizeLB:6,sizeUB:6"`
}
