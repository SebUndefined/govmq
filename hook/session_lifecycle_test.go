package hook

import (
	"encoding/json"
	"fmt"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	body := "{\"peer_addr\":\"127.0.0.1\",\"peer_port\":8888,\"username\":\"username\",\"password\":\"password\",\"mountpoint\":\"\",\"client_id\":\"clientid\",\"clean_session\":false}"
	wh := AuthOnRegister{}
	err := json.Unmarshal([]byte(body), &wh)
	if err != nil {
		t.Fatalf(`The body %s cannot be parsed. Error: %s`, body, err)
	}
	fmt.Println(wh)
}
