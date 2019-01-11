package codec

import (
	"testing"
)

type reqNuLogin struct {
	TXN                 string `fesl:"TXN"`
	ReturnEncryptedInfo bool   `fesl:"returnEncryptedInfo"`
	MacAddr             string `fesl:"macAddr,omitempty"`
}

type reqNuLoginServer struct {
	reqNuLogin

	AccountName     string `fesl:"nuid"`
	AccountPassword string `fesl:"password"`
}

type reqNuLoginClient struct {
	reqNuLogin

	EncryptedInfo string `fesl:"encryptedInfo"`
}

func TestEncode(t *testing.T) {
	pkt := reqNuLoginClient{
		reqNuLogin: reqNuLogin{
			TXN:                 "NuLogin",
			ReturnEncryptedInfo: false,
		},
		EncryptedInfo: "1234",
	}

	enc := NewEncoder()
	err := enc.Encode(pkt)
	if err != nil {
		t.Fatal(err)
	}

	f := DecodeFESL(enc.wr.Bytes())
	expect(t, f, "TXN", "NuLogin")
	expect(t, f, "returnEncryptedInfo", "0")
	expect(t, f, "encryptedInfo", "1234")
	expectNotSet(t, f, "macAddr")
}

func expect(t *testing.T, v Fields, key, value string) {
	val, ok := v[key]
	if !ok {
		t.Fatalf("Expected to see key %s", key)
	}
	if val != value {
		t.Fatalf("Unexpected value, %s; should be %s", val, value)
	}
}

func expectNotSet(t *testing.T, v Fields, key string) {
	if _, ok := v[key]; ok {
		t.Fatalf("Not expected key %s to be set", key)
	}
}
