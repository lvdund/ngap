package ngap

import (
	"bytes"
	"fmt"
	"io"
	"ngap/ie"
	"os"
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	NgSetupRequest_file, err := os.Open("./test_msg/NgSetupRequest.bin")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// Đảm bảo đóng tệp sau khi sử dụng
	defer NgSetupRequest_file.Close()
	tests := []struct {
		name       string
		buf       io.Reader
		resultPdu *NgapPdu
		check     bool
	}{
		{
			name: "NgSetupRequest",
			buf: bytes.NewReader([]byte{}),
			resultPdu: &NgapPdu{
				Present: NgapPduInitiatingMessage,
				Message: NgapMessage{
					ProcedureCode: ProcedureCode(ie.ProcedureCodeNGSetup),
					Criticality:   ie.Criticality{Value: ie.CriticalityPresentReject},
					Msg: &NGSetupRequest{
						RanNodeName: []byte{1},
					},
				},
			},
		},
		{
			name: "NgSetupRequest from UERANSIM",
			buf: NgSetupRequest_file,
			resultPdu: &NgapPdu{
				Present: NgapPduInitiatingMessage,
				Message: NgapMessage{
					ProcedureCode: ProcedureCode(ie.ProcedureCodeNGSetup),
					Criticality:   ie.Criticality{Value: ie.CriticalityPresentReject},
					Msg: &NGSetupRequest{
					},
				},
			},
		},
	}

	for _, pdu := range tests {
		t.Run(pdu.name, func(t *testing.T) {
			t.Parallel()
			buff, _ := io.ReadAll(pdu.buf)
			decodePdu, err, _ := NgapDecode(buff)
			if err != nil {
				t.Errorf("NgapDecode() NGSetupRequest fail = %v", err)
			}
			if !reflect.DeepEqual(decodePdu, pdu.resultPdu) {
				t.Errorf("NgapDecode() = %v, want %v", decodePdu, pdu.resultPdu)
			}
			
		})
	}
}
