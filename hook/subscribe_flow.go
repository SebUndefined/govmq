package hook

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
