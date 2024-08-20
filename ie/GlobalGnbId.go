package ie

type GlobalGnbId struct {
	PlmnIdentity *PlmnIdentity
	ChoiceGnbId  *ChoiceGnbId
}

type ChoiceGnbId struct {
	GnbId *GnbId
}

type GnbId struct {
	GnbId []byte //`bitstring:"sizeLB:22,sizeUB:32"`
}
