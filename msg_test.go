package ngap

import (
	"fmt"
	"testing"

	"github.com/lvdund/ngap/ies"
)

func Test_NGSetupRequest(t *testing.T) {
	msg := ies.NGSetupRequest{
		GlobalRANNodeID: ies.GlobalRANNodeID{
			Choice: 1,
			GlobalGNBID: &ies.GlobalGNBID{
				PLMNIdentity: []byte{0x13, 0x30, 0x01},
				GNBID: ies.GNBID{
					Choice: 1,
					GNBID:  []byte{0x02, 0x02},
				},
			},
		},
		SupportedTAList: []ies.SupportedTAItem{
			ies.SupportedTAItem{
				TAC: []byte{0x13, 0x30, 0x01},
				BroadcastPLMNList: []ies.BroadcastPLMNItem{
					ies.BroadcastPLMNItem{
						PLMNIdentity: []byte{0x13, 0x30, 0x01},
						TAISliceSupportList: []ies.SliceSupportItem{
							ies.SliceSupportItem{
								SNSSAI: ies.SNSSAI{SST: []byte{1}},
							},
						},
					},
				},
			},
		},
		DefaultPagingDRX: ies.PagingDRX{Value: 1},
	}
	var b []byte
	var err error
	b, err = NgapEncode(&msg)
	if err != nil {
		fmt.Println("Err encode:", err)
		return
	}
	fmt.Println("encode:", b)

	fmt.Println("================================")
	fmt.Println("================================")

	fmt.Println("Decode:")

	var pdu NgapPdu
	if pdu, err, _ = NgapDecode(b); err != nil {
		fmt.Println("NgapDecode err:", err)
		return
	}
	fmt.Println("================================")
	fmt.Println("Decoded:", pdu)
	decode_msg := pdu.Message.Msg.(*ies.NGSetupRequest)
	fmt.Println(decode_msg.GlobalRANNodeID.Choice)
	fmt.Println("GlobalGNBID PLMNIdentity:", decode_msg.GlobalRANNodeID.GlobalGNBID.PLMNIdentity)
	fmt.Println("GlobalGNBID GNBID:", decode_msg.GlobalRANNodeID.GlobalGNBID.GNBID.Choice, decode_msg.GlobalRANNodeID.GlobalGNBID.GNBID.GNBID)
	fmt.Println("SupportedTAList")
	fmt.Println("TAC:", decode_msg.SupportedTAList[0].TAC)
	fmt.Println("PLMNIdentity:", decode_msg.SupportedTAList[0].BroadcastPLMNList[0].PLMNIdentity)
	fmt.Println("SNSSAI:", decode_msg.SupportedTAList[0].BroadcastPLMNList[0].TAISliceSupportList[0].SNSSAI.SST)
}
