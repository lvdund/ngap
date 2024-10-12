package ie

type HandoverCommandIEs struct {
	AMFUENGAPID                          AMFUENGAPID                          `,reject,mandatory`
	RANUENGAPID                          RANUENGAPID                          `,reject,mandatory`
	HandoverType                         HandoverType                         `,reject,mandatory`
	NASSecurityParametersFromNGRAN       NASSecurityParametersFromNGRAN       `,reject,conditional`
	PDUSessionResourceHandoverList       PDUSessionResourceHandoverList       `,ignore,optional`
	PDUSessionResourceToReleaseListHOCmd PDUSessionResourceToReleaseListHOCmd `,ignore,optional`
	TargetToSourceTransparentContainer   TargetToSourceTransparentContainer   `,reject,mandatory`
	CriticalityDiagnostics               CriticalityDiagnostics               `,ignore,optional`
}
