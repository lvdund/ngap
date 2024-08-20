package ie

type ConnectionEstablishmentIndication struct {
MessageType	*MessageType
AmfUeNgapId	*AmfUeNgapId
RanUeNgapId	*RanUeNgapId
UeRadioCapability	*UeRadioCapability
EndIndication		*EndIndication	
SNssai	*SNssai
AllowedNssai	*AllowedNssai
UeDifferentiationInformation	*UeDifferentiationInformation
DlCpSecurityInformation	*DlCpSecurityInformation
NbIotUePriority	*NbIotUePriority
EnhancedCoverageRestriction	*EnhancedCoverageRestriction
CeModeBRestricted	*CeModeBRestricted
UeRadioCapabilityId	*UeRadioCapabilityId
MaskedImeisv	*MaskedImeisv
OldAmf	*AmfName
}
