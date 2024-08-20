package ie

type Ie5GSTmsi struct {
	AmfSetId   *AmfSetId
	AmfPointer *AmfPointer
	Ie5GTmsi   []byte //`bitstring:"sizeLB:4,sizeUB:4"`
}
