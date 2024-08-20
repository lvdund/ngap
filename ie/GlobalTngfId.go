package ie

type GlobalTngfId struct {
	PlmnIdentity *PlmnIdentity
	ChoiceTngfId *ChoiceTngfId
}

type ChoiceTngfId struct {
	TngfId *TngfId
}

type TngfId struct {
	TngfId *[]byte
}
