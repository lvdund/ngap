package ie

type UeIdentityIndexValue struct {
	ChoiceUeIdentityIndexValue *ChoiceUeIdentityIndexValue
}

type ChoiceUeIdentityIndexValue struct {
	IndexLength10 *IndexLength10
}

type IndexLength10 struct {
	IndexLength10 []byte //`bitstring:"sizeLB:10,sizeUB:10"`
}
