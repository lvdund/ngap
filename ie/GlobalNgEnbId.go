package ie

import "ngap/aper"

type GlobalNgEnbId struct {
	PlmnIdentity  PlmnIdentity  //`bitstring:"sizeLB:0,sizeUB:150"`
	ChoiceNgEnbId ChoiceNgEnbId //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceNgEnbId struct {
	MacroNgEnbId      MacroNgEnbId      //`bitstring:"sizeLB:0,sizeUB:150"`
	ShortMacroNgEnbId ShortMacroNgEnbId //`bitstring:"sizeLB:0,sizeUB:150"`
	LongMacroNgEnbId  LongMacroNgEnbId  //`bitstring:"sizeLB:0,sizeUB:150"`
}

type MacroNgEnbId struct {
	MacroNgEnbId aper.BitString //`bitstring:"sizeLB:20,sizeUB:20"`
}

type ShortMacroNgEnbId struct {
	ShortMacroNgEnbId aper.BitString //`bitstring:"sizeLB:18,sizeUB:18"`
}

type LongMacroNgEnbId struct {
	LongMacroNgEnbId aper.BitString //`bitstring:"sizeLB:21,sizeUB:21"`
}
