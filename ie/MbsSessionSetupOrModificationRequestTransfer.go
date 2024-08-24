package ie

import "ngap/aper"

type MbsSessionSetupOrModificationRequestTransfer struct {
MbsSessionTnlInformation5Gc	MbsSessionTnlInformation5Gc	//`bitstring:"sizeLB:0,sizeUB:150"`
MbsQosFlowsToBeSetupOrModifiedList	MbsQosFlowsToBeSetupList	//`bitstring:"sizeLB:0,sizeUB:150"`
MbsSessionFsaIdList	MbsSessionFsaIdList	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type MbsSessionFsaIdList struct {
MbsFrequencySelectionAreaIdentity	aper.OctetString	//`octetstring:"sizeLB:3,sizeUB:3"`
}
