package ie

type MbsSessionTnlInformation5Gc struct {
	ChoiceSessionType ChoiceSessionType `bitstring:"sizeLB:0,sizeUB:150"`
}

type MbsSessionTnlInformation5GcItem struct {
	MbsAreaSessionId                 MbsAreaSessionId                 `bitstring:"sizeLB:0,sizeUB:150"`
	SharedNgUMulticastTnlInformation SharedNgUMulticastTnlInformation `bitstring:"sizeLB:0,sizeUB:150"`
}
