package ie

type PDUSessionResourceSetupRequestIEs struct {
	AMFUENGAPID                      AMFUENGAPID                      `,reject,mandatory`
	RANUENGAPID                      RANUENGAPID                      `,reject,mandatory`
	RANPagingPriority                RANPagingPriority                `,ignore,optional`
	NASPDU                           NASPDU                           `,reject,optional`
	PDUSessionResourceSetupListSUReq PDUSessionResourceSetupListSUReq `,reject,mandatory`
	UEAggregateMaximumBitRate        UEAggregateMaximumBitRate        `,ignore,optional`
}
