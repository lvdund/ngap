package ie

import "ngap/aper"

type ExtendedRatRestrictionInformation struct {
	PrimaryRatRestriction   aper.BitString `bitstring:"sizeLB:0,sizeUB:150"`
	SecondaryRatRestriction aper.BitString `bitstring:"sizeLB:0,sizeUB:150"`
}
