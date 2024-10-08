package ie

type ExpectedUEBehaviour struct {
	ExpectedUEActivityBehaviour ExpectedUEActivityBehaviour `True,OPTIONAL`
	ExpectedHOInterval          ExpectedHOInterval          `False,OPTIONAL`
	ExpectedUEMobility          ExpectedUEMobility          `False,OPTIONAL`
	ExpectedUEMovingTrajectory  ExpectedUEMovingTrajectory  `False,OPTIONAL`
	IEExtensions                ExpectedUEBehaviourExtIEs   `False,OPTIONAL`
}
