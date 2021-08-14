package hook

import (
	"fmt"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	var test = &AuthOnSubscribe{
		BaseWebHook: BaseWebHook{
			ClientId:   "ClientIdTest",
			MountPoint: "MountPointTest",
		},
		Username: "UsernameTest"}
	fmt.Println(test.ToJson())
}
