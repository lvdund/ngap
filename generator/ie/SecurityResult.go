package ie

type SecurityResult struct {
	IntegrityProtectionResult       IntegrityProtectionResult       `False,`
	ConfidentialityProtectionResult ConfidentialityProtectionResult `False,`
	IEExtensions                    SecurityResultExtIEs            `False,OPTIONAL`
}
