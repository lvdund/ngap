package ie

type MobilityRestrictionList struct {
	ServingPLMN              PLMNIdentity                  `False,`
	EquivalentPLMNs          EquivalentPLMNs               `False,OPTIONAL`
	RATRestrictions          RATRestrictions               `False,OPTIONAL`
	ForbiddenAreaInformation ForbiddenAreaInformation      `False,OPTIONAL`
	ServiceAreaInformation   ServiceAreaInformation        `False,OPTIONAL`
	IEExtensions             MobilityRestrictionListExtIEs `False,OPTIONAL`
}
