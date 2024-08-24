package ie

import "ngap/aper"

type TargetId struct {
ChoiceTargetId	ChoiceTargetId	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceTargetId struct {
NgRan	NgRan	//`bitstring:"sizeLB:0,sizeUB:150"`
EUtran	EUtran	//`bitstring:"sizeLB:0,sizeUB:150"`
TargetRncId	TargetRncId	//`bitstring:"sizeLB:0,sizeUB:150"`
TargetHomeEnbId	TargetHomeEnbId	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type TargetRncId struct {
Lai	Lai	//`bitstring:"sizeLB:0,sizeUB:150"`
RncId	RncId	//`bitstring:"sizeLB:0,sizeUB:150"`
ExtendedRncId	ExtendedRncId	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type TargetHomeEnbId struct {
PlmnIdentity	PlmnIdentity	//`bitstring:"sizeLB:0,sizeUB:150"`
HomeEnbId	aper.BitString	//`bitstring:"sizeLB:28,sizeUB:28"`
SelectedEpsTai	EpsTai	//`bitstring:"sizeLB:0,sizeUB:150"`
}
