package ie

type MbsSessionFailedToSetupList struct {
MbsSessionFailedToSetupList	*[]MbsSessionFailedToSetupItem
}

type MbsSessionFailedToSetupItem struct {
MbsSessionId	*MbsSessionId
MbsAreaSessionId	*MbsAreaSessionId
Cause	*Cause
}
