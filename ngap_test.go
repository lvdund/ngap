package ngap

import (
	"fmt"
	"os"
	"testing"
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
			msg := decode.Message.Msg.(*NGSetupRequest)
			fmt.Println("Message", msg)
		}
	}
}

/*
func TestEncode(t *testing.T) {
	for _, testCase := range tests {
		msg, _ := testCase.resultPdu.Message.Msg.(NgapMessageEncoder)
		if encoded, err := NgapEncode(msg); err != nil {
			t.Errorf("NgapEncode() NGSetupRequest fail = %v", err)
			return
		} else if !bytes.Equal(encoded, testCase.buf) {
			t.Errorf("Final buffer = % X\n, want %.8b\n", encoded, testCase.buf)
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
			fmt.Println()
			fmt.Println("Present:", decode.Present)
			fmt.Println("ProcedureCode:", decode.Message.ProcedureCode.Value)
			fmt.Println("Criticality:", decode.Message.Criticality.Value)
			msg := decode.Message.Msg.(*NGSetupRequest)
			fmt.Println("Message", msg)
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
	{
		name: "NgSetupRequest",
		buf:  []byte{0x00, 0x15, 0x00, 0x0F, 0x00, 0x00, 0x01, 0x00, 0x66, 0x00, 0x08, 0x00, 0x02, 0x02, 0xF8, 0x39, 0x02, 0xF8, 0x39},
		resultPdu: &NgapPdu{
			Present: NgapPduInitiatingMessage,
			Message: NgapMessage{
				ProcedureCode: ie.ProcedureCode{Value: aper.Integer(ie.ProcedureCodeNGSetup)},
				Criticality:   ie.Criticality{Value: ie.CriticalityPresentReject},
				Msg: &NGSetupRequest{
					SupportedTaList: &ie.SupportedTAList{
						List: []*ie.SupportedTaItem{
							&ie.SupportedTaItem{
								Tac: ie.Tac{Tac: aper.OctetString{
									0x02, 0xf8, 0x39,
								}},
							},
							&ie.SupportedTaItem{
								Tac: ie.Tac{Tac: aper.OctetString{
									0x02, 0xf8, 0x39,
								}},
							},
						},
					},
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
*/
