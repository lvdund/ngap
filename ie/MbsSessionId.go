package ie

type MbsSessionId struct {
	Tmgi []byte //`bitstring:"sizeLB:6,sizeUB:6"`
	Nid  *Nid
}
