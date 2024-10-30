package ngap

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/lvdund/ngap/aper"
	"github.com/lvdund/ngap/ies"
)

func TestDecodeUeRanSim(t *testing.T) {
	if f, err := os.Open("./test_msg/ngsetuprequest.bin"); err != nil {
		t.Errorf("Fail to read file : %+v", err)
	} else {
		defer f.Close()
		if decode, err, _ := NgapDecode(f); err != nil {
			t.Errorf("NgapDecode() NGSetupRequest fail = %v", err)
		} else {
			fmt.Println()
			fmt.Println("Present:", decode.Present)
			fmt.Println("ProcedureCode:", decode.Message.ProcedureCode.Value)
			fmt.Println("Criticality:", decode.Message.Criticality.Value)
			msg := decode.Message.Msg.(*ies.NGSetupRequest)
			fmt.Printf("Message rannodename %s\n", msg.RANNodeName.Value)
			fmt.Println("Message GlobalRanNodeId:", msg.GlobalRANNodeID)
			mcc, mnc := ies.PlmnidToMccMnc(msg.GlobalRANNodeID.GlobalGNBID.PLMNIdentity.Value)
			fmt.Println("Message PLMNIdentity:", mcc, mnc, msg.GlobalRANNodeID.GlobalGNBID.PLMNIdentity.Value)
		}
	}
}

func TestEncode(t *testing.T) {
	msg, _ := test.resultPdu.Message.Msg.(NgapMessageEncoder)
	if encoded, err := NgapEncode(msg); err != nil {
		t.Errorf("NgapEncode() NGSetupRequest fail = %v", err)
		return
	} else if !bytes.Equal(encoded, test.buf) {
		t.Errorf("Final buffer = % X\n, want %.8b\n", encoded, test.buf)
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
