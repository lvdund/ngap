package ie

type HandoverRequiredIEs struct {
	AMFUENGAPID                        AMFUENGAPID                        `,reject,mandatory`
	RANUENGAPID                        RANUENGAPID                        `,reject,mandatory`
	HandoverType                       HandoverType                       `,reject,mandatory`
	Cause                              Cause                              `,ignore,mandatory`
	TargetID                           TargetID                           `,reject,mandatory`
	DirectForwardingPathAvailability   DirectForwardingPathAvailability   `,ignore,optional`
	PDUSessionResourceListHORqd        PDUSessionResourceListHORqd        `,reject,mandatory`
	SourceToTargetTransparentContainer SourceToTargetTransparentContainer `,reject,mandatory`
}
