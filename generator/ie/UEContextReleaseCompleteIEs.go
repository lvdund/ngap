package ie

type UEContextReleaseCompleteIEs struct {
	AMFUENGAPID                                AMFUENGAPID                                `,ignore,mandatory`
	RANUENGAPID                                RANUENGAPID                                `,ignore,mandatory`
	UserLocationInformation                    UserLocationInformation                    `,ignore,optional`
	InfoOnRecommendedCellsAndRANNodesForPaging InfoOnRecommendedCellsAndRANNodesForPaging `,ignore,optional`
	PDUSessionResourceListCxtRelCpl            PDUSessionResourceListCxtRelCpl            `,reject,optional`
	CriticalityDiagnostics                     CriticalityDiagnostics                     `,ignore,optional`
}
