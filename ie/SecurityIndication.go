package ie

type SecurityIndication struct {
IntegrityProtectionIndication	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
ConfidentialityProtectionIndication	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
MaximumIntegrityProtectedDataRateUplink	MaximumIntegrityProtectedDataRate	//`bitstring:"sizeLB:0,sizeUB:150"`
MaximumIntegrityProtectedDataRateDownlink	MaximumIntegrityProtectedDataRate	//`bitstring:"sizeLB:0,sizeUB:150"`
}
