package hook

type AuthOnSubscribeWH struct {
	BaseWebHook
	Username string `json:"username,omitempty"`
}
