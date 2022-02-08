package govmq

// ####################################################################
// ######################## AUTH_ON_SUBSCRIBE #########################
// ####################################################################

type AuthOnSubscribe struct {
	MountPoint string  `json:"mountpoint"`
	ClientId   string  `json:"client_id"`
	Username   *string `json:"username"`
	Topic      []Topic `json:"topic"`
}

type Topic struct {
	Topic string `json:"topic"`
	Qos   int    `json:"qos"`
}

type AuthOnSubscribeM3 struct {
	AuthOnSubscribe `json:",inline"`
}

type AuthOnSubscribeM5 struct {
	AuthOnSubscribe `json:",inline"`
	Properties      interface{} `json:"properties"`
}

// ####################################################################
// ########################## ON_SUBSCRIBE ############################
// ####################################################################

type OnSubscribe struct {
	MountPoint string  `json:"mountpoint"`
	ClientId   string  `json:"client_id"`
	Username   *string `json:"username"`
	Topic      []Topic `json:"topic"`
}

type OnSubscribeM3 struct {
	OnSubscribe `json:",inline"`
}

type OnSubscribeM5 struct {
	OnSubscribe `json:",inline"`
	Properties  interface{} `json:"properties"`
}

/*####################################################################
########################## ON_UNSUBSCRIBE ##########################
####################################################################*/

// OnUnSubscribe The on_subscribe and on_subscribe_m5 hooks allow your plugin to get
type OnUnSubscribe struct {
	MountPoint string   `json:"mountpoint"`
	ClientId   string   `json:"client_id"`
	Username   *string  `json:"username"`
	Topic      []string `json:"topic"`
}

type OnUnSubscribeM3 struct {
	OnUnSubscribe `json:",inline"`
}

type OnUnSubscribeM5 struct {
	OnUnSubscribe `json:",inline"`
	Properties    interface{} `json:"properties"`
}
