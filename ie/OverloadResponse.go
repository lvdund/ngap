package ie

type OverloadResponse struct {
ChoiceOverloadResponse	ChoiceOverloadResponse	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceOverloadResponse struct {
OverloadAction	OverloadAction	//`bitstring:"sizeLB:0,sizeUB:150"`
}
