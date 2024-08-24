package ie

import "ngap/aper"

type MdtConfigurationEutra struct {
MdtActivation	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
ChoiceAreaScopeOfMdt	ChoiceAreaScopeOfMdt	//`bitstring:"sizeLB:0,sizeUB:150"`
MdtMode	aper.OctetString	//`octetstring:"sizeLB:0,sizeUB:150"`
SignallingBasedMdtPlmnList	MdtPlmnList	//`bitstring:"sizeLB:0,sizeUB:150"`
}
