package ie

type LocationReportIEs struct {
	AMFUENGAPID                    AMFUENGAPID                    `,reject,mandatory`
	RANUENGAPID                    RANUENGAPID                    `,reject,mandatory`
	UserLocationInformation        UserLocationInformation        `,ignore,mandatory`
	UEPresenceInAreaOfInterestList UEPresenceInAreaOfInterestList `,ignore,optional`
	LocationReportingRequestType   LocationReportingRequestType   `,ignore,mandatory`
}
