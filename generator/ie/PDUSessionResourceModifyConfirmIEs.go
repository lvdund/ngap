package ie

type PDUSessionResourceModifyConfirmIEs struct {
	AMFUENGAPID                                AMFUENGAPID                                `,ignore,mandatory`
	RANUENGAPID                                RANUENGAPID                                `,ignore,mandatory`
	PDUSessionResourceModifyListModCfm         PDUSessionResourceModifyListModCfm         `,ignore,optional`
	PDUSessionResourceFailedToModifyListModCfm PDUSessionResourceFailedToModifyListModCfm `,ignore,optional`
	CriticalityDiagnostics                     CriticalityDiagnostics                     `,ignore,optional`
}
