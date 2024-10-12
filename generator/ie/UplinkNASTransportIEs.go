package ie

type UplinkNASTransportIEs struct {
	AMFUENGAPID             AMFUENGAPID             `,reject,mandatory`
	RANUENGAPID             RANUENGAPID             `,reject,mandatory`
	NASPDU                  NASPDU                  `,reject,mandatory`
	UserLocationInformation UserLocationInformation `,ignore,mandatory`
}
