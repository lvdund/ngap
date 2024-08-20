package ie

type GlobalNgEnbId struct {
	PlmnIdentity  *PlmnIdentity
	ChoiceNgEnbId *ChoiceNgEnbId
}

type ChoiceNgEnbId struct {
	MacroNgEnbId      *MacroNgEnbId
	ShortMacroNgEnbId *ShortMacroNgEnbId
	LongMacroNgEnbId  *LongMacroNgEnbId
}

type MacroNgEnbId struct {
	MacroNgEnbId []byte //`bitstring:"sizeLB:20,sizeUB:20"`
}

type ShortMacroNgEnbId struct {
	ShortMacroNgEnbId []byte //`bitstring:"sizeLB:18,sizeUB:18"`
}

type LongMacroNgEnbId struct {
	LongMacroNgEnbId []byte //`bitstring:"sizeLB:21,sizeUB:21"`
}
