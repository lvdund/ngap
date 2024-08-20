package ie

type MbsSessionTnlInformationNgRan struct {
ChoiceSessionType	*ChoiceSessionType
}

type ChoiceSessionType struct {
LocationIndependent	*LocationIndependent
LocationDependent	*LocationDependent
}

type LocationIndependent struct {
SharedNgUUnicastTnlInformation	*UpTransportLayerInformation
}

type LocationDependent struct {
MbsSessionTnlInformationNgRanList	*[]MbsSessionTnlInformationNgRanItem
}

type MbsSessionTnlInformationNgRanItem struct {
MbsAreaSessionId	*MbsAreaSessionId
SharedNgUUnicastTnlInformation	*UpTransportLayerInformation
}
