package ie

type NgReset struct {
	MessageType     MessageType     //`bitstring:"sizeLB:0,sizeUB:150"`
	Cause           Cause           //`bitstring:"sizeLB:0,sizeUB:150"`
	ChoiceResetType ChoiceResetType //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceResetType struct {
	NgInterface       NgInterface       //`bitstring:"sizeLB:0,sizeUB:150"`
	PartOfNgInterface PartOfNgInterface //`bitstring:"sizeLB:0,sizeUB:150"`
}

type NgInterface struct {
	ResetAll []byte //`bitstring:"sizeLB:0,sizeUB:150"`
}

type PartOfNgInterface struct {
	UeAssociatedLogicalNgConnectionList UeAssociatedLogicalNgConnectionList //`bitstring:"sizeLB:0,sizeUB:150"`
}
