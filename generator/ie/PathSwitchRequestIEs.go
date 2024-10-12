package ie

type PathSwitchRequestIEs struct {
	RANUENGAPID                              RANUENGAPID                              `,reject,mandatory`
	SourceAMFUENGAPID                        AMFUENGAPID                              `,reject,mandatory`
	UserLocationInformation                  UserLocationInformation                  `,ignore,mandatory`
	UESecurityCapabilities                   UESecurityCapabilities                   `,ignore,mandatory`
	PDUSessionResourceToBeSwitchedDLList     PDUSessionResourceToBeSwitchedDLList     `,reject,mandatory`
	PDUSessionResourceFailedToSetupListPSReq PDUSessionResourceFailedToSetupListPSReq `,ignore,optional`
}
