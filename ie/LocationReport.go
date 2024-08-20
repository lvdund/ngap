package ie

type LocationReport struct {
	MessageType                    *MessageType
	AmfUeNgapId                    *AmfUeNgapId
	RanUeNgapId                    *RanUeNgapId
	UserLocationInformation        *UserLocationInformation
	UePresenceInAreaOfInterestList *UePresenceInAreaOfInterestList
	LocationReportingRequestType   *LocationReportingRequestType
}
