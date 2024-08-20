package ie

type SecondaryRatDataUsageReport struct {
	MessageType                             *MessageType
	AmfUeNgapId                             *AmfUeNgapId
	RanUeNgapId                             *RanUeNgapId
	PduSessionResourceSecondaryRatUsageList *[]PduSessionResourceSecondaryRatUsageItem
	HandoverFlag                            *[]byte
	UserLocationInformation                 *UserLocationInformation
}

type PduSessionResourceSecondaryRatUsageItem struct {
	PduSessionId                        *PduSessionId
	SecondaryRatDataUsageReportTransfer *[]byte
}
