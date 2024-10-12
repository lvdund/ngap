package ie

type SecondaryRATDataUsageReportIEs struct {
	AMFUENGAPID                             AMFUENGAPID                             `,ignore,mandatory`
	RANUENGAPID                             RANUENGAPID                             `,ignore,mandatory`
	PDUSessionResourceSecondaryRATUsageList PDUSessionResourceSecondaryRATUsageList `,ignore,mandatory`
	HandoverFlag                            HandoverFlag                            `,ignore,optional`
	UserLocationInformation                 UserLocationInformation                 `,ignore,optional`
}
