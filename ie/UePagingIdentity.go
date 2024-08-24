package ie

type UePagingIdentity struct {
ChoiceUePagingIdentity	ChoiceUePagingIdentity	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceUePagingIdentity struct {
Ie5GSTmsi	Ie5GSTmsi	//`bitstring:"sizeLB:0,sizeUB:150"`
}
