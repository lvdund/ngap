package ie

type NGSetupResponseIEs struct {
	AMFName                AMFName                `,reject,mandatory`
	ServedGUAMIList        ServedGUAMIList        `,reject,mandatory`
	RelativeAMFCapacity    RelativeAMFCapacity    `,ignore,mandatory`
	PLMNSupportList        PLMNSupportList        `,reject,mandatory`
	CriticalityDiagnostics CriticalityDiagnostics `,ignore,optional`
	UERetentionInformation UERetentionInformation `,ignore,optional`
}
