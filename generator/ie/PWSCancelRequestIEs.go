package ie

type PWSCancelRequestIEs struct {
	MessageIdentifier        MessageIdentifier        `,reject,mandatory`
	SerialNumber             SerialNumber             `,reject,mandatory`
	WarningAreaList          WarningAreaList          `,ignore,optional`
	CancelAllWarningMessages CancelAllWarningMessages `,reject,optional`
}
