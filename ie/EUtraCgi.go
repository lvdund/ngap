package ie

type EUtraCgi struct {
	PlmnIdentity      *PlmnIdentity
	EUtraCellIdentity []byte //`bitstring:"sizeLB:28,sizeUB:28"`
}
