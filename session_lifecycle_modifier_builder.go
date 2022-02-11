package govmq

type AuthOnRegisterModifierBuilder struct {
	actions []func(r *AuthOnRegisterModifier)
}

func (rb *AuthOnRegisterModifierBuilder) WithMountpoint(mountpoint string) *AuthOnRegisterModifierBuilder {
	rb.actions = append(rb.actions, func(r *AuthOnRegisterModifier) {
		r.MountPoint = &mountpoint
	})
	return rb
}

type AuthOnRegisterM3ModifierBuilder struct {
	actions []func(r *AuthOnRegisterM3Modifier)
}

func (rb *AuthOnRegisterM3ModifierBuilder) WithBase(f func(*AuthOnRegisterModifierBuilder) *AuthOnRegisterModifierBuilder) *AuthOnRegisterM3ModifierBuilder {

	rb.actions = append(rb.actions, func(r *AuthOnRegisterM3Modifier) {
		r.CleanSession = &c
	})
	return rb
}

func (rb *AuthOnRegisterM3ModifierBuilder) WithCleanSession(c bool) *AuthOnRegisterM3ModifierBuilder {
	rb.actions = append(rb.actions, func(r *AuthOnRegisterM3Modifier) {
		r.CleanSession = &c
	})
	return rb
}
