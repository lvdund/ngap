package ie

type PDUSessionResourceModifyResponseIEs struct {
	AMFUENGAPID                                AMFUENGAPID                                `,ignore,mandatory`
	RANUENGAPID                                RANUENGAPID                                `,ignore,mandatory`
	PDUSessionResourceModifyListModRes         PDUSessionResourceModifyListModRes         `,ignore,optional`
	PDUSessionResourceFailedToModifyListModRes PDUSessionResourceFailedToModifyListModRes `,ignore,optional`
	UserLocationInformation                    UserLocationInformation                    `,ignore,optional`
	CriticalityDiagnostics                     CriticalityDiagnostics                     `,ignore,optional`
}
