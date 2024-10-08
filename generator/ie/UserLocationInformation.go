package ie

type UserLocationInformation struct {
	UserLocationInformationEUTRA UserLocationInformationEUTRA  `True,`
	UserLocationInformationNR    UserLocationInformationNR     `True,`
	UserLocationInformationN3IWF UserLocationInformationN3IWF  `True,`
	ChoiceExtensions             UserLocationInformationExtIEs `False,`
}
