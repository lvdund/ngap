package ie

type SonConfigurationTransfer struct {
TargetRanNodeIdSon	*TargetRanNodeIdSon
SourceRanNodeId	*SourceRanNodeId
SonInformation	*SonInformation
XnTnlConfigurationInfo	*XnTnlConfigurationInfo
}

type TargetRanNodeIdSon struct {
GlobalRanNodeId	*GlobalRanNodeId
SelectedTai	*Tai
NrCgi	*NrCgi
}
