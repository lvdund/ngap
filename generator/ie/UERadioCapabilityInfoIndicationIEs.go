package ie

type UERadioCapabilityInfoIndicationIEs struct {
	AMFUENGAPID                AMFUENGAPID                `,reject,mandatory`
	RANUENGAPID                RANUENGAPID                `,reject,mandatory`
	UERadioCapability          UERadioCapability          `,ignore,mandatory`
	UERadioCapabilityForPaging UERadioCapabilityForPaging `,ignore,optional`
}
