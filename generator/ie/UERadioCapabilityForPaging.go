package ie

type UERadioCapabilityForPaging struct {
	UERadioCapabilityForPagingOfNR    UERadioCapabilityForPagingOfNR    `False,OPTIONAL`
	UERadioCapabilityForPagingOfEUTRA UERadioCapabilityForPagingOfEUTRA `False,OPTIONAL`
	IEExtensions                      UERadioCapabilityForPagingExtIEs  `False,OPTIONAL`
}
