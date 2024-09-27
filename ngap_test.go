package ngap

import (
	"bytes"
	"fmt"
	"ngap/aper"
	"ngap/ie"

	//"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	// NgSetupRequest_file, err := os.Open("./test_msg/NgSetupRequest.bin")
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	return
	// }
	// defer NgSetupRequest_file.Close()

	for _, pdu := range tests {
		encode, err := NgapEncode(*pdu.resultPdu)
		_ = encode
		if err != nil {
			t.Errorf("NgapEncode() NGSetupRequest fail = %v", err)
			return
		} else if !bytes.Equal(encode.GetBuf(), pdu.buf) {
			t.Errorf("Final buffer = %.8b\n, want %.8b\n", encode.GetBuf(), pdu.buf)
		}
	}
	// now: err when encode octetstring - RanNodeName
}

func TestDecode(t *testing.T) {
	for _, pdu := range tests {
		// decode, err, _ := NgapDecode(pdu.buf)
		decode, err := HandleNgap(pdu.buf)
		if err != nil {
			t.Errorf("NgapDecode() NGSetupRequest fail = %v", err)
		} else {
			fmt.Println(decode)
		}
	}
}

var tests = []struct {
	name      string
	buf       []byte
	resultPdu *NgapPdu
	check     bool
}{
	// buf - NgSetupRequest from UERANSIM with RanNodeName == "a"
	{
		name: "NgSetupRequest",
		buf:  []byte{0x00, 0x15, 0x00, 0x0A, 0x00, 0x00, 0x01, 0x00, 0x52, 0x40, 0x03, 0x00, 0x00, 0x61},
		resultPdu: &NgapPdu{
			Present: NgapPduInitiatingMessage,
			Message: NgapMessage{
				ProcedureCode: ie.ProcedureCode{Value: aper.Integer(ie.ProcedureCodeNGSetup)},
				Criticality:   ie.Criticality{Value: ie.CriticalityPresentReject},
				Msg: &NGSetupRequest{
					RanNodeName: &ie.RANNodeName{Value: "a"},
				},
			},
		},
	},
	// {
	// 	name: "NgSetupRequest from UERANSIM",
	// 	buf:  NgSetupRequest_file,
	// 	resultPdu: &NgapPdu{
	// 		Present: NgapPduInitiatingMessage,
	// 		Message: NgapMessage{
	// 			ProcedureCode: ProcedureCode(ie.ProcedureCodeNGSetup),
	// 			Criticality:   ie.Criticality{Value: ie.CriticalityPresentReject},
	// 			Msg:           &NGSetupRequest{},
	// 		},
	// 	},
	// },
}
