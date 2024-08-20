package ie

type QmcDeactivation struct {
QoeReferenceList	*QoeReferenceList
}

type QoeReferenceList struct {
QoeReference	[]byte	//`bitstring:"sizeLB:6,sizeUB:6"`
}
