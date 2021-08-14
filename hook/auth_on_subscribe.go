package hook

type AuthOnSubscribe struct {
	BaseWebHook
	Username string `json:"username,omitempty"`
}
