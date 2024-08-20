package ie

type Lai struct {
	PlmnIdentity *PlmnIdentity
	Lac          []byte //`bitstring:"sizeLB:2,sizeUB:2"`
}
