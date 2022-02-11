package govmq

type BuilderOption func(r *AuthOnRegisterModifier) (*AuthOnRegisterModifier, error)

type MountpointMember interface {
	SetMountpoint(m string)
}

type AuthOnRegisterModifierBuilder struct {
	actions []BuilderOption
}

func WithMountpoint(query interface{}, args ...interface{}) BuilderOption {
	return func(r *AuthOnRegisterModifier) (*AuthOnRegisterModifier, error) {
		ret := db.Where(query, args)
		return ret, ret.Error
	}
}

func (rb *AuthOnRegisterModifierBuilder) WithMountpoint(mountpoint string) *AuthOnRegisterModifierBuilder {
	rb.actions = append(rb.actions, func(r *AuthOnRegisterModifier) {
		r.MountPoint = &mountpoint
	})
	return rb
}
