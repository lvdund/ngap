package ie

type PDUSessionResourceReleaseCommandIEs struct {
	AMFUENGAPID                           AMFUENGAPID                           `,reject,mandatory`
	RANUENGAPID                           RANUENGAPID                           `,reject,mandatory`
	RANPagingPriority                     RANPagingPriority                     `,ignore,optional`
	NASPDU                                NASPDU                                `,ignore,optional`
	PDUSessionResourceToReleaseListRelCmd PDUSessionResourceToReleaseListRelCmd `,reject,mandatory`
}
