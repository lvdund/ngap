package ie

type TargetId struct {
	ChoiceTargetId *ChoiceTargetId
}

type ChoiceTargetId struct {
	NgRan           *NgRan
	EUtran          *EUtran
	TargetRncId     *TargetRncId
	TargetHomeEnbId *TargetHomeEnbId
}

type TargetRncId struct {
	Lai           *Lai
	RncId         *RncId
	ExtendedRncId *ExtendedRncId
}

type TargetHomeEnbId struct {
	PlmnIdentity   *PlmnIdentity
	HomeEnbId      []byte //`bitstring:"sizeLB:28,sizeUB:28"`
	SelectedEpsTai *EpsTai
}
