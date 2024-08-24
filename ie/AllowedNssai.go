package ie

type AllowedNssai struct {
	AllowedSNssaiList []AllowedSNssaiItem //`bitstring:"sizeLB:0,sizeUB:150"`
}

type AllowedSNssaiItem struct {
	SNssai SNssai //`bitstring:"sizeLB:0,sizeUB:150"`
}
