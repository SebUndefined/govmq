package govmq

type AuthOnSubscribeModifier struct {
	Topics []TopicConfig `json:"topics"`
}

type AuthOnSubscribeM3Modifier struct {
	AuthOnSubscribeModifier `json:",inline"`
}

type AuthOnSubscribeM5Modifier struct {
	AuthOnSubscribeModifier `json:",inline"`
}

type OnUnsubscribeModifier struct {
	Topics []string `json:"topics"`
}

type OnUnsubscribeM3Modifier struct {
	OnUnsubscribeModifier `json:",inline"`
}

type OnUnsubscribeM5Modifier struct {
	OnUnsubscribeModifier `json:",inline"`
}
