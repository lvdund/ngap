package ies

import (
	"bytes"
	"encoding/hex"
	"testing"
)

// A message whose IEs are all optional and absent encodes with an empty
// protocolIEs container; the octets are what free5gc/ngap writes for it.
func TestMessageWithNoIEsEncodes(t *testing.T) {
	var buf bytes.Buffer
	if err := (&NGResetAcknowledge{}).Encode(&buf); err != nil {
		t.Fatal(err)
	}
	if got, want := hex.EncodeToString(buf.Bytes()), "20140003000000"; got != want {
		t.Fatalf("encoded %s, want %s", got, want)
	}
}

