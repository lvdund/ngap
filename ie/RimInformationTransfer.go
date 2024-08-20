package ie

type RimInformationTransfer struct {
TargetRanNodeIdRim	*TargetRanNodeIdRim
SourceRanNodeId	*SourceRanNodeId
RimInformation	*RimInformation
}

type TargetRanNodeIdRim struct {
GlobalRanNodeId	*GlobalRanNodeId
SelectedTai	*Tai
}

type SourceRanNodeId struct {
GlobalRanNodeId	*GlobalRanNodeId
SelectedTai	*Tai
}
