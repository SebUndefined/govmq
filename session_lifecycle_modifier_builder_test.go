package govmq

import (
	"encoding/json"
	"testing"
)

func TestAuthOnRegisterM3ModifierBuilder(t *testing.T) {
	b := AuthOnRegisterM3ResponseBuilder{}
	b.WithMaxMessageSize(65535).
		WithMaxOfflineMessages(10000).
		WithMountpoint("test")
	body := "{\"result\":\"ok\",\"modifiers\":{\"max_message_size\":65535,\"max_offline_messages\":10000,\"retry_interval\":\"test\"}}"
	var res []byte
	res, err := json.Marshal(b.Build())

	if err != nil {
		t.Fatalf(`The body %s cannot be parsed. Error: %s`, body, err)
	}
	if string(res) != body {
		t.Fatalf(`The body %s is different of %s`, body, string(res))
	}
}
