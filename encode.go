package ngap

import "bytes"

func NgapEncode(msg NgapMessageEncoder) (wire []byte, err error) {
	var buf bytes.Buffer
	if err = msg.Encode(&buf); err == nil {
		wire = buf.Bytes()
	}
	return
}
