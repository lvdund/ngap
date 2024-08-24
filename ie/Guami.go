package ie

import "ngap/aper"

type Guami struct {
	PlmnIdentity PlmnIdentity   //`bitstring:"sizeLB:0,sizeUB:150"`
	AmfRegionId  aper.BitString //`bitstring:"sizeLB:8,sizeUB:8"`
	AmfSetId     AmfSetId       //`bitstring:"sizeLB:0,sizeUB:150"`
	AmfPointer   AmfPointer     //`bitstring:"sizeLB:0,sizeUB:150"`
}
