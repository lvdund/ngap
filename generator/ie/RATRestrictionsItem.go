package ie

type RATRestrictionsItem struct {
	PLMNIdentity              PLMNIdentity              `False,`
	RATRestrictionInformation RATRestrictionInformation `False,`
	IEExtensions              RATRestrictionsItemExtIEs `False,OPTIONAL`
}
