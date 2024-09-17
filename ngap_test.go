package ngap

import (
	"fmt"
	"ngap/aper"
	"ngap/ie"
	"reflect"
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
		if err != nil {
			fmt.Printf("NgapEncode() NGSetupRequest fail = %v", err)
			return
		} else if !reflect.DeepEqual(encode.GetBuf(), pdu.buf) {
			fmt.Printf("Encoded compare err: \n\thas: % X\n\twant: % X", encode.GetBuf(), pdu.buf)
		} else {
			fmt.Println("Decode reflect")
		}
	}
}

var tests = []struct {
	name      string
	buf       []byte
	resultPdu *NgapPdu
	check     bool
}{
	{
		name: "NgSetupRequest",
		buf:  []byte{0x00, 0x15, 0x00, 0x03, 0x00, 0x00, 0x00},
		resultPdu: &NgapPdu{
			Present: NgapPduInitiatingMessage,
			Message: NgapMessage{
				ProcedureCode: ie.ProcedureCode{Value: aper.Integer(ie.ProcedureCodeNGSetup)},
				Criticality:   ie.Criticality{Value: ie.CriticalityPresentReject},
				Msg:           &NGSetupRequest{},
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
