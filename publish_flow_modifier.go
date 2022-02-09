package govmq

type AuthOnPublishModifier struct {
	MountPoint string `json:"mountpoint"`
	Retain     bool   `json:"retain"`
	Topic      string `json:"topic"`
	QoS        QoS    `json:"qos"`
	Throttle   int    `json:"throttle"`
	Payload    string `json:"payload"`
}

type AuthOnPublishM3Modifier struct {
	AuthOnPublishModifier `json:",inline"`
}

type AuthOnPublishM5Modifier struct {
	AuthOnPublishModifier `json:",inline"`
	Properties            struct {
		PayloadFormatIndicator string      `json:"p_payload_format_indicator"`
		ContentType            string      `json:"p_content_type"`
		MessageExpiryInterval  int         `json:"p_message_expiry_interval"`
		UserProperty           interface{} `json:"p_user_property"`
		ResponseTopic          string      `json:"p_response_topic"`
		CorrelationData        string      `json:"p_correlation_data"`
	} `json:"properties"`
}

// On deliver

type OnDeliverModifier struct {
	Retain  bool   `json:"retain"`
	Topic   string `json:"topic"`
	QoS     QoS    `json:"qos"`
	Payload string `json:"payload"`
}

type OnDeliverM3Modifier struct {
	AuthOnRegisterModifier `json:",inline"`
}

type OnDeliverM5Modifier struct {
	AuthOnRegisterModifier `json:",inline"`
	CleanStart             bool `json:"clean_start"`
	Properties             struct {
		PayloadFormatIndicator string      `json:"p_payload_format_indicator"`
		ContentType            string      `json:"p_content_type"`
		UserProperty           interface{} `json:"p_user_property"`
		ResponseTopic          string      `json:"p_response_topic"`
		CorrelationData        string      `json:"p_correlation_data"`
	} `json:"properties"`
}
