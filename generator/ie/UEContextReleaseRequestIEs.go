package ie

type UEContextReleaseRequestIEs struct {
	AMFUENGAPID                     AMFUENGAPID                     `,reject,mandatory`
	RANUENGAPID                     RANUENGAPID                     `,reject,mandatory`
	PDUSessionResourceListCxtRelReq PDUSessionResourceListCxtRelReq `,reject,optional`
	Cause                           Cause                           `,ignore,mandatory`
}
