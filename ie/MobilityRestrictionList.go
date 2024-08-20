package ie

type MobilityRestrictionList struct {
ServingPlmn	*PlmnIdentity
EquivalentPlmns	*EquivalentPlmns
RatRestrictions	*RatRestrictions
ForbiddenAreaInformation	*ForbiddenAreaInformation
ServiceAreaInformation	*ServiceAreaInformation
LastEUtranPlmnIdentity	*PlmnIdentity
CoreNetworkTypeRestrictionForServingPlmn	*[]byte
CoreNetworkTypeRestrictionForEquivalentPlmns	*CoreNetworkTypeRestrictionForEquivalentPlmns
NpnMobilityInformation	*NpnMobilityInformation
}

type EquivalentPlmns struct {
PlmnIdentity	*PlmnIdentity
}

type RatRestrictions struct {
PlmnIdentity	*PlmnIdentity
RatRestrictionInformation	*[]byte
ExtendedRatRestrictionInformation	*ExtendedRatRestrictionInformation
}

type ForbiddenAreaInformation struct {
PlmnIdentity	*PlmnIdentity
ForbiddenTacs	*ForbiddenTacs
}

type ForbiddenTacs struct {
Tac	*Tac
}

type ServiceAreaInformation struct {
PlmnIdentity	*PlmnIdentity
AllowedTacs	*AllowedTacs
NotAllowedTacs	*NotAllowedTacs
}

type AllowedTacs struct {
Tac	*Tac
}

type NotAllowedTacs struct {
Tac	*Tac
}

type CoreNetworkTypeRestrictionForEquivalentPlmns struct {
PlmnIdentity	*PlmnIdentity
CoreNetworkTypeRestriction	*[]byte
}
