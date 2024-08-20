package ie

type MbsSessionToReleaseList struct {
MbsSessionToReleaseList	*[]MbsSessionToReleaseItem
}

type MbsSessionToReleaseItem struct {
MbsSessionId	*MbsSessionId
Cause	*Cause
}
