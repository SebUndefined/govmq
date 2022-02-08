package govmq

// ####################################################################
// ######################### AUTH_ON_PUBLISH ##########################
// ####################################################################

// AuthOnPublish d
type AuthOnPublish struct {
	MountPoint string  `json:"mountpoint"`
	ClientId   string  `json:"client_id"`
	Username   *string `json:"username"`
	QoS        int     `json:"qos"`
	Topic      string  `json:"topic"`
	Payload    string  `json:"payload"`
	Retain     bool    `json:"retain"`
}

type AuthOnPublishM3 struct {
	AuthOnPublish `json:",inline"`
}

type AuthOnPublishM5 struct {
	AuthOnPublish `json:",inline"`
	Properties    interface{} `json:"properties"`
}

// ####################################################################
// ############################ ON_PUBLISH ############################
// ####################################################################

type OnPublish struct {
	MountPoint string  `json:"mountpoint"`
	ClientId   string  `json:"client_id"`
	Username   *string `json:"username"`
	QoS        int     `json:"qos"`
	Topic      string  `json:"topic"`
	Payload    string  `json:"payload"`
	Retain     bool    `json:"retain"`
}

type OnPublishM3 struct {
	AuthOnPublish `json:",inline"`
}

type OnPublishM5 struct {
	AuthOnPublish `json:",inline"`
	Properties    interface{} `json:"properties"`
}

// ####################################################################
// ######################### ON_OFFLINE_MESSAGE #######################
// ####################################################################

type OnOfflineMessage struct {
	MountPoint string  `json:"mountpoint"`
	ClientId   string  `json:"client_id"`
	Username   *string `json:"username"`
	QoS        int     `json:"qos"`
	Topic      string  `json:"topic"`
	Payload    string  `json:"payload"`
	Retain     bool    `json:"retain"`
}

// ####################################################################
// ########################### ON_DELIVER #############################
// ####################################################################

type OnDeliver struct {
	MountPoint string  `json:"mountpoint"`
	ClientId   string  `json:"client_id"`
	Username   *string `json:"username"`
	Topic      string  `json:"topic"`
	Payload    string  `json:"payload"`
}

type OnDeliverM3 struct {
	MountPoint string  `json:"mountpoint"`
	ClientId   string  `json:"client_id"`
	Username   *string `json:"username"`
	Topic      string  `json:"topic"`
	Payload    string  `json:"payload"`
}

type OnDeliverM5 struct {
	MountPoint string      `json:"mountpoint"`
	ClientId   string      `json:"client_id"`
	Username   *string     `json:"username"`
	Topic      string      `json:"topic"`
	Payload    string      `json:"payload"`
	Properties interface{} `json:"properties"`
}
