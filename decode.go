package ngap

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/lvdund/ngap/ies"
)

// decode a Ngap message from io.Reader
func NgapDecode(buf []byte) (pdu NgapPdu, err error, diagnostics *ies.CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewBuffer(buf))
	//1. decode extention bit
	var b bool
	if b, err = r.ReadBool(); err != nil {
		return
	}
	_ = b
	//2. decode present		//choice among InitiatingMessage, SuccessfulOutcome and UnsuccessfulOutcome
	c, err := r.ReadChoice(2, false)
	if err != nil {
		return
	}
	present := uint8(c)
	//3. decode procedure code
	v, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false)
	if err != nil {
		return
	}
	var procedureCode ies.ProcedureCode = ies.ProcedureCode{Value: aper.Integer(v)}
	//4. decode criticality
	e, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return
	}
	var criticality ies.Criticality = ies.Criticality{Value: aper.Enumerated(e)}
	//5. decode message content
	var containerBytes []byte
	if containerBytes, err = r.ReadOpenType(); err != nil {
		return
	}

	//prepare message for decoding
	message := createMessage(present, procedureCode)
	if message == nil {
		err = fmt.Errorf("Unknown message") //TODO: create a right error message
		return
	}

	var diagnosticsItems []ies.CriticalityDiagnosticsIEItem
	//decode all IEs within the message
	if err, diagnosticsItems = message.Decode(containerBytes); err != nil {
		return
	}

	pdu = NgapPdu{
		Present: present,
		Message: NgapMessage{
			ProcedureCode: procedureCode,
			Criticality:   criticality,
			Msg:           message,
		},
	}

	//in case there was any critical diagnostics, create a report
	diagnostics = ies.BuildDiagnostics(present, procedureCode, criticality, diagnosticsItems)
	return
}

func TransferDecode(ioR io.Reader) (pdu NgapPdu, err error, diagnostics *ies.CriticalityDiagnostics) {
	r := aper.NewReader(ioR)
	if _, err = r.ReadBool(); err != nil {
		return
	}
	return
}
