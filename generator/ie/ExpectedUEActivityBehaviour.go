package ie

type ExpectedUEActivityBehaviour struct {
	ExpectedActivityPeriod                 ExpectedActivityPeriod                 `False,OPTIONAL`
	ExpectedIdlePeriod                     ExpectedIdlePeriod                     `False,OPTIONAL`
	SourceOfUEActivityBehaviourInformation SourceOfUEActivityBehaviourInformation `False,OPTIONAL`
	IEExtensions                           ExpectedUEActivityBehaviourExtIEs      `False,OPTIONAL`
}
