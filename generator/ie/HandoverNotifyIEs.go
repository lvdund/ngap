package ie

type HandoverNotifyIEs struct {
	AMFUENGAPID             AMFUENGAPID             `,reject,mandatory`
	RANUENGAPID             RANUENGAPID             `,reject,mandatory`
	UserLocationInformation UserLocationInformation `,ignore,mandatory`
}
