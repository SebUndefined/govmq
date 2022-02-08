package govmq

type AuthOnRegisterResponse struct {
	MountPoint     string `json:"mountpoint"`
	MaxMessageSize int    `json:"max_message_size"`
}

type AuthOnRegisterM3Response struct {
	AuthOnRegisterResponse `json:",inline"`
	CleanSession           bool `json:"clean_session"`
}

type AuthOnRegisterM5Response struct {
	AuthOnRegisterResponse `json:",inline"`
	CleanStart             bool `json:"clean_start"`
}
