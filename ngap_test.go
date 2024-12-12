package ngap

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/lvdund/ngap/aper"
	"github.com/lvdund/ngap/ies"
	"github.com/lvdund/ngap/utils"
)

func TestOptional(t *testing.T) {
	optionals := []byte{0x0}
	aper.SetBit(optionals, 1)
	fmt.Printf("%b\n\n", optionals)

	s := ies.AreaOfInterest{
		AreaOfInterestTAIList: &ies.AreaOfInterestTAIList{Value: []*ies.AreaOfInterestTAIItem{&ies.AreaOfInterestTAIItem{
			TAI: &ies.TAI{
				PLMNIdentity: &ies.PLMNIdentity{Value: aper.OctetString{0x02, 0xf8, 0x39}},
				TAC:          &ies.TAC{Value: aper.OctetString{0x02, 0xf8, 0x39}},
			},
		}}},
		AreaOfInterestRANNodeList: &ies.AreaOfInterestRANNodeList{Value: []*ies.AreaOfInterestRANNodeItem{&ies.AreaOfInterestRANNodeItem{
			GlobalRANNodeID: &ies.GlobalRANNodeID{
				Choice: ies.GlobalRANNodeIDPresentGlobalGNBID,
				GlobalGNBID: &ies.GlobalGNBID{
					PLMNIdentity: &ies.PLMNIdentity{Value: aper.OctetString{0x02, 0xf8, 0x39}},
					GNBID: &ies.GNBID{
						Choice: ies.GNBIDPresentGNBID,
						GNBID: &aper.BitString{
							Bytes:   []byte{0x45, 0x46, 0x47},
							NumBits: 24,
						},
					},
				},
			},
		}}},
	}

	var buf bytes.Buffer
	w := aper.NewWriter(&buf)
	if err := s.Encode(w); err != nil {
		fmt.Println("err:", err)
	}
	aper.FlushWrite(w)
	fmt.Printf("encode: %b\n", buf.Bytes())

	// decode
	ss := ies.AreaOfInterest{}
	if err := ss.Decode(aper.NewReader(bytes.NewBuffer(buf.Bytes()))); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Decode:", ss)
		fmt.Println("PLMNIdentity", ss.AreaOfInterestTAIList.Value[0].TAI.PLMNIdentity.Value)
		fmt.Println("TAC", ss.AreaOfInterestTAIList.Value[0].TAI.TAC.Value)
		fmt.Println("GlobalRANNodeID Choice", ss.AreaOfInterestRANNodeList.Value[0].GlobalRANNodeID.Choice)
		fmt.Println("\tPLMNIdentity Choice", ss.AreaOfInterestRANNodeList.Value[0].GlobalRANNodeID.GlobalGNBID.PLMNIdentity.Value)
		fmt.Println("\tGlobalGNBID Choice", ss.AreaOfInterestRANNodeList.Value[0].GlobalRANNodeID.GlobalGNBID.GNBID.Choice)
		fmt.Println("\t\tGNBID", ss.AreaOfInterestRANNodeList.Value[0].GlobalRANNodeID.GlobalGNBID.GNBID.GNBID)
	}
}

func TestDecodeUeRanSim(t *testing.T) {
	if f, err := os.Open("./test_msg/ngsetuprequest.bin"); err != nil {
		t.Errorf("Fail to read file : %+v", err)
	} else {
		defer f.Close()
		if decode, err, _ := NgapDecode(f); err != nil {
			t.Errorf("NgapDecode() NGSetupRequest fail = %v", err)
		} else {
			msg := decode.Message.Msg.(*ies.NGSetupRequest)
			// GlobalRANNodeID
			fmt.Println("GlobalGNBID - GNBID:", msg.GlobalRANNodeID.GlobalGNBID.GNBID.Choice,
				msg.GlobalRANNodeID.GlobalGNBID.GNBID.GNBID)
			mcc, mnc := utils.PlmnidToMccMnc(msg.GlobalRANNodeID.GlobalGNBID.PLMNIdentity.Value)
			fmt.Println("GlobalGNBID - PLMNIdentity:", mcc, mnc, msg.GlobalRANNodeID.GlobalGNBID.PLMNIdentity.Value)
			// RANNodeName
			fmt.Printf("RANNodeName %s\n", msg.RANNodeName.Value)
		}
	}
}

func TestNil(t *testing.T) {
	msg := ies.AMFConfigurationUpdateAcknowledge{}
	o, err := NgapEncode(&msg)
	fmt.Println(o, err)
}

func TestEncodeDecode(t *testing.T) {
	if decode, err, _ := NgapDecode(bytes.NewBuffer(ngsetupreq)); err != nil {
		fmt.Println(err)
	} else {
		msg := decode.Message.Msg.(*ies.NGSetupRequest)

		fmt.Printf("Message rannodename %s\n", msg.RANNodeName.Value)

		mcc, mnc := utils.PlmnidToMccMnc(msg.GlobalRANNodeID.GlobalGNBID.PLMNIdentity.Value)
		fmt.Println("GlobalGNBID PLMNIdentity:", mcc, mnc, msg.GlobalRANNodeID.GlobalGNBID.PLMNIdentity.Value)

		fmt.Println("GlobalGNBID GNBID Bytes:", msg.GlobalRANNodeID.GlobalGNBID.GNBID.GNBID.Bytes)
		fmt.Println("GlobalGNBID GNBID NumBits:", msg.GlobalRANNodeID.GlobalGNBID.GNBID.GNBID.NumBits)

		fmt.Println("SupportedTAList TAC:", msg.SupportedTAList.Value[0].TAC.Value)
		fmt.Println("SupportedTAList PLMNIdentity:", msg.SupportedTAList.Value[0].BroadcastPLMNList.Value[0].PLMNIdentity)
		snssai := msg.SupportedTAList.Value[0].BroadcastPLMNList.Value[0].TAISliceSupportList.Value[0].SNSSAI
		fmt.Println("SNSSAI:", snssai.SST.Value, snssai.SD.Value)

		fmt.Println("PagingDRX", msg.DefaultPagingDRX.Value)
	}
}

func BenchmarkPerformace(b *testing.B) {
	msg := test.resultPdu
	out := test.buf
	for i := 0; i < b.N; i++ {
		if _, err, _ := NgapDecode(bytes.NewBuffer(out)); err != nil {
			fmt.Println("Decode:", err)
			return
		}
		if _, err := NgapEncode(msg.Message.Msg.(NgapMessageEncoder)); err != nil {
			fmt.Println("Decode:", err)
			return
		}
	}
}
func TestNgap(t *testing.T) {
	msg := test.resultPdu
	out := test.buf
	if _, err, _ := NgapDecode(bytes.NewBuffer(out)); err != nil {
		fmt.Println("Decode:", err)
		return
	}
	if b , err := NgapEncode(msg.Message.Msg.(NgapMessageEncoder)); err != nil {
		fmt.Println("Decode:", err)
		return
	} else {
		fmt.Printf("encode: %0b\n\t%v\n", b, b)
	}
}

var str string = "a"
var oct aper.OctetString = aper.OctetString(str)

var test = struct {
	name      string
	buf       []byte
	resultPdu *NgapPdu
	check     bool
}{
	name: "NgSetupRequest",
	buf:  []byte{0x00, 0x15, 0x00, 0x0A, 0x00, 0x00, 0x01, 0x00, 0x52, 0x40, 0x03, 0x00, 0x00, 0x61},
	resultPdu: &NgapPdu{
		Present: ies.NgapPduInitiatingMessage,
		Message: NgapMessage{
			ProcedureCode: ies.ProcedureCode{Value: aper.Integer(ies.ProcedureCode_NGSetup)},
			Criticality:   ies.Criticality{Value: ies.Criticality_PresentReject},
			Msg: &ies.NGSetupRequest{
				RANNodeName: &ies.RANNodeName{Value: oct},
			},
		},
	},
}

var ngsetupreq []byte = []byte{0, 21, 0, 59, 0, 0, 4, 0, 27, 0, 8, 0, 2, 248, 57, 16, 69, 70, 71, 0, 82, 64, 15, 6, 0, 109, 121, 53, 103, 82, 65, 78, 84, 101, 115, 116, 101, 114, 0, 102, 0, 16, 0, 0, 0, 0, 1, 0, 2, 248, 57, 0, 0, 16, 8, 1, 2, 3, 0, 21, 64, 1, 64}
