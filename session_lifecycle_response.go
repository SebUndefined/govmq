package govmq

type AuthOnRegisterResponse struct {
	MountPoint               string      `json:"mountpoint"`
	MaxMessageSize           int         `json:"max_message_size"`
	SharedSubscriptionPolicy interface{} `json:"shared_subscription_policy"`
	MaxOnlineMessages        int         `json:"max_online_messages "`
	MaxOfflineMessages       int         `json:"max_offline_messages "`
}

type AuthOnRegisterM3Response struct {
	AuthOnRegisterResponse `json:",inline"`
	CleanSession           bool `json:"clean_session"`
}

type AuthOnRegisterM5Response struct {
	AuthOnRegisterResponse `json:",inline"`
	CleanStart             bool `json:"clean_start"`
	Properties             struct {
		SessionExpiryInterval int `json:"p_session_expiry_interval"`
	} `json:"properties"`
}
