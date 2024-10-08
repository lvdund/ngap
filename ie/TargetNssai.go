package ie

type TargetNssai struct {
	TargetSNssaiList []TargetSNssaiItem `bitstring:"sizeLB:0,sizeUB:150"`
}

type TargetSNssaiItem struct {
	SNssai SNssai `bitstring:"sizeLB:0,sizeUB:150"`
}
