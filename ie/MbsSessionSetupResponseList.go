package ie

type MbsSessionSetupResponseList struct {
MbsSessionSetupResponseList	*[]MbsSessionSetupResponseItem
}

type MbsSessionSetupResponseItem struct {
MbsSessionId	*MbsSessionId
MbsAreaSessionId	*MbsAreaSessionId
}
