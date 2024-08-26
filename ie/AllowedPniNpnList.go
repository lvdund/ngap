package ie

type AllowedPniNpnList struct {
	AllowedPniNpnItem AllowedPniNpnItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type AllowedPniNpnItem struct {
	PlmnIdentity          PlmnIdentity          `bitstring:"sizeLB:0,sizeUB:150"`
	PniNpnRestricted      []byte                `bitstring:"sizeLB:0,sizeUB:150"`
	AllowedCagListPerPlmn AllowedCagListPerPlmn `bitstring:"sizeLB:0,sizeUB:150"`
}

type AllowedCagListPerPlmn struct {
	CagId CagId `bitstring:"sizeLB:0,sizeUB:150"`
}
