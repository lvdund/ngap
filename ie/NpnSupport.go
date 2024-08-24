package ie

type NpnSupport struct {
ChoiceNpnSupport	ChoiceNpnSupport	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceNpnSupport struct {
Snpn	Snpn	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type Snpn struct {
Nid	Nid	//`bitstring:"sizeLB:0,sizeUB:150"`
}
