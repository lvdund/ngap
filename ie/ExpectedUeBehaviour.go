package ie

type ExpectedUeBehaviour struct {
ExpectedUeActivityBehaviour	*ExpectedUeActivityBehaviour
ExpectedHoInterval	*[]byte
ExpectedUeMobility	*[]byte
ExpectedUeMovingTrajectory	*ExpectedUeMovingTrajectory
}

type ExpectedUeMovingTrajectory struct {
ExpectedUeMovingTrajectoryItem	*ExpectedUeMovingTrajectoryItem
}

type ExpectedUeMovingTrajectoryItem struct {
NgRanCgi	*NgRanCgi
TimeStayedInCell	uint16	//`bitstring:"sizeLB:0,sizeUB:4095"`
}
