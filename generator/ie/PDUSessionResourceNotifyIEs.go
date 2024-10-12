package ie

type PDUSessionResourceNotifyIEs struct {
	AMFUENGAPID                       AMFUENGAPID                       `,reject,mandatory`
	RANUENGAPID                       RANUENGAPID                       `,reject,mandatory`
	PDUSessionResourceNotifyList      PDUSessionResourceNotifyList      `,reject,optional`
	PDUSessionResourceReleasedListNot PDUSessionResourceReleasedListNot `,ignore,optional`
	UserLocationInformation           UserLocationInformation           `,ignore,optional`
}
