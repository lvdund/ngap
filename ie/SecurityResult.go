package ie

type SecurityResult struct {
	IntegrityProtectionResult       []byte `bitstring:"sizeLB:0,sizeUB:150"`
	ConfidentialityProtectionResult []byte `bitstring:"sizeLB:0,sizeUB:150"`
}
