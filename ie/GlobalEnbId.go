package ie

type GlobalEnbId struct {
	PlmnIdentity *PlmnIdentity
	ChoiceEnbId  *ChoiceEnbId
}

type ChoiceEnbId struct {
	MacroEnbId      *MacroEnbId
	HomeEnbId       *HomeEnbId
	ShortMacroEnbId *ShortMacroEnbId
	LongMacroEnbId  *LongMacroEnbId
}

type MacroEnbId struct {
	MacroEnbId []byte //`bitstring:"sizeLB:20,sizeUB:20"`
}

type HomeEnbId struct {
	HomeEnbId []byte //`bitstring:"sizeLB:28,sizeUB:28"`
}

type ShortMacroEnbId struct {
	ShortMacroEnbId []byte //`bitstring:"sizeLB:18,sizeUB:18"`
}

type LongMacroEnbId struct {
	LongMacroEnbId []byte //`bitstring:"sizeLB:21,sizeUB:21"`
}
