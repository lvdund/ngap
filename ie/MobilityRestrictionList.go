package ie

import "ngap/aper"

type MobilityRestrictionList struct {
ServingPlmn	PlmnIdentity	//`bitstring:"sizeLB:0,sizeUB:150"`
EquivalentPlmns	EquivalentPlmns	//`bitstring:"sizeLB:0,sizeUB:150"`
RatRestrictions	RatRestrictions	//`bitstring:"sizeLB:0,sizeUB:150"`
ForbiddenAreaInformation	ForbiddenAreaInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
ServiceAreaInformation	ServiceAreaInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
LastEUtranPlmnIdentity	PlmnIdentity	//`bitstring:"sizeLB:0,sizeUB:150"`
CoreNetworkTypeRestrictionForServingPlmn	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
CoreNetworkTypeRestrictionForEquivalentPlmns	CoreNetworkTypeRestrictionForEquivalentPlmns	//`bitstring:"sizeLB:0,sizeUB:150"`
NpnMobilityInformation	NpnMobilityInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type EquivalentPlmns struct {
PlmnIdentity	PlmnIdentity	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type RatRestrictions struct {
PlmnIdentity	PlmnIdentity	//`bitstring:"sizeLB:0,sizeUB:150"`
RatRestrictionInformation	aper.BitString	//`bitstring:"sizeLB:0,sizeUB:150"`
ExtendedRatRestrictionInformation	ExtendedRatRestrictionInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type ForbiddenAreaInformation struct {
PlmnIdentity	PlmnIdentity	//`bitstring:"sizeLB:0,sizeUB:150"`
ForbiddenTacs	ForbiddenTacs	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type ForbiddenTacs struct {
Tac	Tac	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type ServiceAreaInformation struct {
PlmnIdentity	PlmnIdentity	//`bitstring:"sizeLB:0,sizeUB:150"`
AllowedTacs	AllowedTacs	//`bitstring:"sizeLB:0,sizeUB:150"`
NotAllowedTacs	NotAllowedTacs	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type AllowedTacs struct {
Tac	Tac	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type NotAllowedTacs struct {
Tac	Tac	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type CoreNetworkTypeRestrictionForEquivalentPlmns struct {
PlmnIdentity	PlmnIdentity	//`bitstring:"sizeLB:0,sizeUB:150"`
CoreNetworkTypeRestriction	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
}
