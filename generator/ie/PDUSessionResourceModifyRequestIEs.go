package ie

type PDUSessionResourceModifyRequestIEs struct {
	AMFUENGAPID                        AMFUENGAPID                        `,reject,mandatory`
	RANUENGAPID                        RANUENGAPID                        `,reject,mandatory`
	RANPagingPriority                  RANPagingPriority                  `,ignore,optional`
	PDUSessionResourceModifyListModReq PDUSessionResourceModifyListModReq `,reject,mandatory`
}
