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

func TestDecodeUeRanSim(t *testing.T) {
	if f, err := os.Open("./test_msg/ngsetuprequest.bin"); err != nil {
		t.Errorf("Fail to read file : %+v", err)
	} else {
		defer f.Close()
		if decode, err, _ := NgapDecode(f); err != nil {
			t.Errorf("NgapDecode() NGSetupRequest fail = %v", err)
		} else {
			_ = decode
			fmt.Println()
			fmt.Println("Present:", decode.Present)
			fmt.Println("ProcedureCode:", decode.Message.ProcedureCode.Value)
			fmt.Println("Criticality:", decode.Message.Criticality.Value)
			msg := decode.Message.Msg.(*ies.NGSetupRequest)
			fmt.Printf("Message rannodename %s\n", msg.RANNodeName.Value)
			fmt.Println("Message GlobalRanNodeId:", msg.GlobalRANNodeID)
			mcc, mnc := utils.PlmnidToMccMnc(msg.GlobalRANNodeID.GlobalGNBID.PLMNIdentity.Value)
			fmt.Println("Message PLMNIdentity:", mcc, mnc, msg.GlobalRANNodeID.GlobalGNBID.PLMNIdentity.Value)
			fmt.Println("Message GNBID Choice:", msg.GlobalRANNodeID.GlobalGNBID.GNBID.Choice)
			if msg.GlobalRANNodeID.GlobalGNBID.GNBID.GNBID != nil {
				fmt.Println("Message GNBID:", msg.GlobalRANNodeID.GlobalGNBID.GNBID.GNBID)
			}
			fmt.Println("Message GNBID:", msg.GlobalRANNodeID.GlobalGNBID.GNBID.GNBID)
		}
	}
}

func TestNil(t *testing.T) {
	msg := ies.AMFConfigurationUpdateAcknowledge{}
	o, err := NgapEncode(&msg)
	fmt.Println(o, err)
}

func TestEncodeDecode(t *testing.T) {
	msg, _ := test.resultPdu.Message.Msg.(NgapMessageEncoder)
	var encoded []byte
	var err error
	if encoded, err = NgapEncode(msg); err != nil {
		t.Errorf("NgapEncode() NGSetupRequest fail = %v", err)
		return
	}
	fmt.Println()
	if decode, err, _ := NgapDecode(bytes.NewBuffer(encoded)); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println()
		fmt.Println("Present:", decode.Present)
		fmt.Println("ProcedureCode:", decode.Message.ProcedureCode.Value)
		fmt.Println("Criticality:", decode.Message.Criticality.Value)
		msg := decode.Message.Msg.(*ies.NGSetupRequest)
		fmt.Printf("Message rannodename %s\n", msg.RANNodeName.Value)
		mcc, mnc := utils.PlmnidToMccMnc(msg.GlobalRANNodeID.GlobalGNBID.PLMNIdentity.Value)
		fmt.Println("Message PLMNIdentity:", mcc, mnc, msg.GlobalRANNodeID.GlobalGNBID.PLMNIdentity.Value)
		fmt.Println("Message N3IWFID choice:", msg.GlobalRANNodeID.GlobalGNBID.GNBID.Choice)
		fmt.Println("Message GNBID Bytes:", msg.GlobalRANNodeID.GlobalGNBID.GNBID.GNBID.Bytes)
		fmt.Println("Message GNBID NumBits:", msg.GlobalRANNodeID.GlobalGNBID.GNBID.GNBID.NumBits)
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
				GlobalRANNodeID: &ies.GlobalRANNodeID{
					Choice: ies.GlobalRANNodeIDPresentGlobalGNBID,
					GlobalGNBID: &ies.GlobalGNBID{
						PLMNIdentity: &ies.PLMNIdentity{Value: aper.OctetString{0x02, 0xf8, 0x39}},
						GNBID: &ies.GNBID{
							Choice: 1,
							GNBID: &aper.BitString{
								Bytes:   []byte{0x45, 0x46, 0x47},
								NumBits: 24,
							},
						},
					},
				},
			},
		},
	},
}
