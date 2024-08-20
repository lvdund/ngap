package ie

type UeContextReleaseCommand struct {
	MessageType     *MessageType
	ChoiceUeNgapIds *ChoiceUeNgapIds
	Cause           *Cause
}

type ChoiceUeNgapIds struct {
	UeNgapIdPair *UeNgapIdPair
	AmfUeNgapId  *AmfUeNgapId
}

type UeNgapIdPair struct {
	AmfUeNgapId *AmfUeNgapId
	RanUeNgapId *RanUeNgapId
}
