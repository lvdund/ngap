package ngap

import (
	"ngap/ie"

	"github.com/sirupsen/logrus"
)

// This is to describe how NGAP handler should be written
func HandleNgap(wire []byte) (pdu NgapPdu, err error) {
	var diagnostics *ie.CriticalityDiagnostics
	if pdu, err, diagnostics = NgapDecode(wire); err != nil {
		return
	} else {
		if diagnostics != nil {
			//TODO: send error indication
		}
	}

	switch pdu.Present {
	case NgapPduInitiatingMessage:
		handleInitiatingMessage(&pdu)
	case NgapPduSuccessfulOutcome:
		handleSuccessfulOutcome(&pdu)
	case NgapPduUnsuccessfulOutcome:
	}
	return
}

func handleInitiatingMessage(pdu *NgapPdu) {
	logrus.Infoln("handleInitiatingMessage -", pdu.Message.ProcedureCode.Value)
	switch int64(pdu.Message.ProcedureCode.Value) {
	case ie.ProcedureCodeNGSetup: //NGSetupRequest
		msg := pdu.Message.Msg.(*NGSetupRequest)
		//then handle the message
		handleNGSetupRequest(msg)
	default:
		//TODO:
	}
}

func handleSuccessfulOutcome(pdu *NgapPdu) {
	switch int64(pdu.Message.ProcedureCode.Value) {
	case ie.ProcedureCodeNGSetup: //NGSetupResponse
		// msg := pdu.Message.Msg(*NGSetupResponse)
		//then handle the message
		// handleNGSetupResponse(msg)
	default:
		//TODO:
	}
}

func handleNGSetupRequest(msg *NGSetupRequest) {
	logrus.Infoln("handleNGSetupRequest")
	//TODO: add example code
	
}

// func handleNGSetupResponse(msg *NGSetupResponse) {
// 	//TODO: add example code
// }
