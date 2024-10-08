package ie

type CoreNetworkAssistanceInformationForInactive struct {
	UEIdentityIndexValue            UEIdentityIndexValue                              `False,`
	UESpecificDRX                   PagingDRX                                         `False,OPTIONAL`
	PeriodicRegistrationUpdateTimer PeriodicRegistrationUpdateTimer                   `False,`
	MICOModeIndication              MICOModeIndication                                `False,OPTIONAL`
	TAIListForInactive              TAIListForInactive                                `False,`
	ExpectedUEBehaviour             ExpectedUEBehaviour                               `True,OPTIONAL`
	IEExtensions                    CoreNetworkAssistanceInformationForInactiveExtIEs `False,OPTIONAL`
}
