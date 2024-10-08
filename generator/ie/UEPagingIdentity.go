package ie

type UEPagingIdentity struct {
	FiveGSTMSI       FiveGSTMSI             `True,`
	ChoiceExtensions UEPagingIdentityExtIEs `False,`
}
