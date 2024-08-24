package ie

import "ngap/aper"

type GlobalEnbId struct {
	PlmnIdentity PlmnIdentity //`bitstring:"sizeLB:0,sizeUB:150"`
	ChoiceEnbId  ChoiceEnbId  //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceEnbId struct {
	MacroEnbId      MacroEnbId      //`bitstring:"sizeLB:0,sizeUB:150"`
	HomeEnbId       HomeEnbId       //`bitstring:"sizeLB:0,sizeUB:150"`
	ShortMacroEnbId ShortMacroEnbId //`bitstring:"sizeLB:0,sizeUB:150"`
	LongMacroEnbId  LongMacroEnbId  //`bitstring:"sizeLB:0,sizeUB:150"`
}

type MacroEnbId struct {
	MacroEnbId aper.BitString //`bitstring:"sizeLB:20,sizeUB:20"`
}

type HomeEnbId struct {
	HomeEnbId aper.BitString //`bitstring:"sizeLB:28,sizeUB:28"`
}

type ShortMacroEnbId struct {
	ShortMacroEnbId aper.BitString //`bitstring:"sizeLB:18,sizeUB:18"`
}

type LongMacroEnbId struct {
	LongMacroEnbId aper.BitString //`bitstring:"sizeLB:21,sizeUB:21"`
}
