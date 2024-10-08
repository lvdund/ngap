package ie

type PDUSessionResourceItemCxtRelReq struct {
	PDUSessionID PDUSessionID                          `False,`
	IEExtensions PDUSessionResourceItemCxtRelReqExtIEs `False,OPTIONAL`
}
