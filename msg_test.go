package ngap

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/lvdund/ngap/ies"
	"github.com/lvdund/ngap/utils"
)

func Test_NGSetupRequest(t *testing.T) {
	msg := ies.NGSetupRequest{
		GlobalRANNodeID: ies.GlobalRANNodeID{
			Choice: 1,
			GlobalGNBID: &ies.GlobalGNBID{
				PLMNIdentity: []byte{0x13, 0x30, 0x01},
				GNBID: ies.GNBID{
					Choice: 1,
					GNBID:  []byte{0x12, 0x34, 0x56, 0x78},
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
	var cridia *ies.CriticalityDiagnostics
	if pdu, err, cridia = NgapDecode(b); err != nil {
		fmt.Println("NgapDecode err:", err)
		return
	}
	fmt.Println("================================ CriticalityDiagnostics:", cridia)
	fmt.Println("================================ MSG NGSetupRequest:")
	decode_msg := pdu.Message.Msg.(*ies.NGSetupRequest)
	fmt.Println(decode_msg.GlobalRANNodeID.GlobalGNBID.GNBID)
	fmt.Println("GlobalGNBID PLMNIdentity:", decode_msg.GlobalRANNodeID.GlobalGNBID.PLMNIdentity)
	fmt.Println("GlobalGNBID GNBID:", decode_msg.GlobalRANNodeID.GlobalGNBID.GNBID.Choice, decode_msg.GlobalRANNodeID.GlobalGNBID.GNBID.GNBID)
	fmt.Println("SupportedTAList")
	fmt.Println("TAC:", decode_msg.SupportedTAList[0].TAC)
	fmt.Println("PLMNIdentity:", decode_msg.SupportedTAList[0].BroadcastPLMNList[0].PLMNIdentity)
	fmt.Println("SNSSAI:", decode_msg.SupportedTAList[0].BroadcastPLMNList[0].TAISliceSupportList[0].SNSSAI.SST)
}

var m1 []byte = []byte{0, 14, 0, 128, 160, 0, 0, 9, 0, 10, 0, 2, 0, 1, 0, 85, 0, 2, 0, 1, 0, 28, 0, 7, 0, 2, 248, 57, 202, 254, 0, 0, 0, 0, 5, 2, 1, 1, 2, 3, 0, 119, 0, 9, 8, 0, 4, 0, 0, 0, 0, 0, 0, 0, 94, 0, 32, 225, 111, 144, 208, 212, 77, 76, 241, 133, 37, 88, 8, 123, 238, 46, 143, 36, 178, 16, 214, 75, 211, 29, 163, 250, 176, 158, 131, 118, 105, 18, 9, 0, 36, 64, 4, 0, 2, 248, 57, 0, 34, 64, 8, 17, 16, 0, 0, 0, 255, 255, 0, 0, 38, 64, 52, 51, 126, 2, 216, 233, 89, 233, 1, 126, 0, 66, 1, 1, 119, 0, 11, 242, 2, 248, 57, 202, 254, 0, 0, 0, 0, 1, 84, 7, 0, 2, 248, 57, 0, 0, 1, 21, 5, 4, 1, 1, 2, 3, 33, 1, 0, 94, 1, 6, 22, 1, 44}

func TestM1(t *testing.T) {
	var pdu NgapPdu
	var err error
	var diagnostics *ies.CriticalityDiagnostics
	pdu, err, diagnostics = NgapDecode(m1)
	if err != nil {
		fmt.Println(err)
		fmt.Println(diagnostics)
		return
	}
	msg1 := pdu.Message.Msg.(*ies.InitialContextSetupRequest)
	fmt.Println(msg1.AMFUENGAPID)
	fmt.Println(msg1.RANUENGAPID)
	fmt.Println(msg1.GUAMI)
	fmt.Println(msg1.AllowedNSSAI)
	fmt.Println(msg1.UESecurityCapabilities)
	fmt.Println(msg1.SecurityKey)
	fmt.Println(msg1.MobilityRestrictionList)
	fmt.Println(msg1.MaskedIMEISV)
	fmt.Println(msg1.NASPDU)
}

var m2 []byte = []byte{0, 29, 0, 128, 200, 0, 0, 4, 0, 10, 0, 2, 0, 1, 0, 85, 0, 2, 0, 1, 0, 74, 0, 128, 166, 0, 64, 1, 103, 126, 2, 98, 205, 169, 177, 3, 126, 0, 104, 1, 0, 88, 46, 1, 1, 194, 17, 0, 35, 1, 0, 6, 49, 49, 1, 1, 255, 1, 1, 0, 14, 33, 17, 9, 16, 1, 1, 1, 1, 255, 255, 255, 255, 128, 3, 2, 0, 6, 33, 18, 1, 1, 255, 2, 6, 6, 3, 232, 6, 3, 232, 41, 5, 1, 10, 60, 0, 1, 34, 4, 1, 1, 2, 3, 121, 0, 12, 1, 32, 65, 1, 1, 9, 3, 32, 65, 1, 1, 8, 37, 9, 8, 105, 110, 116, 101, 114, 110, 101, 116, 18, 1, 64, 32, 1, 2, 3, 53, 0, 0, 4, 0, 130, 0, 10, 12, 59, 154, 202, 0, 48, 59, 154, 202, 0, 0, 139, 0, 10, 1, 240, 192, 168, 56, 101, 0, 0, 0, 2, 0, 134, 0, 1, 0, 0, 136, 0, 13, 4, 1, 0, 0, 9, 28, 0, 48, 0, 0, 8, 28, 0, 0, 110, 64, 10, 12, 119, 53, 148, 0, 48, 59, 154, 202, 0}

func TestM2(t *testing.T) {
	var pdu NgapPdu
	var err error
	var diagnostics *ies.CriticalityDiagnostics
	pdu, err, diagnostics = NgapDecode(m2)
	if err != nil {
		fmt.Println(err)
		fmt.Println(diagnostics)
		return
	}
	msg1 := pdu.Message.Msg.(*ies.PDUSessionResourceSetupRequest)
	fmt.Println(msg1.PDUSessionResourceSetupListSUReq[0].PDUSessionResourceSetupRequestTransfer)

	fmt.Println("-------------------------------------")
	fmt.Println("-------------------------------------")
	fmt.Println("-------------------------------------")

	msg2 := ies.PDUSessionResourceSetupRequest{
		AMFUENGAPID: 1,
		RANUENGAPID: 1,
		PDUSessionResourceSetupListSUReq: []ies.PDUSessionResourceSetupItemSUReq{{
			PDUSessionID: 1,
			// PDUSessionNASPDU:                       []byte{1, 1, 1},
			SNSSAI: ies.SNSSAI{SST: []byte{1}, SD: []byte{1, 2, 3}},
			// PDUSessionResourceSetupRequestTransfer: []byte{2, 2, 2},
		}},
		UEAggregateMaximumBitRate: &ies.UEAggregateMaximumBitRate{
			UEAggregateMaximumBitRateDL: 1000000000,
			UEAggregateMaximumBitRateUL: 2000000000,
		},
	}
	var b []byte
	b, err = NgapEncode(&msg2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b)
}

func TestTransfer1(t *testing.T) {
	fmt.Println(getPDUSessionResourceSetupResponseTransfer("192.168.1.1", 1, 1))
}

func getPDUSessionResourceSetupResponseTransfer(ipv4 string, teid uint32, qosId int64) []byte {
	data := ies.PDUSessionResourceSetupResponseTransfer{}

	dowlinkTeid := make([]byte, 4)
	binary.BigEndian.PutUint32(dowlinkTeid, teid)
	ipNgap := utils.IPAddressToNgap(ipv4, "")
	data.DLQosFlowPerTNLInformation = ies.QosFlowPerTNLInformation{
		UPTransportLayerInformation: ies.UPTransportLayerInformation{
			Choice: ies.UPTransportLayerInformationPresentGtptunnel,
			GTPTunnel: &ies.GTPTunnel{
				GTPTEID:               dowlinkTeid,
				TransportLayerAddress: ipNgap,
			},
		},
		AssociatedQosFlowList: []ies.AssociatedQosFlowItem{
			ies.AssociatedQosFlowItem{QosFlowIdentifier: qosId},
		},
	}

	var buf []byte
	var err error
	if buf, err = data.Encode(); err != nil {
		fmt.Println(err)
		return nil
	}
	return buf
}
