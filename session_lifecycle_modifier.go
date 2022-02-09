package govmq

type AuthOnRegisterModifier struct {
	MountPoint               string      `json:"mountpoint,omitempty"`
	MaxMessageSize           int         `json:"max_message_size,omitempty"`
	SharedSubscriptionPolicy interface{} `json:"shared_subscription_policy,omitempty"`
	MaxOnlineMessages        int         `json:"max_online_messages,omitempty"`
	MaxOfflineMessages       int         `json:"max_offline_messages,omitempty"`
}

type AuthOnRegisterM3Modifier struct {
	AuthOnRegisterModifier `json:",inline"`
	CleanSession           bool `json:"clean_session,omitempty"`
}

type AuthOnRegisterM5Modifier struct {
	AuthOnRegisterModifier `json:",inline"`
	CleanStart             bool `json:"clean_start,omitempty"`
	Properties             struct {
		SessionExpiryInterval int `json:"p_session_expiry_interval,omitempty"`
	} `json:"properties"`
}
