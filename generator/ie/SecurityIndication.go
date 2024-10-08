package ie

type SecurityIndication struct {
	IntegrityProtectionIndication       IntegrityProtectionIndication       `False,`
	ConfidentialityProtectionIndication ConfidentialityProtectionIndication `False,`
	MaximumIntegrityProtectedDataRateUL MaximumIntegrityProtectedDataRate   `False,OPTIONAL`
	IEExtensions                        SecurityIndicationExtIEs            `False,OPTIONAL`
}
