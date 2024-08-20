package ie

type MbsSessionSetupOrModificationRequestTransfer struct {
MbsSessionTnlInformation5Gc	*MbsSessionTnlInformation5Gc
MbsQosFlowsToBeSetupOrModifiedList	*MbsQosFlowsToBeSetupList
MbsSessionFsaIdList	*MbsSessionFsaIdList
}

type MbsSessionFsaIdList struct {
MbsFrequencySelectionAreaIdentity	[]byte	//`bitstring:"sizeLB:3,sizeUB:3"`
}
