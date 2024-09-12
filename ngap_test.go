package ngap

import (
	"bytes"
	"fmt"
	"io"
	"ngap/ie"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEncode(t *testing.T) {
	NgSetupRequest_file, err := os.Open("./test_msg/NgSetupRequest.bin")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// Đảm bảo đóng tệp sau khi sử dụng
	defer NgSetupRequest_file.Close()

	for _, pdu := range tests {
		t.Run(pdu.name, func(t *testing.T) {
			t.Parallel()
			logrus.Println("-------------", pdu.resultPdu)
			encode, err := NgapEncode(*pdu.resultPdu)
			if err != nil {
				t.Errorf("NgapEncode() NGSetupRequest fail = %v", err)
			} else {
				t.Logf("Encoded: %0b", encode)
			}
		})
	}
}

var tests = []struct {
	name      string
	buf       io.Reader
	resultPdu *NgapPdu
	check     bool
}{
	{
		name: "NgSetupRequest",
		buf:  bytes.NewReader([]byte{}),
		resultPdu: &NgapPdu{
			Present: NgapPduInitiatingMessage,
			Message: NgapMessage{
				ProcedureCode: ProcedureCode(ie.ProcedureCodeNGSetup),
				Criticality:   ie.Criticality{Value: ie.CriticalityPresentReject},
				Msg: &NGSetupRequest{
					RanNodeName: []byte{1},
					DefaultPagingDrx: &ie.PagingDrx{PagingDrx: []byte{2}},
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
