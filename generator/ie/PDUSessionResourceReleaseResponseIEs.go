package ie

type PDUSessionResourceReleaseResponseIEs struct {
	AMFUENGAPID                          AMFUENGAPID                          `,ignore,mandatory`
	RANUENGAPID                          RANUENGAPID                          `,ignore,mandatory`
	PDUSessionResourceReleasedListRelRes PDUSessionResourceReleasedListRelRes `,ignore,mandatory`
	UserLocationInformation              UserLocationInformation              `,ignore,optional`
	CriticalityDiagnostics               CriticalityDiagnostics               `,ignore,optional`
}
