package ie

type NrCgi struct {
	PlmnIdentity   *PlmnIdentity
	NrCellIdentity []byte //`bitstring:"sizeLB:36,sizeUB:36"`
}
