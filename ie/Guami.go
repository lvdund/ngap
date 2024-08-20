package ie

type Guami struct {
	PlmnIdentity *PlmnIdentity
	AmfRegionId  []byte //`bitstring:"sizeLB:8,sizeUB:8"`
	AmfSetId     *AmfSetId
	AmfPointer   *AmfPointer
}
