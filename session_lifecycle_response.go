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

type AuthOnRegisterResponseBuilder struct {
	actions []func(response *AuthOnRegisterResponse)
}

func (rb *AuthOnRegisterResponseBuilder) WithMountpoint(mountpoint string) *AuthOnRegisterResponseBuilder {
	rb.actions = append(rb.actions, func(r *AuthOnRegisterResponse) {
		r.MountPoint = mountpoint
	})
	return rb
}

func (rb *AuthOnRegisterResponseBuilder) WithMaxMessageSize(m int) *AuthOnRegisterResponseBuilder {
	rb.actions = append(rb.actions, func(r *AuthOnRegisterResponse) {
		r.MaxMessageSize = m
	})
	return rb
}

func (rb *AuthOnRegisterResponseBuilder) WithMaxOnlineMessages(m int) *AuthOnRegisterResponseBuilder {
	rb.actions = append(rb.actions, func(r *AuthOnRegisterResponse) {
		r.MaxOnlineMessages = m
	})
	return rb
}

func (rb *AuthOnRegisterResponseBuilder) WithMaxOfflineMessages(m int) *AuthOnRegisterResponseBuilder {
	rb.actions = append(rb.actions, func(r *AuthOnRegisterResponse) {
		r.MaxOfflineMessages = m
	})
	return rb
}

func (rb *AuthOnRegisterResponseBuilder) WithSharedSubscriptionPolicy(s interface{}) *AuthOnRegisterResponseBuilder {
	rb.actions = append(rb.actions, func(r *AuthOnRegisterResponse) {
		r.SharedSubscriptionPolicy = s
	})
	return rb
}

func (rb *AuthOnRegisterResponseBuilder) Build() *AuthOnRegisterResponse {
	r := AuthOnRegisterResponse{}
	for _, a := range rb.actions {
		a(&r)
	}
	return &r
}

type AuthOnRegisterM3ResponseBuilder struct {
	AuthOnRegisterResponseBuilder
	actions []func(response *AuthOnRegisterM3Response)
}

func (rb *AuthOnRegisterM3ResponseBuilder) WithCleanSession(c bool) *AuthOnRegisterM3ResponseBuilder {
	rb.actions = append(rb.actions, func(r *AuthOnRegisterM3Response) {
		r.CleanSession = c
	})
	return rb
}

func (rb *AuthOnRegisterM3ResponseBuilder) Build() *AuthOnRegisterM3Response {
	r := AuthOnRegisterResponse{}
	for _, a := range rb.AuthOnRegisterResponseBuilder.actions {
		a(&r)
	}
	rM3 := AuthOnRegisterM3Response{}
	rM3.AuthOnRegisterResponse = r
	for _, a := range rb.actions {
		a(&rM3)
	}
	return &rM3
}

type AuthOnRegisterM5ResponseBuilder struct {
	AuthOnRegisterResponseBuilder
	actions []func(response *AuthOnRegisterM5Response)
}

func (rb *AuthOnRegisterM5ResponseBuilder) WithCleanStart(c bool) *AuthOnRegisterM5ResponseBuilder {
	rb.actions = append(rb.actions, func(r *AuthOnRegisterM5Response) {
		r.CleanStart = c
	})
	return rb
}

func (rb *AuthOnRegisterM5ResponseBuilder) Build() *AuthOnRegisterM5Response {
	r := AuthOnRegisterResponse{}
	for _, a := range rb.AuthOnRegisterResponseBuilder.actions {
		a(&r)
	}
	rM5 := AuthOnRegisterM5Response{}
	rM5.AuthOnRegisterResponse = r
	for _, a := range rb.actions {
		a(&rM5)
	}
	return &rM5
}
