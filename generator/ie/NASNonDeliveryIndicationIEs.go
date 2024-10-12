package ie

type NASNonDeliveryIndicationIEs struct {
	AMFUENGAPID AMFUENGAPID `,reject,mandatory`
	RANUENGAPID RANUENGAPID `,reject,mandatory`
	NASPDU      NASPDU      `,ignore,mandatory`
	Cause       Cause       `,ignore,mandatory`
}
