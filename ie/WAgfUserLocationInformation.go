package ie

type WAgfUserLocationInformation struct {
ChoiceWAgfUserLocationInformation	*ChoiceWAgfUserLocationInformation
}

type ChoiceWAgfUserLocationInformation struct {
GlobalLineId	*GlobalLineId
HfcNodeId	*HfcNodeId
GlobalCableId	*GlobalCableId
HfcNodeIdNew	*HfcNodeIdNew
GlobalCableIdNew	*GlobalCableIdNew
}

type GlobalLineId struct {
GlobalLineId	*[]byte
LineType	*[]byte
Tai	*Tai
}

type HfcNodeId struct {
HfcNodeId	*[]byte
}

type GlobalCableId struct {
GlobalCableId	*[]byte
}

type HfcNodeIdNew struct {
HfcNodeId	*[]byte
Tai	*Tai
}

type GlobalCableIdNew struct {
GlobalCableId	*[]byte
Tai	*Tai
}
