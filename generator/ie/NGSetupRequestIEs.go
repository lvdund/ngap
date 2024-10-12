package ie

type NGSetupRequestIEs struct {
	GlobalRANNodeID        GlobalRANNodeID        `,reject,mandatory`
	RANNodeName            RANNodeName            `,ignore,optional`
	SupportedTAList        SupportedTAList        `,reject,mandatory`
	DefaultPagingDRX       PagingDRX              `,ignore,mandatory`
	UERetentionInformation UERetentionInformation `,ignore,optional`
}
