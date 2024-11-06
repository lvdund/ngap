package ngap

import (
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/lvdund/ngap/ies"
)

// decode a Ngap message from io.Reader
func NgapDecode(ioR io.Reader) (pdu NgapPdu, err error, diagnostics *ies.CriticalityDiagnostics) {
	r := aper.NewReader(ioR)
	//1. decode extention bit
	var b bool
	if b, err = r.ReadBool(); err != nil {
		return
	}
	_ = b
	// fmt.Println("extenstion bit:", b)
	//2. decode present		//choice among InitiatingMessage, SuccessfulOutcome and UnsuccessfulOutcome
	// v, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 2}, false)
	c, err := r.ReadChoice(2, false)
	if err != nil {
		return
	}
	present := uint8(c)
	// fmt.Println("present:", present)
	//3. decode procedure code
	v, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false)
	if err != nil {
		return
	}
	// fmt.Printf("procedureCode:%d\n", v)
	var procedureCode ies.ProcedureCode = ies.ProcedureCode{Value: aper.Integer(v)}
	//4. decode criticality
	e, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return
	}
	var criticality ies.Criticality = ies.Criticality{Value: aper.Enumerated(e)}
	// fmt.Println("criticality:", e)
	//5. decode message content
	var containerBytes []byte
	if containerBytes, err = r.ReadOpenType(); err != nil {
		return
	}
	//fmt.Printf("exten=%v - containerBytes readopen type= %.8b\n", b, containerBytes)

	//prepare message for decoding
	message := createMessage(present, procedureCode)
	if message == nil {
		err = fmt.Errorf("Unknown message") //TODO: create a right error message
		return
	}

	var diagnosticsItems []ies.CriticalityDiagnostics
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
