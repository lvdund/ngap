package ngap

import (
	"encoding/hex"
	"testing"

	"github.com/lvdund/ngap/ies"
)

// A message with an empty protocolIEs container - free5gc/ngap's encoding of
// an NG Reset Acknowledge for the whole interface - decodes, and re-encodes
// to the same octets.
func TestMessageWithNoIEsRoundTrips(t *testing.T) {
	const wire = "20140003000000"
	b, _ := hex.DecodeString(wire)
	pdu, err, _ := NgapDecode(b)
	if err != nil {
		t.Fatalf("decode: %v", err)
	}
	if _, ok := pdu.Message.Msg.(*ies.NGResetAcknowledge); !ok {
		t.Fatalf("decoded %T, want NGResetAcknowledge", pdu.Message.Msg)
	}
	out, err := NgapEncode(pdu.Message.Msg.(*ies.NGResetAcknowledge))
	if err != nil {
		t.Fatalf("encode: %v", err)
	}
	if got := hex.EncodeToString(out); got != wire {
		t.Fatalf("re-encoded %s, want %s", got, wire)
	}
}
