package govmq

type AuthOnRegisterModifierBuilder struct {
	actions []func(response *AuthOnRegisterModifier)
}

func (rb *AuthOnRegisterModifierBuilder) WithMountpoint(mountpoint string) *AuthOnRegisterModifierBuilder {
	rb.actions = append(rb.actions, func(r *AuthOnRegisterModifier) {
		r.MountPoint = &mountpoint
	})
	return rb
}

func (rb *AuthOnRegisterModifierBuilder) WithMaxMessageSize(m int) *AuthOnRegisterModifierBuilder {
	rb.actions = append(rb.actions, func(r *AuthOnRegisterModifier) {
		r.MaxMessageSize = &m
	})
	return rb
}

func (rb *AuthOnRegisterModifierBuilder) WithMaxOnlineMessages(m int) *AuthOnRegisterModifierBuilder {
	rb.actions = append(rb.actions, func(r *AuthOnRegisterModifier) {
		r.MaxOnlineMessages = &m
	})
	return rb
}

func (rb *AuthOnRegisterModifierBuilder) WithMaxOfflineMessages(m int) *AuthOnRegisterModifierBuilder {
	rb.actions = append(rb.actions, func(r *AuthOnRegisterModifier) {
		r.MaxOfflineMessages = &m
	})
	return rb
}

func (rb *AuthOnRegisterModifierBuilder) WithSharedSubscriptionPolicy(s interface{}) *AuthOnRegisterModifierBuilder {
	rb.actions = append(rb.actions, func(r *AuthOnRegisterModifier) {
		r.SharedSubscriptionPolicy = s
	})
	return rb
}

func (rb *AuthOnRegisterModifierBuilder) Build() *AuthOnRegisterModifier {
	r := AuthOnRegisterModifier{}
	for _, a := range rb.actions {
		a(&r)
	}
	return &r
}

type AuthOnRegisterM3ModifierBuilder struct {
	AuthOnRegisterModifierBuilder
	actions []func(response *AuthOnRegisterM3Modifier)
}

func (rb *AuthOnRegisterM3ModifierBuilder) WithCleanSession(c bool) *AuthOnRegisterM3ModifierBuilder {
	rb.actions = append(rb.actions, func(r *AuthOnRegisterM3Modifier) {
		r.CleanSession = &c
	})
	return rb
}

func (rb *AuthOnRegisterM3ModifierBuilder) Build() *AuthOnRegisterM3Modifier {
	rM3 := AuthOnRegisterM3Modifier{
		AuthOnRegisterModifier: rb.AuthOnRegisterModifierBuilder.Build(),
	}
	for _, a := range rb.actions {
		a(&rM3)
	}
	return &rM3
}

type AuthOnRegisterM5ResponseBuilder struct {
	AuthOnRegisterModifierBuilder
	actions []func(response *AuthOnRegisterM5Modifier)
}

func (rb *AuthOnRegisterM5ResponseBuilder) WithCleanStart(c bool) *AuthOnRegisterM5ResponseBuilder {
	rb.actions = append(rb.actions, func(r *AuthOnRegisterM5Modifier) {
		r.CleanStart = &c
	})
	return rb
}

func (rb *AuthOnRegisterM5ResponseBuilder) Build() *AuthOnRegisterM5Modifier {
	rM5 := AuthOnRegisterM5Modifier{
		AuthOnRegisterModifier: rb.AuthOnRegisterModifierBuilder.Build(),
	}
	for _, a := range rb.actions {
		a(&rM5)
	}
	return &rM5
}
