package govmq

type BuilderOption func(r *AuthOnRegisterModifier)

type AuthOnRegisterModifierBuilder struct {
	actions []BuilderOption
}

func (rb *AuthOnRegisterModifierBuilder) WithMountpoint(mountpoint string) *AuthOnRegisterModifierBuilder {
	rb.actions = append(rb.actions, func(r *AuthOnRegisterModifier) {
		r.MountPoint = &mountpoint
	})
	return rb
}
