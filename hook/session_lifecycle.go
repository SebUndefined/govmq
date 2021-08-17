package hook

type AuthOnRegister struct {
	MountPoint   string `json:"mountpoint,omitempty"`
	ClientId     string `json:"client_id,omitempty"`
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	PeerAddress  string `json:"peer_address,omitempty"`
	PeerPort     int    `json:"peer_port,omitempty"`
	CleanSession bool   `json:"clean_session,omitempty"`
}
