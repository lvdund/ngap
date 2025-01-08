package ngap

import (
	"io"

	"github.com/lvdund/ngap/ies"
)

// hold a decoded Ngap message
type NgapPdu struct {
	Present uint8 //choice among InitiatingMessage, SuccessfulOutcome and UnsuccessfulOutcome
	Message NgapMessage
}

// represent InitiatingMessage, SuccessfulOutcome or UnsuccessfulOutcome
type NgapMessage struct {
	ProcedureCode ies.ProcedureCode
	Criticality   ies.Criticality
	Msg           MessageUnmarshaller //to be decoded message
}

// interface to message decoder all message need to implement this interface
type MessageUnmarshaller interface {
	Decode([]byte) (error, []ies.CriticalityDiagnosticsIEItem)
}

type NgapMessageEncoder interface {
	Encode(io.Writer) error
}
