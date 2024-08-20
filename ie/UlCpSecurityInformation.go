package ie

type UlCpSecurityInformation struct {
	UlNasMac   []byte //`bitstring:"sizeLB:16,sizeUB:16"`
	UlNasCount []byte //`bitstring:"sizeLB:5,sizeUB:5"`
}
