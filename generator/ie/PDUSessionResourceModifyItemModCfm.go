package ie

import "gen/aper"

type PDUSessionResourceModifyItemModCfm struct {
	PDUSessionID                            PDUSessionID                             `False,`
	PDUSessionResourceModifyConfirmTransfer aper.OctetString                         `False,`
	IEExtensions                            PDUSessionResourceModifyItemModCfmExtIEs `False,OPTIONAL`
}
