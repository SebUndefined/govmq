package hook

import (
	"encoding/json"
	"testing"
)

func TestAuthOnRegisterSuccess(t *testing.T) {
	body := "{\"peer_addr\":\"127.0.0.1\",\"peer_port\":8888,\"username\":\"username\",\"password\":\"password\",\"mountpoint\":\"\",\"client_id\":\"clientid\",\"clean_session\":false}"
	wh := AuthOnRegister{}
	err := json.Unmarshal([]byte(body), &wh)
	if err != nil {
		t.Fatalf(`The body %s cannot be parsed. Error: %s`, body, err)
	}
}

func TestAuthOnRegisterFail(t *testing.T) {
	body := "{\"peer_addr\":\"127.0.0.1\",\"peer_port\":8888,\"username\":1234,\"password\":\"password\",\"mountpoint\":\"\",\"client_id\":\"clientid\",\"clean_session\":false}"
	wh := AuthOnRegister{}
	err := json.Unmarshal([]byte(body), &wh)
	if err == nil {
		t.Fatalf("No error raised.")
	}
}

func TestTest(t *testing.T) {
	body := "{\"client_id\":\"clientid\",\"mountpoint\":\"\"}"
	wh := OnClientWakeUp{}
	err := json.Unmarshal([]byte(body), &wh)
	if err != nil {
		t.Fatalf(`The body %s cannot be parsed. Error: %s`, body, err)
	}
}
