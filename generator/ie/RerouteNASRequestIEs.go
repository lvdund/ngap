package ie

import "gen/aper"

type RerouteNASRequestIEs struct {
	RANUENGAPID                         RANUENGAPID                         `,reject,mandatory`
	AMFUENGAPID                         AMFUENGAPID                         `,ignore,optional`
	NGAPMessage                         aper.OctetString                    `,reject,mandatory`
	AMFSetID                            AMFSetID                            `,reject,mandatory`
	AllowedNSSAI                        AllowedNSSAI                        `,reject,optional`
	SourceToTargetAMFInformationReroute SourceToTargetAMFInformationReroute `,ignore,optional`
}
